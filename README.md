# gRPC Long-lived Streaming

https://github.com/omri86/longlived-grpc

This repository holds a minimalistic example of a gRPC long-lived streaming application.

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
2. To install, run `make install`

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

## grpc ui

1. ensure env `GRPC_REFLECTION` has not the value as any of `0`, `off`, `no`
2. install [grpcui](https://github.com/fullstorydev/grpcui)
3. start the server: `longlived-grpc`
4. start the grpcui: `grpcui -plaintext localhost:7070`

## 

1. ensure env `GRPC_CHANNELZ` has not the value as any of `0`, `off`, `no`
2. install channelzcli `go install github.com/kazegusuri/channelzcli@latest`
3. start longlived-grpc: `longlived-grpc`

```sh
$ gurl -pb -r :7080/client/start
{"status":200,"message":"OK","data":{"clientID":"26gv4BBtGfMMqSdeGADwUrwzMUH"}}
$ gurl -pb -r :7080/client/start
{"status":200,"message":"OK","data":{"clientID":"26gv4GDw4ZrCNH2vKtKRw0bNTvv"}}
$ channelzcli -k --addr localhost:7070 list channel
ID	Name                                                                            	State	Channel	SubChannel	Calls	Success	Fail	LastCall
2	127.0.0.1:7070                                                                  	READY	0      	1         	191   	188   	0     	639ms
$ channelzcli -k --addr localhost:7070 list serversocket
ID	ServerID	Name                                    	RemoteName          	Local               	Remote              	Started	Success	Fail	LastStream
6	1       	127.0.0.1:49171 -> 127.0.0.1:7070       	<none>              	[127.0.0.1]:7070	[127.0.0.1]:49171	206   	203   	0     	411ms
23	1       	[::1]:49657 -> [::1]:7070               	<none>              	[::1]:7070      	[::1]:49657     	4     	3     	0     	2ms
$ gurl -pb -r :7080/client/start
{"status":200,"message":"OK","data":{"clientID":"26gv7JBDIXPEp36BtGJNBRfsRKs"}}
$ channelzcli -k --addr localhost:7070 list server
ID	Name	LocalAddr	Calls	Success	Fail	LastCall
1	<none>	[::]:7070   	8     	6     	0     	0ms
$ channelzcli -k --addr localhost:7070 list serversocket
ID	ServerID	Name                                    	RemoteName          	Local               	Remote              	Started	Success	Fail	LastStream
6	1       	127.0.0.1:49661 -> 127.0.0.1:7070       	<none>              	[127.0.0.1]:7070	[127.0.0.1]:49661	10    	9     	0     	642ms
8	1       	[::1]:49673 -> [::1]:7070               	<none>              	[::1]:7070      	[::1]:49673     	4     	3     	0     	1ms
$ channelzcli -k --addr localhost:7070 list channel
ID	Name                                                                            	State	Channel	SubChannel	Calls	Success	Fail	LastCall
2	127.0.0.1:7070                                                                  	READY	0      	1         	24    	23    	0     	739ms
$ channelzcli -k --addr localhost:7070 describe server 1
ID: 	1
Name:
Calls:
  Started:        	49
  Succeeded:      	47
  Failed:         	0
  LastCallStarted:	2022-03-21 10:18:33.439509 +0000 UTC
$ channelzcli -k --addr localhost:7070 describe channel 2
ID:       	2
Name:     	127.0.0.1:7070
State:    	READY
Target:   	127.0.0.1:7070
Calls:
  Started:    	60
  Succeeded:  	59
  Failed:     	0
  LastCallStarted:	2022-03-21 10:18:52.151899 +0000 UTC
Socket:   	<none>
Channels:   	<none>
Subchannels:
  ID	Name	State	Start 	Succeeded	Failed
  3		READY	60    	59      	0
Trace:
  NumEvents:	13
  CreationTimestamp:	2022-03-21 10:17:50.107535 +0000 UTC
  Events
    Severity	Description                                                                     	Timestamp
    INFO	Channel Created                                                                 	2022-03-21 10:17:50.107634 +0000 UTC
    INFO	original dial target is: "127.0.0.1:7070"                                       	2022-03-21 10:17:50.107665 +0000 UTC
    INFO	dial target "127.0.0.1:7070" parse failed: parse "127.0.0.1:7070": first path segment in URL cannot contain colon	2022-03-21 10:17:50.107688 +0000 UTC
    INFO	fallback to scheme "passthrough"                                                	2022-03-21 10:17:50.10769 +0000 UTC
    INFO	parsed dial target is: {Scheme:passthrough Authority: Endpoint:127.0.0.1:7070 URL:{Scheme:passthrough Opaque: User: Host: Path:/127.0.0.1:7070 RawPath: ForceQuery:false RawQuery: Fragment: RawFragment:}}	2022-03-21 10:17:50.107763 +0000 UTC
    INFO	Channel authority set to "127.0.0.1:7070"                                       	2022-03-21 10:17:50.107768 +0000 UTC
    INFO	ccResolverWrapper: sending update to cc: {[{127.0.0.1:7070  <nil> <nil> 0 <nil>}] <nil> <nil>}	2022-03-21 10:17:50.107791 +0000 UTC
    INFO	Resolver state updated: {Addresses:[{Addr:127.0.0.1:7070 ServerName: Attributes:<nil> BalancerAttributes:<nil> Type:0 Metadata:<nil>}] ServiceConfig:<nil> Attributes:<nil>} (resolver returned new addresses)	2022-03-21 10:17:50.107801 +0000 UTC
    INFO	ClientConn switching balancer to "round_robin"                                  	2022-03-21 10:17:50.107809 +0000 UTC
    INFO	Channel switches to new LB policy "round_robin"                                 	2022-03-21 10:17:50.107812 +0000 UTC
    INFO	Subchannel(id:3) created                                                        	2022-03-21 10:17:50.107866 +0000 UTC
    INFO	Channel Connectivity change to CONNECTING                                       	2022-03-21 10:17:50.108088 +0000 UTC
    INFO	Channel Connectivity change to READY                                            	2022-03-21 10:17:50.108585 +0000 UTC
$ channelzcli -k --addr localhost:7070 describe serversocket 6
ID:       	6
Name:     	127.0.0.1:49661 -> 127.0.0.1:7070
Local:    	[127.0.0.1]:7070
Remote:   	[127.0.0.1]:49661
Streams:
  Started:    	90
  Succeeded:  	89
  Failed:     	0
  LastCreated:	2022-03-21 10:19:22.172461 +0000 UTC
Messages:
  Sent:    	178
  Recieved:  	90
  LastSent:	2022-03-21 10:19:22.172537 +0000 UTC
  LastReceived:	2022-03-21 10:19:22.172503 +0000 UTC
Options:
Security:
  Model: none
```