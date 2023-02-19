// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: simple.proto

package testgrpc

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SimpleServiceClient is the client API for SimpleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleServiceClient interface {
	// unary RPC
	RPCRequest(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// Server Streaming
	ServerStreaming(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (SimpleService_ServerStreamingClient, error)
	// Client Streaming
	ClientStreaming(ctx context.Context, opts ...grpc.CallOption) (SimpleService_ClientStreamingClient, error)
	// Bi-Directional Streaming
	StreamingBiDirectional(ctx context.Context, opts ...grpc.CallOption) (SimpleService_StreamingBiDirectionalClient, error)
}

type simpleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleServiceClient(cc grpc.ClientConnInterface) SimpleServiceClient {
	return &simpleServiceClient{cc}
}

func (c *simpleServiceClient) RPCRequest(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/SimpleService/RPCRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleServiceClient) ServerStreaming(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (SimpleService_ServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &SimpleService_ServiceDesc.Streams[0], "/SimpleService/ServerStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &simpleServiceServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SimpleService_ServerStreamingClient interface {
	Recv() (*SimpleResponse, error)
	grpc.ClientStream
}

type simpleServiceServerStreamingClient struct {
	grpc.ClientStream
}

func (x *simpleServiceServerStreamingClient) Recv() (*SimpleResponse, error) {
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *simpleServiceClient) ClientStreaming(ctx context.Context, opts ...grpc.CallOption) (SimpleService_ClientStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &SimpleService_ServiceDesc.Streams[1], "/SimpleService/ClientStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &simpleServiceClientStreamingClient{stream}
	return x, nil
}

type SimpleService_ClientStreamingClient interface {
	Send(*SimpleRequest) error
	CloseAndRecv() (*SimpleResponse, error)
	grpc.ClientStream
}

type simpleServiceClientStreamingClient struct {
	grpc.ClientStream
}

func (x *simpleServiceClientStreamingClient) Send(m *SimpleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *simpleServiceClientStreamingClient) CloseAndRecv() (*SimpleResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *simpleServiceClient) StreamingBiDirectional(ctx context.Context, opts ...grpc.CallOption) (SimpleService_StreamingBiDirectionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &SimpleService_ServiceDesc.Streams[2], "/SimpleService/StreamingBiDirectional", opts...)
	if err != nil {
		return nil, err
	}
	x := &simpleServiceStreamingBiDirectionalClient{stream}
	return x, nil
}

type SimpleService_StreamingBiDirectionalClient interface {
	Send(*SimpleRequest) error
	Recv() (*SimpleResponse, error)
	grpc.ClientStream
}

type simpleServiceStreamingBiDirectionalClient struct {
	grpc.ClientStream
}

func (x *simpleServiceStreamingBiDirectionalClient) Send(m *SimpleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *simpleServiceStreamingBiDirectionalClient) Recv() (*SimpleResponse, error) {
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SimpleServiceServer is the server API for SimpleService service.
// All implementations must embed UnimplementedSimpleServiceServer
// for forward compatibility
type SimpleServiceServer interface {
	// unary RPC
	RPCRequest(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// Server Streaming
	ServerStreaming(*SimpleRequest, SimpleService_ServerStreamingServer) error
	// Client Streaming
	ClientStreaming(SimpleService_ClientStreamingServer) error
	// Bi-Directional Streaming
	StreamingBiDirectional(SimpleService_StreamingBiDirectionalServer) error
	mustEmbedUnimplementedSimpleServiceServer()
}

// UnimplementedSimpleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSimpleServiceServer struct {
}

func (UnimplementedSimpleServiceServer) RPCRequest(context.Context, *SimpleRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RPCRequest not implemented")
}
func (UnimplementedSimpleServiceServer) ServerStreaming(*SimpleRequest, SimpleService_ServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreaming not implemented")
}
func (UnimplementedSimpleServiceServer) ClientStreaming(SimpleService_ClientStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreaming not implemented")
}
func (UnimplementedSimpleServiceServer) StreamingBiDirectional(SimpleService_StreamingBiDirectionalServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingBiDirectional not implemented")
}
func (UnimplementedSimpleServiceServer) mustEmbedUnimplementedSimpleServiceServer() {}

// UnsafeSimpleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleServiceServer will
// result in compilation errors.
type UnsafeSimpleServiceServer interface {
	mustEmbedUnimplementedSimpleServiceServer()
}

func RegisterSimpleServiceServer(s grpc.ServiceRegistrar, srv SimpleServiceServer) {
	s.RegisterService(&SimpleService_ServiceDesc, srv)
}

func _SimpleService_RPCRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServiceServer).RPCRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SimpleService/RPCRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServiceServer).RPCRequest(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleService_ServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SimpleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SimpleServiceServer).ServerStreaming(m, &simpleServiceServerStreamingServer{stream})
}

type SimpleService_ServerStreamingServer interface {
	Send(*SimpleResponse) error
	grpc.ServerStream
}

type simpleServiceServerStreamingServer struct {
	grpc.ServerStream
}

func (x *simpleServiceServerStreamingServer) Send(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SimpleService_ClientStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SimpleServiceServer).ClientStreaming(&simpleServiceClientStreamingServer{stream})
}

type SimpleService_ClientStreamingServer interface {
	SendAndClose(*SimpleResponse) error
	Recv() (*SimpleRequest, error)
	grpc.ServerStream
}

type simpleServiceClientStreamingServer struct {
	grpc.ServerStream
}

func (x *simpleServiceClientStreamingServer) SendAndClose(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *simpleServiceClientStreamingServer) Recv() (*SimpleRequest, error) {
	m := new(SimpleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SimpleService_StreamingBiDirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SimpleServiceServer).StreamingBiDirectional(&simpleServiceStreamingBiDirectionalServer{stream})
}

type SimpleService_StreamingBiDirectionalServer interface {
	Send(*SimpleResponse) error
	Recv() (*SimpleRequest, error)
	grpc.ServerStream
}

type simpleServiceStreamingBiDirectionalServer struct {
	grpc.ServerStream
}

func (x *simpleServiceStreamingBiDirectionalServer) Send(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *simpleServiceStreamingBiDirectionalServer) Recv() (*SimpleRequest, error) {
	m := new(SimpleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SimpleService_ServiceDesc is the grpc.ServiceDesc for SimpleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimpleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SimpleService",
	HandlerType: (*SimpleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RPCRequest",
			Handler:    _SimpleService_RPCRequest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreaming",
			Handler:       _SimpleService_ServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreaming",
			Handler:       _SimpleService_ClientStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamingBiDirectional",
			Handler:       _SimpleService_StreamingBiDirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "simple.proto",
}
