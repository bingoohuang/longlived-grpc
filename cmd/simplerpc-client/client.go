package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"

	sgRPC "github.com/bingoohuang/longlivedgprc/protos/simple/testgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const socket = "127.0.0.1:2008"

func main() {
	// grpc uses HTTP 2 which is by default uses SSL
	// we use insecure (we can also use the credentials)
	conn, err := grpc.Dial(socket, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Could not connect to : ", socket)
	}
	log.Println("Connected to ", socket)
	defer conn.Close()
	client := sgRPC.NewSimpleServiceClient(conn) // Using this connection to use the SimpleService
	// unary request
	makeUnaryRequest(client)

	// Client streaming
	makeClientStreaming(client)

	// Server Streaming
	makeServerStreaming(client)

	// Bi-Directional
	makeBidirectional(client)
}

func makeUnaryRequest(c sgRPC.SimpleServiceClient) {
	log.Println("Making Unary Request")
	req := &sgRPC.SimpleRequest{RequestNeed: "To test!", RequestId: math.MaxUint64}
	log.Printf("Request - %+v", req)
	res, err := c.RPCRequest(context.Background(), req)
	handleAndFatalError(err)
	log.Printf("Response - %+v", res)
}

func makeClientStreaming(c sgRPC.SimpleServiceClient) {
	log.Println("Client Streaming")
	stream, err := c.ClientStreaming(context.Background())
	handleAndFatalError(err)

	for i := 1; i < 10; i++ {
		req := fmt.Sprintf("Request number : %d", i)
		r := &sgRPC.SimpleRequest{RequestNeed: req, RequestId: math.MaxUint64}
		log.Printf("Request - %+v", r)
		stream.Send(r)
	}
	response, err := stream.CloseAndRecv()
	handleAndFatalError(err)

	log.Printf("Response - %+v", response)
}

func makeServerStreaming(c sgRPC.SimpleServiceClient) {
	log.Println("Server Streaming")
	req := &sgRPC.SimpleRequest{RequestNeed: "Need stream response", RequestId: math.MaxUint64}
	log.Printf("Request - %+v", req)
	serverStream, err := c.ServerStreaming(context.Background(), req)
	handleAndFatalError(err)

	for {
		response, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		handleAndFatalError(err)

		log.Printf("Response - %+v", response)
	}
}

func makeBidirectional(c sgRPC.SimpleServiceClient) {
	log.Println("Bi-Directional Streaming")
	biStream, err := c.StreamingBiDirectional(context.Background())
	handleAndFatalError(err)
	defer biStream.CloseSend()
	var id uint64 = math.MaxUint64

	// here the communication sequence is completely depends on how the server is implemented.
	// if the server is implemetend to give response to all the response at the end or
	// one after another it all compeletely depends on the implementation
	for i := uint64(0); i < 10; i++ {
		req := fmt.Sprintf("My request %d", i)
		r := &sgRPC.SimpleRequest{RequestNeed: req, RequestId: id - i}
		log.Printf("Request - %+v", r)
		biStream.Send(r)
		reply, err := biStream.Recv()

		handleAndPrintError(err)
		log.Printf("Response - %+v", reply)
	}
}

func handleAndPrintError(e error) {
	if e != nil {
		log.Println(e)
	}
}

func handleAndFatalError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
