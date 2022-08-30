package com.example.server.impl;

import io.grpc.ServerBuilder;
import io.grpc.protobuf.services.ProtoReflectionService;
import java.io.IOException;
import java.util.concurrent.TimeUnit;
import lombok.extern.slf4j.Slf4j;

/**
 * 服务端
 *
 * @author label
 */
@Slf4j
public class Server {

  private io.grpc.Server grpcServer;

  public void start() throws IOException {
    /* The port on which the server should run */
    int port = 50051;
    grpcServer = ServerBuilder.forPort(port)
        .addService(new SimpleImpl())
        .addService(ProtoReflectionService.newInstance())
        .build().start();
    log.info("Server started, listening on " + port);
    registerShutdownHook();
  }

  private void registerShutdownHook() {
    Runtime.getRuntime().addShutdownHook(new Thread(() -> {
      // Use stderr here since the logger may have been reset by its JVM shutdown hook.
      log.info("*** shutting down gRPC server since JVM is shutting down");
      try {
        Server.this.stop();
      } catch (InterruptedException e) {
        log.error("Interrupted!", e);
        // Restore interrupted state...
        Thread.currentThread().interrupt();
      }
      log.info("*** server shut down");
    }));
  }

  private void stop() throws InterruptedException {
    if (grpcServer != null) {
      grpcServer.shutdown().awaitTermination(30, TimeUnit.SECONDS);
    }
  }

  /**
   * Await termination on the main thread since the grpc library uses daemon threads.
   */
  public void blockUntilShutdown() throws InterruptedException {
    if (grpcServer != null) {
      grpcServer.awaitTermination();
    }
  }
}
