package simple.testgrpc;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.49.0)",
    comments = "Source: simple.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class SimpleServiceGrpc {

  private SimpleServiceGrpc() {}

  public static final String SERVICE_NAME = "SimpleService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<simple.testgrpc.Simple.SimpleRequest,
      simple.testgrpc.Simple.SimpleResponse> getRPCRequestMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RPCRequest",
      requestType = simple.testgrpc.Simple.SimpleRequest.class,
      responseType = simple.testgrpc.Simple.SimpleResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<simple.testgrpc.Simple.SimpleRequest,
      simple.testgrpc.Simple.SimpleResponse> getRPCRequestMethod() {
    io.grpc.MethodDescriptor<simple.testgrpc.Simple.SimpleRequest, simple.testgrpc.Simple.SimpleResponse> getRPCRequestMethod;
    if ((getRPCRequestMethod = SimpleServiceGrpc.getRPCRequestMethod) == null) {
      synchronized (SimpleServiceGrpc.class) {
        if ((getRPCRequestMethod = SimpleServiceGrpc.getRPCRequestMethod) == null) {
          SimpleServiceGrpc.getRPCRequestMethod = getRPCRequestMethod =
              io.grpc.MethodDescriptor.<simple.testgrpc.Simple.SimpleRequest, simple.testgrpc.Simple.SimpleResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RPCRequest"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  simple.testgrpc.Simple.SimpleRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  simple.testgrpc.Simple.SimpleResponse.getDefaultInstance()))
              .setSchemaDescriptor(new SimpleServiceMethodDescriptorSupplier("RPCRequest"))
              .build();
        }
      }
    }
    return getRPCRequestMethod;
  }

  private static volatile io.grpc.MethodDescriptor<simple.testgrpc.Simple.SimpleRequest,
      simple.testgrpc.Simple.SimpleResponse> getServerStreamingMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ServerStreaming",
      requestType = simple.testgrpc.Simple.SimpleRequest.class,
      responseType = simple.testgrpc.Simple.SimpleResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<simple.testgrpc.Simple.SimpleRequest,
      simple.testgrpc.Simple.SimpleResponse> getServerStreamingMethod() {
    io.grpc.MethodDescriptor<simple.testgrpc.Simple.SimpleRequest, simple.testgrpc.Simple.SimpleResponse> getServerStreamingMethod;
    if ((getServerStreamingMethod = SimpleServiceGrpc.getServerStreamingMethod) == null) {
      synchronized (SimpleServiceGrpc.class) {
        if ((getServerStreamingMethod = SimpleServiceGrpc.getServerStreamingMethod) == null) {
          SimpleServiceGrpc.getServerStreamingMethod = getServerStreamingMethod =
              io.grpc.MethodDescriptor.<simple.testgrpc.Simple.SimpleRequest, simple.testgrpc.Simple.SimpleResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ServerStreaming"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  simple.testgrpc.Simple.SimpleRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  simple.testgrpc.Simple.SimpleResponse.getDefaultInstance()))
              .setSchemaDescriptor(new SimpleServiceMethodDescriptorSupplier("ServerStreaming"))
              .build();
        }
      }
    }
    return getServerStreamingMethod;
  }

  private static volatile io.grpc.MethodDescriptor<simple.testgrpc.Simple.SimpleRequest,
      simple.testgrpc.Simple.SimpleResponse> getClientStreamingMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ClientStreaming",
      requestType = simple.testgrpc.Simple.SimpleRequest.class,
      responseType = simple.testgrpc.Simple.SimpleResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.CLIENT_STREAMING)
  public static io.grpc.MethodDescriptor<simple.testgrpc.Simple.SimpleRequest,
      simple.testgrpc.Simple.SimpleResponse> getClientStreamingMethod() {
    io.grpc.MethodDescriptor<simple.testgrpc.Simple.SimpleRequest, simple.testgrpc.Simple.SimpleResponse> getClientStreamingMethod;
    if ((getClientStreamingMethod = SimpleServiceGrpc.getClientStreamingMethod) == null) {
      synchronized (SimpleServiceGrpc.class) {
        if ((getClientStreamingMethod = SimpleServiceGrpc.getClientStreamingMethod) == null) {
          SimpleServiceGrpc.getClientStreamingMethod = getClientStreamingMethod =
              io.grpc.MethodDescriptor.<simple.testgrpc.Simple.SimpleRequest, simple.testgrpc.Simple.SimpleResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.CLIENT_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ClientStreaming"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  simple.testgrpc.Simple.SimpleRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  simple.testgrpc.Simple.SimpleResponse.getDefaultInstance()))
              .setSchemaDescriptor(new SimpleServiceMethodDescriptorSupplier("ClientStreaming"))
              .build();
        }
      }
    }
    return getClientStreamingMethod;
  }

  private static volatile io.grpc.MethodDescriptor<simple.testgrpc.Simple.SimpleRequest,
      simple.testgrpc.Simple.SimpleResponse> getStreamingBiDirectionalMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "StreamingBiDirectional",
      requestType = simple.testgrpc.Simple.SimpleRequest.class,
      responseType = simple.testgrpc.Simple.SimpleResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.BIDI_STREAMING)
  public static io.grpc.MethodDescriptor<simple.testgrpc.Simple.SimpleRequest,
      simple.testgrpc.Simple.SimpleResponse> getStreamingBiDirectionalMethod() {
    io.grpc.MethodDescriptor<simple.testgrpc.Simple.SimpleRequest, simple.testgrpc.Simple.SimpleResponse> getStreamingBiDirectionalMethod;
    if ((getStreamingBiDirectionalMethod = SimpleServiceGrpc.getStreamingBiDirectionalMethod) == null) {
      synchronized (SimpleServiceGrpc.class) {
        if ((getStreamingBiDirectionalMethod = SimpleServiceGrpc.getStreamingBiDirectionalMethod) == null) {
          SimpleServiceGrpc.getStreamingBiDirectionalMethod = getStreamingBiDirectionalMethod =
              io.grpc.MethodDescriptor.<simple.testgrpc.Simple.SimpleRequest, simple.testgrpc.Simple.SimpleResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.BIDI_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "StreamingBiDirectional"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  simple.testgrpc.Simple.SimpleRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  simple.testgrpc.Simple.SimpleResponse.getDefaultInstance()))
              .setSchemaDescriptor(new SimpleServiceMethodDescriptorSupplier("StreamingBiDirectional"))
              .build();
        }
      }
    }
    return getStreamingBiDirectionalMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static SimpleServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SimpleServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SimpleServiceStub>() {
        @java.lang.Override
        public SimpleServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SimpleServiceStub(channel, callOptions);
        }
      };
    return SimpleServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static SimpleServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SimpleServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SimpleServiceBlockingStub>() {
        @java.lang.Override
        public SimpleServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SimpleServiceBlockingStub(channel, callOptions);
        }
      };
    return SimpleServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static SimpleServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SimpleServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SimpleServiceFutureStub>() {
        @java.lang.Override
        public SimpleServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SimpleServiceFutureStub(channel, callOptions);
        }
      };
    return SimpleServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class SimpleServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * unary RPC
     * </pre>
     */
    public void rPCRequest(simple.testgrpc.Simple.SimpleRequest request,
        io.grpc.stub.StreamObserver<simple.testgrpc.Simple.SimpleResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRPCRequestMethod(), responseObserver);
    }

    /**
     * <pre>
     * Server Streaming
     * </pre>
     */
    public void serverStreaming(simple.testgrpc.Simple.SimpleRequest request,
        io.grpc.stub.StreamObserver<simple.testgrpc.Simple.SimpleResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getServerStreamingMethod(), responseObserver);
    }

    /**
     * <pre>
     * Client Streaming
     * </pre>
     */
    public io.grpc.stub.StreamObserver<simple.testgrpc.Simple.SimpleRequest> clientStreaming(
        io.grpc.stub.StreamObserver<simple.testgrpc.Simple.SimpleResponse> responseObserver) {
      return io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall(getClientStreamingMethod(), responseObserver);
    }

    /**
     * <pre>
     * Bi-Directional Streaming
     * </pre>
     */
    public io.grpc.stub.StreamObserver<simple.testgrpc.Simple.SimpleRequest> streamingBiDirectional(
        io.grpc.stub.StreamObserver<simple.testgrpc.Simple.SimpleResponse> responseObserver) {
      return io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall(getStreamingBiDirectionalMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getRPCRequestMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                simple.testgrpc.Simple.SimpleRequest,
                simple.testgrpc.Simple.SimpleResponse>(
                  this, METHODID_RPCREQUEST)))
          .addMethod(
            getServerStreamingMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                simple.testgrpc.Simple.SimpleRequest,
                simple.testgrpc.Simple.SimpleResponse>(
                  this, METHODID_SERVER_STREAMING)))
          .addMethod(
            getClientStreamingMethod(),
            io.grpc.stub.ServerCalls.asyncClientStreamingCall(
              new MethodHandlers<
                simple.testgrpc.Simple.SimpleRequest,
                simple.testgrpc.Simple.SimpleResponse>(
                  this, METHODID_CLIENT_STREAMING)))
          .addMethod(
            getStreamingBiDirectionalMethod(),
            io.grpc.stub.ServerCalls.asyncBidiStreamingCall(
              new MethodHandlers<
                simple.testgrpc.Simple.SimpleRequest,
                simple.testgrpc.Simple.SimpleResponse>(
                  this, METHODID_STREAMING_BI_DIRECTIONAL)))
          .build();
    }
  }

  /**
   */
  public static final class SimpleServiceStub extends io.grpc.stub.AbstractAsyncStub<SimpleServiceStub> {
    private SimpleServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SimpleServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SimpleServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * unary RPC
     * </pre>
     */
    public void rPCRequest(simple.testgrpc.Simple.SimpleRequest request,
        io.grpc.stub.StreamObserver<simple.testgrpc.Simple.SimpleResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRPCRequestMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Server Streaming
     * </pre>
     */
    public void serverStreaming(simple.testgrpc.Simple.SimpleRequest request,
        io.grpc.stub.StreamObserver<simple.testgrpc.Simple.SimpleResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getServerStreamingMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Client Streaming
     * </pre>
     */
    public io.grpc.stub.StreamObserver<simple.testgrpc.Simple.SimpleRequest> clientStreaming(
        io.grpc.stub.StreamObserver<simple.testgrpc.Simple.SimpleResponse> responseObserver) {
      return io.grpc.stub.ClientCalls.asyncClientStreamingCall(
          getChannel().newCall(getClientStreamingMethod(), getCallOptions()), responseObserver);
    }

    /**
     * <pre>
     * Bi-Directional Streaming
     * </pre>
     */
    public io.grpc.stub.StreamObserver<simple.testgrpc.Simple.SimpleRequest> streamingBiDirectional(
        io.grpc.stub.StreamObserver<simple.testgrpc.Simple.SimpleResponse> responseObserver) {
      return io.grpc.stub.ClientCalls.asyncBidiStreamingCall(
          getChannel().newCall(getStreamingBiDirectionalMethod(), getCallOptions()), responseObserver);
    }
  }

  /**
   */
  public static final class SimpleServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<SimpleServiceBlockingStub> {
    private SimpleServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SimpleServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SimpleServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * unary RPC
     * </pre>
     */
    public simple.testgrpc.Simple.SimpleResponse rPCRequest(simple.testgrpc.Simple.SimpleRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRPCRequestMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Server Streaming
     * </pre>
     */
    public java.util.Iterator<simple.testgrpc.Simple.SimpleResponse> serverStreaming(
        simple.testgrpc.Simple.SimpleRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getServerStreamingMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class SimpleServiceFutureStub extends io.grpc.stub.AbstractFutureStub<SimpleServiceFutureStub> {
    private SimpleServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SimpleServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SimpleServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * unary RPC
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<simple.testgrpc.Simple.SimpleResponse> rPCRequest(
        simple.testgrpc.Simple.SimpleRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRPCRequestMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_RPCREQUEST = 0;
  private static final int METHODID_SERVER_STREAMING = 1;
  private static final int METHODID_CLIENT_STREAMING = 2;
  private static final int METHODID_STREAMING_BI_DIRECTIONAL = 3;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final SimpleServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(SimpleServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_RPCREQUEST:
          serviceImpl.rPCRequest((simple.testgrpc.Simple.SimpleRequest) request,
              (io.grpc.stub.StreamObserver<simple.testgrpc.Simple.SimpleResponse>) responseObserver);
          break;
        case METHODID_SERVER_STREAMING:
          serviceImpl.serverStreaming((simple.testgrpc.Simple.SimpleRequest) request,
              (io.grpc.stub.StreamObserver<simple.testgrpc.Simple.SimpleResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_CLIENT_STREAMING:
          return (io.grpc.stub.StreamObserver<Req>) serviceImpl.clientStreaming(
              (io.grpc.stub.StreamObserver<simple.testgrpc.Simple.SimpleResponse>) responseObserver);
        case METHODID_STREAMING_BI_DIRECTIONAL:
          return (io.grpc.stub.StreamObserver<Req>) serviceImpl.streamingBiDirectional(
              (io.grpc.stub.StreamObserver<simple.testgrpc.Simple.SimpleResponse>) responseObserver);
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class SimpleServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    SimpleServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return simple.testgrpc.Simple.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("SimpleService");
    }
  }

  private static final class SimpleServiceFileDescriptorSupplier
      extends SimpleServiceBaseDescriptorSupplier {
    SimpleServiceFileDescriptorSupplier() {}
  }

  private static final class SimpleServiceMethodDescriptorSupplier
      extends SimpleServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    SimpleServiceMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (SimpleServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new SimpleServiceFileDescriptorSupplier())
              .addMethod(getRPCRequestMethod())
              .addMethod(getServerStreamingMethod())
              .addMethod(getClientStreamingMethod())
              .addMethod(getStreamingBiDirectionalMethod())
              .build();
        }
      }
    }
    return result;
  }
}
