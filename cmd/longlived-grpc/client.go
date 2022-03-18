package main

import (
	"context"
	"log"
	"longlived-gprc/protos"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func startClient(address string) {
	// Create multiple clients and start receiving data
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		client, err := mkLonglivedClient(address, int32(i))
		if err != nil {
			log.Fatal(err)
		}
		// Dispatch client goroutine
		go client.start(&wg)
		time.Sleep(time.Second * 2)
	}

	wg.Wait()
}

// longlivedClient holds the long lived gRPC client fields
type longlivedClient struct {
	client protos.LonglivedClient // client is the long lived gRPC client
	conn   *grpc.ClientConn       // conn is the client gRPC connection
	id     int32                  // id is the client ID used for subscribing
}

// mkLonglivedClient creates a new client instance
func mkLonglivedClient(address string, id int32) (*longlivedClient, error) {
	conn, err := mkConnection(address)
	if err != nil {
		return nil, err
	}
	return &longlivedClient{
		client: protos.NewLonglivedClient(conn),
		conn:   conn,
		id:     id,
	}, nil
}

// close is not used but is here as an example of how to close the gRPC client connection
func (c *longlivedClient) Close() error {
	return c.conn.Close()
}

// subscribe subscribes to messages from the gRPC server
func (c *longlivedClient) Subscribe() (protos.Longlived_SubscribeClient, error) {
	log.Printf("Subscribing client ID %d", c.id)
	return c.client.Subscribe(context.Background(), &protos.Request{Id: c.id})
}

// unsubscribe unsubscribes to messages from the gRPC server
func (c *longlivedClient) Unsubscribe() error {
	log.Printf("Unsubscribing client ID %d", c.id)
	_, err := c.client.Unsubscribe(context.Background(), &protos.Request{Id: c.id})
	return err
}

func (c *longlivedClient) start(wg *sync.WaitGroup) {
	defer wg.Done()

	var err error
	// stream is the client side of the RPC stream
	var stream protos.Longlived_SubscribeClient
	for {
		if stream == nil {
			if stream, err = c.Subscribe(); err != nil {
				log.Printf("Failed to subscribe: %v", err)
				c.sleep()
				continue // Retry on failure
			}
		}
		response, err := stream.Recv()
		if err != nil {
			log.Printf("Failed to receive message: %v", err)
			// Clearing the stream will force the client to resubscribe on next iteration
			stream = nil
			c.sleep()
			continue // Retry on failure
		}
		log.Printf("Client ID %d got response: %q", c.id, response.Data)
	}
}

// sleep is used to give the server time to unsubscribe the client and reset the stream
func (c *longlivedClient) sleep() {
	time.Sleep(time.Second * 5)
}

func mkConnection(address string) (*grpc.ClientConn, error) {
	if address == ":" || address == "" {
		address = "127.0.0.1:7070"
	} else if strings.HasPrefix(address, ":") {
		address = "127.0.0.1" + address
	}
	return grpc.Dial(address, []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}...)
}
