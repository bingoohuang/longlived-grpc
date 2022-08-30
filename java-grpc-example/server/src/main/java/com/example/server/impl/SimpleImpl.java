package com.example.server.impl;

import io.grpc.stub.StreamObserver;
import java.util.HashMap;
import java.util.Map;
import java.util.Random;
import simple.testgrpc.Simple.SimpleRequest;
import simple.testgrpc.Simple.SimpleResponse;
import simple.testgrpc.SimpleServiceGrpc;

public class SimpleImpl extends SimpleServiceGrpc.SimpleServiceImplBase {

  /**
   * 简单RPC(Simple RPC)：最简单的，也是最常用的 gRPC 通信模式，简单来说就是一请求一应答
   *
   * @param req              请求
   * @param responseObserver 响应
   */
  @Override
  public void rPCRequest(SimpleRequest req, StreamObserver<SimpleResponse> responseObserver) {
    logRequest(req);
    SimpleResponse resp = SimpleResponse.newBuilder()
        .setResponse("ACK :" + req.getRequestNeed())
        .setResponseId(req.getRequestId()).build();
    responseObserver.onNext(resp);
    responseObserver.onCompleted();
  }


  /**
   * 服务端流RPC(Server-streaming RPC)：一请求，多应答
   * <p/>
   *
   * @param req
   * @param responseObserver
   */
  @Override
  public void serverStreaming(SimpleRequest req, StreamObserver<SimpleResponse> responseObserver) {
    logRequest(req);
    // 模拟一请求，多应答
    // 模拟又不确定个数的结果集，每次只响应 1 个结果
    Random random = new Random();
    for (int i = 0; i < new Random().nextInt(10); i++) {
      SimpleResponse resp = SimpleResponse.newBuilder()
          .setResponse(String.format("ACK : %s[%d]", req.getRequestNeed(), i))
          .setResponseId(req.getRequestId()).build();
      responseObserver.onNext(resp);
    }
    responseObserver.onCompleted();
  }


  @Override
  public StreamObserver<SimpleRequest> clientStreaming(
      StreamObserver<SimpleResponse> responseObserver) {
    return new StreamObserver<SimpleRequest>() {
      private Map<Long, String> map = new HashMap<>();

      /**
       * 模拟多次请求
       * @param req
       */
      @Override
      public void onNext(SimpleRequest req) {
        logRequest(req);
        map.put(req.getRequestId(), req.getRequestNeed());
      }

      @Override
      public void onError(Throwable throwable) {
        responseObserver.onError(throwable);
      }

      /**
       * 模拟一次响应，响应计算结果：size 和 values
       */
      @Override
      public void onCompleted() {
        SimpleResponse resp = SimpleResponse.newBuilder().setResponse("ACK :" + map.values())
            .setResponseId(map.size()).build();
        responseObserver.onNext(resp);
        responseObserver.onCompleted();
      }
    };
  }


  @Override
  public StreamObserver<SimpleRequest> streamingBiDirectional(
      StreamObserver<SimpleResponse> responseObserver) {

    return new StreamObserver<SimpleRequest>() {
      /**
       * 模拟多次请求
       * @param req
       */
      @Override
      public void onNext(SimpleRequest req) {
        logRequest(req);
        // 模拟对应单次请求的响应（一个或多个）
        for (int i = 0; i < new Random().nextInt(10); i++) {
          SimpleResponse resp = SimpleResponse.newBuilder()
              .setResponse(String.format("ACK : 请求原文 %s 响应添加[%d]", req.getRequestNeed(), i))
              .setResponseId(req.getRequestId()).build();
          responseObserver.onNext(resp);
        }
      }

      @Override
      public void onError(Throwable throwable) {
        responseObserver.onError(throwable);
      }

      /**
       * 模拟一次响应，响应计算结果：size 和 values
       */
      @Override
      public void onCompleted() {
        // 请求处理完了，响应可以结束或者不结束，这里处理完请求就结束
        responseObserver.onCompleted();
      }
    };
  }

  private void logRequest(SimpleRequest req) {
    System.out.println(String.format("Received req as id: %s need: %s",
        req.getRequestId(),
        req.getRequestNeed()));
  }
}
