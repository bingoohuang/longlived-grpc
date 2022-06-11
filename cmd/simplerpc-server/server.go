package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	sgRPC "github.com/bingoohuang/longlivedgprc/protos/simple/testgrpc"
	"google.golang.org/grpc"
)

const socket = "127.0.0.1:2008"

type Server struct {
	sgRPC.SimpleServiceServer
}

func main() {
	lisn, err := net.Listen("tcp", socket)
	if err != nil {
		log.Fatalln("Errored while Listen to : ", socket, err)
	}
	log.Println("Listening at ", socket)
	s := grpc.NewServer()
	sgRPC.RegisterSimpleServiceServer(s, &Server{}) // registering our grpc server with our grpc service.
	if err := s.Serve(lisn); err != nil {
		log.Fatalln("Errored while Serving : ", socket, err)
	}
}

func (s *Server) RPCRequest(ctx context.Context, req *sgRPC.SimpleRequest) (*sgRPC.SimpleResponse, error) {
	log.Println("Unary request")
	log.Printf("Request - %v\n", req)
	response := &sgRPC.SimpleResponse{Response: "Here is your response"}
	log.Printf("Response - %v\n", response)
	return response, nil
}

func (s *Server) ClientStreaming(stream sgRPC.SimpleService_ClientStreamingServer) error {
	log.Println("ClientStreaming RPC")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			response := &sgRPC.SimpleResponse{Response: "Here is your response"}
			log.Printf("Response - %v\n", response)
			stream.SendAndClose(response)
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Request - %v\n", req)
	}
	return nil
}

func (s *Server) ServerStreaming(req *sgRPC.SimpleRequest, stream sgRPC.SimpleService_ServerStreamingServer) error {
	log.Println("ServerStreaming RPC")
	log.Printf("Request- %v", req)
	for i := 1; i < 10; i++ {
		res := fmt.Sprintf("Here is the response %d", i)
		log.Printf("Response - %v", res)
		stream.Send(&sgRPC.SimpleResponse{Response: res})
	}

	return nil
}

func (s *Server) StreamingBiDirectional(stream sgRPC.SimpleService_StreamingBiDirectionalServer) error {
	log.Println("StreamingBiDirectional RPC")

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Errored in stream Recv", err)
			break
		}
		log.Printf("Request - %v", msg)
		r := fmt.Sprintf("Response for your request - %v", msg.RequestNeed)
		log.Printf("Response - %v\n", r)
		stream.Send(&sgRPC.SimpleResponse{Response: r})
	}

	return nil
}
