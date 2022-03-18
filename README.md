# gRPC Long-lived Streaming

https://github.com/omri86/longlived-grpc

This repository holds a minimalistic example of a gRPC long lived streaming application.


## Resources

1. [gRPC客户端的那些事儿](https://tonybai.com/2021/09/17/those-things-about-grpc-client/)
2. [blog grpc-long-lived-streaming](https://dev.bitolog.com/grpc-long-lived-streaming/)
3. [gRPC is easy to misconfigure](https://www.evanjones.ca/grpc-is-tricky.html)
   - Client keepalive is dangerous: do not use it
   - Servers cannot return errors larger than 7 kiB

gRPC支持四种通信模式

gRPC支持四种通信模式，它们是（以下四张图截自[《gRPC: Up and Running》一书](https://book.douban.com/subject/34796013/）：

1. 简单RPC(Simple RPC)：最简单的，也是最常用的gRPC通信模式，简单来说就是一请求一应答
   ![image](https://user-images.githubusercontent.com/1940588/158950733-a8bb3f1a-9a8f-4b94-9d81-146157207da6.png)
2. 服务端流RPC(Server-streaming RPC)：一请求，多应答
   ![image](https://user-images.githubusercontent.com/1940588/158950767-c2b54650-fd1d-4329-b0dd-bacd7aadf607.png)
3. 客户端流RPC(Client-streaming RPC)：多请求，一应答
   ![image](https://user-images.githubusercontent.com/1940588/158950809-1f9c66ac-508c-44ec-a082-4e985e541bed.png)
4. 双向流RPC(Bidirectional-Streaming RPC)：多请求，多应答
   ![image](https://user-images.githubusercontent.com/1940588/158950837-b605bdc3-8cca-48b2-9362-21766478a899.png)

## Instructions

1. To compile the proto file, run `make protoc`
2. To build, run `make build`

Note that this was tested on protoc version: `libprotoc 3.17.3` 

## Running the server

```
$ longlived-grpc
2022/03/18 12:17:27 Starting server on address [::]:7070
2022/03/18 12:17:27 Starting data generation
2022/03/18 12:17:28 Received subscribe request from ID: 4
2022/03/18 12:17:29 Received subscribe request from ID: 1
2022/03/18 12:17:29 Received subscribe request from ID: 3
2022/03/18 12:17:29 Received subscribe request from ID: 2
2022/03/18 12:17:30 Received subscribe request from ID: 5
2022/03/18 12:17:32 Received subscribe request from ID: 6
2022/03/18 12:17:33 Client ID 5 has disconnected
2022/03/18 12:17:33 Client ID 6 has disconnected
2022/03/18 12:17:33 Client ID 1 has disconnected
2022/03/18 12:17:33 Client ID 2 has disconnected
2022/03/18 12:17:33 Client ID 4 has disconnected
2022/03/18 12:17:33 Client ID 3 has disconnected
2022/03/18 12:17:33 Failed to send data to client: rpc error: code = Unavailable desc = transport is closing
2022/03/18 12:17:33 Failed to send data to client: rpc error: code = Unavailable desc = transport is closing
2022/03/18 12:17:33 Failed to send data to client: rpc error: code = Unavailable desc = transport is closing
2022/03/18 12:17:33 Failed to send data to client: rpc error: code = Unavailable desc = transport is closing
2022/03/18 12:17:33 Failed to send data to client: rpc error: code = Unavailable desc = transport is closing
2022/03/18 12:17:33 Failed to send data to client: rpc error: code = Unavailable desc = transport is closing
2022/03/18 12:17:35 Received subscribe request from ID: 1
2022/03/18 12:17:37 Received subscribe request from ID: 2
2022/03/18 12:17:39 Received subscribe request from ID: 3
```

## Running the client(s)

The client process emulates several clients (default is 10).

```
$ longlived-grpc -client
2022/03/18 12:17:19 Subscribing client ID: 1
2022/03/18 12:17:20 Client ID 1 got response: "data mock for: 1"
2022/03/18 12:17:21 Client ID 1 got response: "data mock for: 1"
2022/03/18 12:17:21 Subscribing client ID: 2
2022/03/18 12:17:22 Client ID 1 got response: "data mock for: 1"
2022/03/18 12:17:22 Client ID 2 got response: "data mock for: 2"
2022/03/18 12:17:23 Client ID 1 got response: "data mock for: 1"
2022/03/18 12:17:23 Client ID 2 got response: "data mock for: 2"
2022/03/18 12:17:23 Subscribing client ID: 3
2022/03/18 12:17:24 Client ID 1 got response: "data mock for: 1"
2022/03/18 12:17:24 Client ID 2 got response: "data mock for: 2"
2022/03/18 12:17:24 Client ID 3 got response: "data mock for: 3"
2022/03/18 12:17:24 Failed to receive message: rpc error: code = Unavailable desc = error reading from server: EOF
2022/03/18 12:17:24 Failed to receive message: rpc error: code = Unavailable desc = error reading from server: EOF
2022/03/18 12:17:24 Failed to receive message: rpc error: code = Unavailable desc = error reading from server: EOF
2022/03/18 12:17:28 Subscribing client ID: 4
2022/03/18 12:17:28 Client ID 4 got response: "data mock for: 4"
2022/03/18 12:17:29 Client ID 4 got response: "data mock for: 4"
2022/03/18 12:17:29 Subscribing client ID: 3
2022/03/18 12:17:29 Subscribing client ID: 2
2022/03/18 12:17:29 Subscribing client ID: 1
2022/03/18 12:17:30 Subscribing client ID: 5
2022/03/18 12:17:30 Client ID 3 got response: "data mock for: 3"
2022/03/18 12:17:30 Client ID 4 got response: "data mock for: 4"
2022/03/18 12:17:30 Client ID 1 got response: "data mock for: 1"
2022/03/18 12:17:30 Client ID 5 got response: "data mock for: 5"
2022/03/18 12:17:30 Client ID 2 got response: "data mock for: 2"
2022/03/18 12:17:31 Client ID 1 got response: "data mock for: 1"
2022/03/18 12:17:31 Client ID 3 got response: "data mock for: 3"
2022/03/18 12:17:31 Client ID 5 got response: "data mock for: 5"
2022/03/18 12:17:31 Client ID 2 got response: "data mock for: 2"
2022/03/18 12:17:31 Client ID 4 got response: "data mock for: 4"
2022/03/18 12:17:32 Subscribing client ID: 6
2022/03/18 12:17:32 Client ID 3 got response: "data mock for: 3"
2022/03/18 12:17:32 Client ID 1 got response: "data mock for: 1"
2022/03/18 12:17:32 Client ID 5 got response: "data mock for: 5"
2022/03/18 12:17:32 Client ID 2 got response: "data mock for: 2"
2022/03/18 12:17:32 Client ID 4 got response: "data mock for: 4"
2022/03/18 12:17:32 Client ID 6 got response: "data mock for: 6"
```

## multiple servers

1. server: `goreman start`
2. client: `longlived-grpc -client -addr "static::7071,:7072,:7073"`