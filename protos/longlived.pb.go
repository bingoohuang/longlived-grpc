// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: protos/longlived.proto

package protos

import (
	context "context"
	reflect "reflect"
	sync "sync"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_longlived_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_protos_longlived_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_protos_longlived_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_longlived_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_protos_longlived_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_protos_longlived_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_protos_longlived_proto protoreflect.FileDescriptor

var file_protos_longlived_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x6c, 0x69, 0x76,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x22, 0x19, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xaa, 0x01, 0x0a, 0x09,
	0x4c, 0x6f, 0x6e, 0x67, 0x6c, 0x69, 0x76, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x09, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x32, 0x0a,
	0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x35, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x64, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_longlived_proto_rawDescOnce sync.Once
	file_protos_longlived_proto_rawDescData = file_protos_longlived_proto_rawDesc
)

func file_protos_longlived_proto_rawDescGZIP() []byte {
	file_protos_longlived_proto_rawDescOnce.Do(func() {
		file_protos_longlived_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_longlived_proto_rawDescData)
	})
	return file_protos_longlived_proto_rawDescData
}

var (
	file_protos_longlived_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
	file_protos_longlived_proto_goTypes  = []interface{}{
		(*Request)(nil),  // 0: protos.Request
		(*Response)(nil), // 1: protos.Response
	}
)

var file_protos_longlived_proto_depIdxs = []int32{
	0, // 0: protos.Longlived.Subscribe:input_type -> protos.Request
	0, // 1: protos.Longlived.Unsubscribe:input_type -> protos.Request
	0, // 2: protos.Longlived.NotifyReceived:input_type -> protos.Request
	1, // 3: protos.Longlived.Subscribe:output_type -> protos.Response
	1, // 4: protos.Longlived.Unsubscribe:output_type -> protos.Response
	1, // 5: protos.Longlived.NotifyReceived:output_type -> protos.Response
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_longlived_proto_init() }
func file_protos_longlived_proto_init() {
	if File_protos_longlived_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_longlived_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_longlived_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_longlived_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_longlived_proto_goTypes,
		DependencyIndexes: file_protos_longlived_proto_depIdxs,
		MessageInfos:      file_protos_longlived_proto_msgTypes,
	}.Build()
	File_protos_longlived_proto = out.File
	file_protos_longlived_proto_rawDesc = nil
	file_protos_longlived_proto_goTypes = nil
	file_protos_longlived_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConnInterface
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LonglivedClient is the client API for Longlived service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LonglivedClient interface {
	Subscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (Longlived_SubscribeClient, error)
	Unsubscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	NotifyReceived(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type longlivedClient struct {
	cc grpc.ClientConnInterface
}

func NewLonglivedClient(cc grpc.ClientConnInterface) LonglivedClient {
	return &longlivedClient{cc}
}

func (c *longlivedClient) Subscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (Longlived_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Longlived_serviceDesc.Streams[0], "/protos.Longlived/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &longlivedSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Longlived_SubscribeClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type longlivedSubscribeClient struct {
	grpc.ClientStream
}

func (x *longlivedSubscribeClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *longlivedClient) Unsubscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protos.Longlived/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *longlivedClient) NotifyReceived(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protos.Longlived/NotifyReceived", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LonglivedServer is the server API for Longlived service.
type LonglivedServer interface {
	Subscribe(*Request, Longlived_SubscribeServer) error
	Unsubscribe(context.Context, *Request) (*Response, error)
	NotifyReceived(context.Context, *Request) (*Response, error)
}

// UnimplementedLonglivedServer can be embedded to have forward compatible implementations.
type UnimplementedLonglivedServer struct {
}

func (*UnimplementedLonglivedServer) Subscribe(*Request, Longlived_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func (*UnimplementedLonglivedServer) Unsubscribe(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}

func (*UnimplementedLonglivedServer) NotifyReceived(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyReceived not implemented")
}

func RegisterLonglivedServer(s *grpc.Server, srv LonglivedServer) {
	s.RegisterService(&_Longlived_serviceDesc, srv)
}

func _Longlived_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LonglivedServer).Subscribe(m, &longlivedSubscribeServer{stream})
}

type Longlived_SubscribeServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type longlivedSubscribeServer struct {
	grpc.ServerStream
}

func (x *longlivedSubscribeServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _Longlived_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LonglivedServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Longlived/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LonglivedServer).Unsubscribe(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Longlived_NotifyReceived_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LonglivedServer).NotifyReceived(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Longlived/NotifyReceived",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LonglivedServer).NotifyReceived(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Longlived_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Longlived",
	HandlerType: (*LonglivedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unsubscribe",
			Handler:    _Longlived_Unsubscribe_Handler,
		},
		{
			MethodName: "NotifyReceived",
			Handler:    _Longlived_NotifyReceived_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Longlived_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/longlived.proto",
}
