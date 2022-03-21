package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"longlived-gprc/protos"
	_ "longlived-gprc/resolver"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type Clients struct {
	wg     sync.WaitGroup
	cancel context.CancelFunc
	conn   *grpc.ClientConn
}

func (c *Clients) Stop() {
	c.cancel()
	c.wg.Wait()
	if err := c.conn.Close(); err != nil {
		log.Printf("close connection failed: %v", err)
	}
}

func (c *Clients) AddWg() {
	c.wg.Add(1)
}

func startClients(address string) *Clients {
	ctx, cancel := context.WithCancel(context.Background())
	conn, err := mkConnection(ctx, address)
	if err != nil {
		log.Fatal(err)
	}

	clients := &Clients{cancel: cancel, conn: conn}

	for i := 1; i <= 10; i++ {
		clients.AddWg()
		client := newLonglivedClient(ctx, conn, int32(i))

		// Dispatch client goroutine
		go client.start(&clients.wg)
	}

	return clients
}

// longlivedClient holds the long-lived gRPC client fields
type longlivedClient struct {
	ctx    context.Context
	client protos.LonglivedClient // client is the long-lived gRPC client
	id     int32                  // id is the client ID used for subscribing
}

// newLonglivedClient creates a new client instance
func newLonglivedClient(ctx context.Context, conn *grpc.ClientConn, id int32) *longlivedClient {
	return &longlivedClient{
		ctx:    ctx,
		client: protos.NewLonglivedClient(conn),
		id:     id,
	}
}

// Subscribe subscribes to message from the gRPC server
func (c *longlivedClient) Subscribe() (protos.Longlived_SubscribeClient, error) {
	log.Printf("Subscribing client ID %d", c.id)
	return c.client.Subscribe(c.ctx, &protos.Request{Id: c.id})
}

// Unsubscribe unsubscribes to message from the gRPC server
func (c *longlivedClient) Unsubscribe() error {
	log.Printf("Unsubscribing client ID %d", c.id)
	_, err := c.client.Unsubscribe(c.ctx, &protos.Request{Id: c.id})
	return err
}

func (c *longlivedClient) start(wg *sync.WaitGroup) {
	defer wg.Done()

	var err error
	// stream is the client side of the RPC stream
	var stream protos.Longlived_SubscribeClient

	for c.ctx.Err() == nil {
		if stream == nil {
			if stream, err = c.Subscribe(); err != nil {
				log.Printf("Failed to subscribe: %v", err)
				c.sleep()
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
			c.sleep()
			continue // Retry on failure
		}

		log.Printf("Client ID %d got response: %q", c.id, response.Data)

		_, _ = c.client.NotifyReceived(c.ctx, &protos.Request{Id: c.id})
	}

	log.Printf("Client ID %d exited", c.id)
}

// sleep is used to give the server time to unsubscribe the client and reset the stream
func (c *longlivedClient) sleep() {
	select {
	case <-c.ctx.Done():
	case <-time.After(5 * time.Second):
	}
}

func mkConnection(ctx context.Context, address string) (*grpc.ClientConn, error) {
	if address == ":" || address == "" {
		address = "127.0.0.1:7070"
	} else if strings.HasPrefix(address, ":") {
		address = "127.0.0.1" + address
	}

	return grpc.DialContext(ctx, address,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
	)
}
