package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	longlivedgrpc "github.com/bingoohuang/longlivedgprc"

	"github.com/bingoohuang/longlivedgprc/protos"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/reflection"
)

type Stoppers struct {
	Values []Stopper
}

func (s Stoppers) Stop() {
	for _, st := range s.Values {
		st.Stop()
	}
}

func (s *Stoppers) DeleteAllClients() (ids []string) {
	var values []Stopper
	for _, st := range s.Values {
		if st.Mode() == ModeClient {
			ids = append(ids, st.GetID())
			st.Stop()
		} else {
			values = append(values, st)
		}
	}
	s.Values = values

	return
}

func (s *Stoppers) DeleteClient(id string) (Stopper, bool) {
	for i, st := range s.Values {
		if st.Mode() == ModeClient && st.GetID() == id {
			st.Stop()
			s.Values = append(s.Values[:i], s.Values[i+1:]...)
			return st, true
		}
	}

	return nil, false
}

func (s *Stoppers) Add(client Stopper) {
	s.Values = append(s.Values, client)
}

func (s Stoppers) List(mode Mode) (ids []string) {
	for _, st := range s.Values {
		if st.Mode() == mode {
			ids = append(ids, st.GetID())
		}
	}

	return
}

func (s Stoppers) StopMode(mode Mode) (ids []string) {
	for _, st := range s.Values {
		if st.Mode() == mode {
			st.Stop()
			ids = append(ids, st.GetID())
		}
	}

	return ids
}

type Rsp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func serverRestHandle(addrs []string) gin.HandlerFunc {
	return func(g *gin.Context) {
		switch g.Param("action") {
		case "start":
			if ids := stoppers.List(ModeServer); len(ids) == 0 {
				for _, addr := range addrs {
					stoppers.Add(startServer(addr))
				}
				g.JSON(200, Rsp{Status: 200, Message: "OK"})
			} else {
				g.JSON(200, Rsp{Status: 304, Message: "Server already started", Data: H{"ids": ids}})
			}
		case "stop":
			ids := stoppers.StopMode(ModeServer)
			g.JSON(200, Rsp{Status: 200, Message: "Stop successfully", Data: H{"ids": ids}})
		default:
			g.JSON(404, Rsp{Status: 404, Message: "unsupported path", Data: H{"path": g.Request.URL.Path}})
		}
	}
}

type Mode int

const (
	ModeServer Mode = iota
	ModeClient
)

type Stopper interface {
	GetID() string
	Mode() Mode
	Start()
	Stop()
}

type Server struct {
	ctx     context.Context
	cancelF context.CancelFunc

	protos.UnimplementedLonglivedServer
	subscribers sync.Map // subscribers is a concurrent map that holds mapping from a client ID to it's subscriber
	*grpc.Server
	ID      string
	Address string
}

func (s *Server) GetID() string { return s.ID }
func (s *Server) Mode() Mode    { return ModeServer }
func (s *Server) Stop() {
	s.cancelF()
	s.Server.Stop()
}

func (s *Server) Start() {
	// Start sending data to subscribers
	go s.mockDataGenerator()

	go func() {
		lis, err := net.Listen("tcp", s.Address)
		if err != nil {
			log.Printf("E! listen on %s failed: %v", s.Address, err)
			return
		}

		log.Printf("Starting server on address %s", lis.Addr().String())

		// Start listening
		if err := s.Server.Serve(lis); err != nil {
			log.Printf("E! listen failed: %v", err)
		} else {
			log.Printf("Server stopped")
		}

		if err = lis.Close(); err != nil {
			log.Printf("close listen failed: %v", err)
		}
	}()
}

func startServer(address string) *Server {
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	ctx, cancelF := context.WithCancel(context.Background())
	server := &Server{ctx: ctx, cancelF: cancelF, Address: address, Server: grpcServer, ID: ksuid.New().String()}

	protos.RegisterLonglivedServer(grpcServer, server)
	if longlivedgrpc.IsEnvEnabled("GRPC_REFLECTION") {
		reflection.Register(grpcServer)
	}
	if longlivedgrpc.IsEnvEnabled("GRPC_CHANNELZ") {
		service.RegisterChannelzServiceToServer(server)
	}

	server.Start()
	return server
}

type subscribe struct {
	stream   protos.Longlived_SubscribeServer // stream is the server side of the RPC stream
	finished chan<- bool                      // finished is used to signal closure of a client subscribing goroutine
}

// Subscribe handles a subscribe request from a client
func (s *Server) Subscribe(request *protos.Request, stream protos.Longlived_SubscribeServer) error {
	peerAddr := longlivedgrpc.GetPeerAddr(stream.Context())
	realAddr := longlivedgrpc.GetRealAddr(stream.Context())
	log.Printf("Subscribed, serverAddr: %q, peerAddr: %q, realAddr: %q", s.Address, peerAddr, realAddr)

	// Handle subscribe request
	log.Printf("Received subscribe request from ID: %s", request.Id)

	fin := make(chan bool)
	// Save the subscriber stream according to the given client ID
	s.subscribers.Store(request.Id, subscribe{stream: stream, finished: fin})

	// Keep this scope alive because once this scope exits - the stream is closed
	for {
		select {
		case <-fin:
			log.Printf("Closing stream for client ID: %s", request.Id)
			return nil
		case <-stream.Context().Done():
			log.Printf("Client ID %s has disconnected", request.Id)
			return nil
		}
	}
}

// NotifyReceived handles a NotifyReceived request from a client
func (s *Server) NotifyReceived(ctx context.Context, request *protos.Request) (*protos.Response, error) {
	peerAddr := longlivedgrpc.GetPeerAddr(ctx)
	realAddr := longlivedgrpc.GetRealAddr(ctx)
	log.Printf("NotifyReceived: %q, serverAddr: %q, peerAddr: %q, realAddr: %q", s.Address, request.Id, peerAddr, realAddr)
	return &protos.Response{Data: fmt.Sprintf("NotifyReceived: %s", request.Id)}, nil
}

// Unsubscribe handles a unsubscribe request from a client
// Note: this function is not called but it here as an example of an unary RPC for unsubscribing clients
func (s *Server) Unsubscribe(ctx context.Context, request *protos.Request) (*protos.Response, error) {
	peerAddr := longlivedgrpc.GetPeerAddr(ctx)
	realAddr := longlivedgrpc.GetRealAddr(ctx)
	log.Printf("Unsubscribe: %q, serverAddr: %q, peerAddr: %q, realAddr: %q", s.Address, request.Id, peerAddr, realAddr)

	v, ok := s.subscribers.Load(request.Id)
	if !ok {
		return nil, fmt.Errorf("failed to load subscriber key: %s", request.Id)
	}
	sub, ok := v.(subscribe)
	if !ok {
		return nil, fmt.Errorf("failed to cast subscriber value: %T", v)
	}
	select {
	case sub.finished <- true:
		log.Printf("Unsubscribed client: %s", request.Id)
	default:
		// Default case is to avoid blocking in case client has already unsubscribed
	}
	s.subscribers.Delete(request.Id)
	return &protos.Response{}, nil
}

func (s *Server) mockDataGenerator() {
	log.Printf("Starting mock data generation")
	defer log.Printf("Stopped mock data generation")
	for s.ctx.Err() == nil {
		Sleep(s.ctx, 3*time.Second)

		// A list of clients to unsubscribe in case of error
		var unsubscribe []string

		// Iterate over all subscribers and send data to each client
		s.subscribers.Range(func(k, v interface{}) bool {
			id, ok := k.(string)
			if !ok {
				log.Printf("Failed to cast subscriber key: %T", k)
				return false
			}
			sub, ok := v.(subscribe)
			if !ok {
				log.Printf("Failed to cast subscriber value: %T", v)
				return false
			}
			// Send data over the gRPC stream to the client
			mockData := fmt.Sprintf("data mock for: %s", id)
			if err := sub.stream.Send(&protos.Response{Data: mockData}); err != nil {
				log.Printf("Failed to send data to client: %v", err)
				select {
				case sub.finished <- true:
					log.Printf("Unsubscribed client: %s", id)
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
