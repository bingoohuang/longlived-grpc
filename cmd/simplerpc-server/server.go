package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	longlivedgrpc "github.com/bingoohuang/longlivedgprc"
	protos "github.com/bingoohuang/longlivedgprc/protos/simple/testgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/reflection"
)

const socket = "127.0.0.1:2008"

type Server struct {
	protos.SimpleServiceServer
}

func main() {
	lisn, err := net.Listen("tcp", socket)
	if err != nil {
		log.Fatalln("Errored while Listen to : ", socket, err)
	}
	log.Println("Listening at ", socket)
	s := grpc.NewServer()
	protos.RegisterSimpleServiceServer(s, &Server{}) // registering our grpc server with our grpc service.
	if longlivedgrpc.IsEnvEnabled("GRPC_REFLECTION") {
		reflection.Register(s)
	}
	if longlivedgrpc.IsEnvEnabled("GRPC_CHANNELZ") {
		service.RegisterChannelzServiceToServer(s)
	}
	if err := s.Serve(lisn); err != nil {
		log.Fatalln("Errored while Serving : ", socket, err)
	}
}

func (s *Server) RPCRequest(ctx context.Context, req *protos.SimpleRequest) (*protos.SimpleResponse, error) {
	log.Println("Unary request")
	log.Printf("Request - %+v", req)
	response := &protos.SimpleResponse{Response: "Here is your response", ResponseId: math.MaxUint64}
	log.Printf("Response - %+v", response)
	return response, nil
}

func (s *Server) ClientStreaming(stream protos.SimpleService_ClientStreamingServer) error {
	log.Println("ClientStreaming RPC")
	var responseID uint64 = math.MaxUint64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			response := &protos.SimpleResponse{Response: "Here is your response", ResponseId: responseID}
			log.Printf("Response - %+v", response)
			responseID--
			stream.SendAndClose(response)
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Request - %+v", req)
	}
	return nil
}

func (s *Server) ServerStreaming(req *protos.SimpleRequest, stream protos.SimpleService_ServerStreamingServer) error {
	log.Println("ServerStreaming RPC")
	log.Printf("Request- %+v", req)
	var responseID uint64 = math.MaxUint64
	for i := uint64(0); i < 10; i++ {
		res := fmt.Sprintf("Here is the response %d", i)
		rsp := &protos.SimpleResponse{Response: res, ResponseId: responseID - i}
		log.Printf("Response - %+v", rsp)
		stream.Send(rsp)
	}

	return nil
}

func (s *Server) StreamingBiDirectional(stream protos.SimpleService_StreamingBiDirectionalServer) error {
	log.Println("StreamingBiDirectional RPC")
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				log.Println("Errored in stream Recv", err)
			}
			break
		}
		log.Printf("Request - %+v", msg)
		r := fmt.Sprintf("Response for your request - %v", msg.RequestNeed)
		rsp := &protos.SimpleResponse{Response: r, ResponseId: msg.RequestId - 1}
		stream.Send(rsp)
		log.Printf("Response - %+v", rsp)
	}

	return nil
}
