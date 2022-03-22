package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"longlived-gprc/protos"
	_ "longlived-gprc/resolver"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

type H map[string]any

func clientRestHandle(addr string) gin.HandlerFunc {
	ctx, cancel := context.WithCancel(context.Background())
	conn, err := mkConnection(ctx, addr)
	if err != nil {
		log.Fatal(err)
	}

	clientContainer := &ClientContainer{cancel: cancel, conn: conn}

	return func(g *gin.Context) {
		switch g.Param("action") {
		case "list":
			clients := stoppers.List(ModeClient)
			g.JSON(200, Rsp{Status: 200, Message: "OK", Data: clients})
		case "start":
			client := startClients(ctx, clientContainer)
			clientID := client.GetID()
			stoppers.Add(client)
			g.JSON(200, Rsp{Status: 200, Message: "OK", Data: H{"clientID": clientID}})
		case "notify":
			c := protos.NewLonglivedClient(clientContainer.conn)
			rsp, err := c.NotifyReceived(ctx, &protos.Request{Id: ksuid.New().String()})
			if err != nil {
				g.JSON(500, Rsp{Status: 500, Message: "error", Data: err.Error()})
			} else {
				g.JSON(200, Rsp{Status: 200, Message: "notified", Data: rsp})
			}
		case "stop":
			clientID := g.Query("id")
			if _, ok := stoppers.DeleteClient(clientID); ok {
				g.JSON(200, Rsp{Status: 200, Message: "stop and deleted", Data: H{"clientID": clientID}})
			} else {
				g.JSON(200, Rsp{Status: 404, Message: "client not found", Data: H{"clientID": clientID}})
			}
		default:
			g.JSON(404, Rsp{Status: 404, Message: "unsupported path", Data: H{"path": g.Request.URL.Path}})
		}
	}
}

type ClientContainer struct {
	wg     sync.WaitGroup
	cancel context.CancelFunc
	conn   *grpc.ClientConn
}

func (c *ClientContainer) Mode() Mode {
	return ModeClient
}

func (c *ClientContainer) Stop() {
	c.cancel()
	c.wg.Wait()
	if err := c.conn.Close(); err != nil {
		log.Printf("close connection failed: %v", err)
	}
}

func (c *ClientContainer) AddWg() {
	c.wg.Add(1)
}

func startClients(ctx context.Context, clients *ClientContainer) *longlivedClient {
	clients.AddWg()
	client := newLonglivedClient(ctx, clients)

	// Dispatch client goroutine
	go client.Start()

	return client
}

// longlivedClient holds the long-lived gRPC client fields
type longlivedClient struct {
	ctx             context.Context
	client          protos.LonglivedClient // client is the long-lived gRPC client
	ID              string                 // id is the client ID used for subscribing
	cancelF         context.CancelFunc
	clientContainer *ClientContainer
}

func (c *longlivedClient) Stop() {
	log.Printf("Unsubscribe client ID: %s", c.ID)
	response, err := c.client.Unsubscribe(c.ctx, &protos.Request{Id: c.ID})
	if err != nil {
		log.Printf("E! unsubscribe failed: %v", err)
	} else {
		log.Printf("unsubscribe successfully, response: %s", response.Data)
	}

	c.cancelF()
}

func (c *longlivedClient) Mode() Mode    { return ModeClient }
func (c *longlivedClient) GetID() string { return c.ID }

// newLonglivedClient creates a new client instance
func newLonglivedClient(ctx context.Context, clientContainer *ClientContainer) *longlivedClient {
	ctx, cancelF := context.WithCancel(ctx)
	return &longlivedClient{
		ctx:             ctx,
		cancelF:         cancelF,
		clientContainer: clientContainer,
		client:          protos.NewLonglivedClient(clientContainer.conn),
		ID:              ksuid.New().String(),
	}
}

// Subscribe subscribes to message from the gRPC server
func (c *longlivedClient) Subscribe() (protos.Longlived_SubscribeClient, error) {
	log.Printf("Subscribing client ID %s", c.ID)
	return c.client.Subscribe(c.ctx, &protos.Request{Id: c.ID})
}

// Unsubscribe unsubscribes to message from the gRPC server
func (c *longlivedClient) Unsubscribe() error {
	log.Printf("Unsubscribing client ID %s", c.ID)
	_, err := c.client.Unsubscribe(c.ctx, &protos.Request{Id: c.ID})
	return err
}

func (c *longlivedClient) Start() {
	defer c.clientContainer.wg.Done()

	var err error
	// stream is the client side of the RPC stream
	var stream protos.Longlived_SubscribeClient

	for c.ctx.Err() == nil {
		if stream == nil {
			if stream, err = c.Subscribe(); err != nil {
				log.Printf("Failed to subscribe: %v", err)
				sleep(c.ctx, 5*time.Second)
				continue // Retry on failure
			}
		}

		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("stream is finished: %v", err)
				break
			}

			log.Printf("Failed to receive message: %v", err)
			// Clearing the stream will force the client to resubscribe on next iteration
			stream = nil
			sleep(c.ctx, 5*time.Second)
			continue // Retry on failure
		}

		log.Printf("Client ID %s got response: %q", c.ID, response.Data)
		_, _ = c.client.NotifyReceived(c.ctx, &protos.Request{Id: c.ID})
	}

	log.Printf("Client ID %s exited", c.ID)
}

// sleep is used to give the server time to unsubscribe the client and reset the stream
func sleep(ctx context.Context, d time.Duration) {
	select {
	case <-ctx.Done():
	case <-time.After(d):
	}
}

func mkConnection(ctx context.Context, address string) (*grpc.ClientConn, error) {
	serviceConfig := fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)
	options := []grpc.DialOption{
		// grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(serviceConfig),
	}

	if !strings.HasPrefix(address, "static:") {
		address = "static:" + address
	}

	return grpc.DialContext(ctx, address, options...)
}
