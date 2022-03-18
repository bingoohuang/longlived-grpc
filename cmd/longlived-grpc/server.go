package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"longlived-gprc/protos"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	pAddr := flag.String("addr", ":7070", "listen address for the Grpc server")
	pClient := flag.Bool("client", false, "start as Grpc client")
	flag.Parse()

	var stopper Stopper
	if *pClient {
		stopper = startClients(*pAddr)
	} else {
		stopper = startServer(*pAddr)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	log.Printf("signal %v received", <-c)
	stopper.Stop()
	log.Print("exit")
}

type Stopper interface {
	Stop()
}

func startServer(address string) *grpc.Server {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("listen on %s failed: %v", address, err)
	}

	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	server := &longlivedServer{}

	// Start sending data to subscribers
	go server.mockDataGenerator()

	// Register the server
	protos.RegisterLonglivedServer(grpcServer, server)

	log.Printf("Starting server on address %s", lis.Addr().String())

	go func() {
		// Start listening
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}()

	return grpcServer
}

type longlivedServer struct {
	protos.UnimplementedLonglivedServer
	subscribers sync.Map // subscribers is a concurrent map that holds mapping from a client ID to it's subscriber
}

type sub struct {
	stream   protos.Longlived_SubscribeServer // stream is the server side of the RPC stream
	finished chan<- bool                      // finished is used to signal closure of a client subscribing goroutine
}

// Subscribe handles a subscribe request from a client
func (s *longlivedServer) Subscribe(request *protos.Request, stream protos.Longlived_SubscribeServer) error {
	// Handle subscribe request
	log.Printf("Received subscribe request from ID: %d", request.Id)

	fin := make(chan bool)
	// Save the subscriber stream according to the given client ID
	s.subscribers.Store(request.Id, sub{stream: stream, finished: fin})

	ctx := stream.Context()
	// Keep this scope alive because once this scope exits - the stream is closed
	for {
		select {
		case <-fin:
			log.Printf("Closing stream for client ID: %d", request.Id)
			return nil
		case <-ctx.Done():
			log.Printf("Client ID %d has disconnected", request.Id)
			return nil
		}
	}
}

// NotifyReceived handles a NotifyReceived request from a client
func (s *longlivedServer) NotifyReceived(ctx context.Context, request *protos.Request) (*protos.Response, error) {
	log.Printf("NotifyReceived: %d", request.Id)
	return &protos.Response{}, nil
}

// Unsubscribe handles a unsubscribe request from a client
// Note: this function is not called but it here as an example of an unary RPC for unsubscribing clients
func (s *longlivedServer) Unsubscribe(ctx context.Context, request *protos.Request) (*protos.Response, error) {
	v, ok := s.subscribers.Load(request.Id)
	if !ok {
		return nil, fmt.Errorf("failed to load subscriber key: %d", request.Id)
	}
	sub, ok := v.(sub)
	if !ok {
		return nil, fmt.Errorf("failed to cast subscriber value: %T", v)
	}
	select {
	case sub.finished <- true:
		log.Printf("Unsubscribed client: %d", request.Id)
	default:
		// Default case is to avoid blocking in case client has already unsubscribed
	}
	s.subscribers.Delete(request.Id)
	return &protos.Response{}, nil
}

func (s *longlivedServer) mockDataGenerator() {
	log.Println("Starting data generation")
	for {
		time.Sleep(time.Second)

		// A list of clients to unsubscribe in case of error
		var unsubscribe []int32

		// Iterate over all subscribers and send data to each client
		s.subscribers.Range(func(k, v interface{}) bool {
			id, ok := k.(int32)
			if !ok {
				log.Printf("Failed to cast subscriber key: %T", k)
				return false
			}
			sub, ok := v.(sub)
			if !ok {
				log.Printf("Failed to cast subscriber value: %T", v)
				return false
			}
			// Send data over the gRPC stream to the client
			mockData := fmt.Sprintf("data mock for: %d", id)
			if err := sub.stream.Send(&protos.Response{Data: mockData}); err != nil {
				log.Printf("Failed to send data to client: %v", err)
				select {
				case sub.finished <- true:
					log.Printf("Unsubscribed client: %d", id)
				default:
					// Default case is to avoid blocking in case client has already unsubscribed
				}
				// In case of error the client would re-subscribe so close the subscriber stream
				unsubscribe = append(unsubscribe, id)
			}
			return true
		})

		// Unsubscribe erroneous client streams
		for _, id := range unsubscribe {
			s.subscribers.Delete(id)
		}
	}
}
