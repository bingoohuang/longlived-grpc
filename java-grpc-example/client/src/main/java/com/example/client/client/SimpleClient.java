/*
 * Copyright 2015 The gRPC Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.example.client.client;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.stub.StreamObserver;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.CountDownLatch;
import javax.annotation.PostConstruct;
import lombok.Getter;
import lombok.SneakyThrows;
import lombok.extern.slf4j.Slf4j;
import lombok.var;
import org.springframework.stereotype.Service;
import simple.testgrpc.Simple.SimpleRequest;
import simple.testgrpc.Simple.SimpleResponse;
import simple.testgrpc.SimpleServiceGrpc;
import simple.testgrpc.SimpleServiceGrpc.SimpleServiceBlockingStub;
import simple.testgrpc.SimpleServiceGrpc.SimpleServiceFutureStub;
import simple.testgrpc.SimpleServiceGrpc.SimpleServiceStub;

/**
 * A simple client that requests a greeting from the .
 */
@Service
@Slf4j
public class SimpleClient {

  private SimpleServiceStub stub;
  private SimpleServiceBlockingStub blockingStub;

  @PostConstruct
  public void init() {
    String user = "world";
    // Access a service running on the local machine on port 50051
    String target = "localhost:50051";
    // Allow passing in the user and target strings as command line arguments

    // Create a communication channel to the server, known as a Channel. Channels are thread-safe
    // and reusable. It is common to create channels at the beginning of your application and reuse
    // them until the application shuts down.
    ManagedChannel channel = ManagedChannelBuilder.forTarget(target)
        // Channels are secure by default (via SSL/TLS). For the example we disable TLS to avoid
        // needing certificates.
        .usePlaintext()
        .build();
    blockingStub = SimpleServiceGrpc.newBlockingStub(channel);
    stub = SimpleServiceGrpc.newStub(channel);

  }


  /**
   * 同步调用：一元RPC
   *
   * @return
   */
  public SimpleResponse rPCRequest(long id, String need) {
    log.info("Will try to SYN {} {}", id, need);
    var response = blockingStub.rPCRequest(newRequest(id, need));
    logResp(response);
    return response;
  }


  @SneakyThrows
  public SimpleResponse clientStreaming(int count) {
    var countDownLatch = new CountDownLatch(1);
    var streamObserver = new SimpleStreamObserver<SimpleResponse>(countDownLatch);
    var streams = stub.clientStreaming(streamObserver);
    for (int i = 1; i <= count; i++) {
      streams.onNext(newRequest(i, String.format("%d/%d", i, count)));
    }
    streams.onCompleted();
    countDownLatch.await();
    return streamObserver.getResp().get(0);
  }

  @SneakyThrows
  public List<SimpleResponse> serverStreaming(long id, String need) {
    log.info("Will try to SYN {} {}", id, need);
    var countDownLatch = new CountDownLatch(1);
    var simpleRequest = newRequest(id, need);
    var streamObserver = new SimpleStreamObserver<SimpleResponse>(countDownLatch);
    stub.serverStreaming(simpleRequest, streamObserver);
    countDownLatch.await();
    return streamObserver.getResp();
  }

  @SneakyThrows
  public List<SimpleResponse> streamingBiDirectional(int count) {
    var countDownLatch = new CountDownLatch(1);
    var streamObserver = new SimpleStreamObserver<SimpleResponse>(countDownLatch);
    var streams = stub.streamingBiDirectional(streamObserver);
    for (int i = 1; i <= count; i++) {
      streams.onNext(newRequest(i, String.format("%d/%d", i, count)));
    }
    streams.onCompleted();
    countDownLatch.await();
    return streamObserver.getResp();
  }


  private void logResp(SimpleResponse response) {
    log.info("Response: " + response.getResponseId() + " -> " + response.getResponse());
  }


  private SimpleRequest newRequest(long id, String need) {
    return SimpleRequest.newBuilder().setRequestId(id).setRequestNeed(need).build();
  }

  /**
   * 处理 stream 类型的结果
   *
   * @param <T> 响应泛型
   */
  static class SimpleStreamObserver<T> implements StreamObserver<T> {

    private CountDownLatch countDownLatch;
    @Getter
    private List<T> resp = new ArrayList<>();

    public SimpleStreamObserver(CountDownLatch countDownLatch) {
      this.countDownLatch = countDownLatch;
    }

    @Override
    public void onNext(T response) {
      log.info("onNext");
      resp.add(response);
    }

    @SneakyThrows
    @Override
    public void onError(Throwable throwable) {
      try {
        log.error("onError", throwable);
        throw throwable;
      } finally {
        countDownLatch.countDown();
      }
    }

    @Override
    public void onCompleted() {
      log.info("onCompleted");
      countDownLatch.countDown();
    }
  }
}
