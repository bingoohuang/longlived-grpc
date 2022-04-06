# blog

In this blog post I’ll explore a way to implement gRPC long-lived streaming. Using gRPC is perfect for cloud native applications mainly since it is modern, bandwidth and CPU efficient and low latency which is exactly what distributed systems require.

If you’re reading this I assume you are already familiar with gRPC. But if you still feel like you need an introduction, please leave a comment below and I will put together a gRPC introductory post as well.

## What is considered a “long” RPC session?
A typical RPC is an immediate request-response mechanism. When referring to a long-lived RPC, some of you might have a different timeframe in mind.

A long-lived RPC is usually greater in an order of magnitude from a regular RPC call. For example – it can last minutes, hours, days and so on, depending on your use case.

Before diving into the implementation let’s first consider the use cases for a long-lived RPC stream:

## Long-lived RPC use cases
Now let’s highlight some of the main use cases for a long-lived RPC. Some of these might be different names for the same use case, but I wanted to make sure I cover the most common ones:

- Watch API – where you want to get notifications (similar to the Kubernetes watch API mechanism) when an API object is modified.
- Notifications – when some backend event occurred and you want to receive an immediate notification on that event.
- Subscribers – where several services can subscribe to events and receive those immediately. This use case can also include an unsubscribe functionality.

Note that all of the use cases I’ve mentioned above could have been solved by using polling. But if you’re reading this I guess that polling is something you would want to avoid. By using a long-lived stream you can have immediate response and reduce the latency of events. Think of the RPC usage as a “pipe” – it is set up and ready to handle events at any given time.

## gRPC failure handling
Some of the perks of using gRPC is that it handles some mechanisms that will help you to handle failures. Some of which are:

- Connection Backoff – When we do a connection to a backend which fails, it is typically desirable to not retry immediately (to avoid flooding the network or the server with requests) and instead do some form of exponential backoff.
- Keepalive – The keepalive ping is a way to check if a channel is currently working by sending HTTP2 pings over the transport. It is sent periodically, and if the ping is not acknowledged by the peer within a certain timeout period, the transport is disconnected.

This does not mean to imply that you won’t need to handle cases where the network fails. You should definitely take that into consideration when designing a production grade system.

## What are we going to build?
All of the code used here is available in the following GitHub repository: https://github.com/omri86/longlived-grpc

For the sake of keeping everything simple and focus on how to utilize gRPC – we are going to create a basic application which consists of a single server and multiple clients. I chose to work with a arbitrary number of 10 clients, but as you will see later this scales easily:

![image](https://user-images.githubusercontent.com/1940588/158938052-a162fe15-72dd-4bce-9382-71b76666f380.png)

gRPC client server app example

## Application flow

Below is the general flow I had in mind while building this application. Note that this does not necessarily have to be in any particular order. Meaning, for example, that the clients can start before the server has started.

1. Server starts and waits for clients to subscribe
1. A client starts and sends a subscribe request to the server
1. The server subscribes the client
1. The server sends data periodically to the client

gRPC application flow

![image](https://user-images.githubusercontent.com/1940588/158938064-8e03f12a-923e-412d-b487-9e0468da4e37.png)


A few notes:

- As mentioned above, we are going to have several clients so the server would send data to each of them.
- Sending data could be event-based. I chose to do it periodically for the sake of simplicity.
- Each of the components handles errors gracefully – more on this later.

## Creating the gRPC proto file

See the full proto file here: longlived.proto

As the title of this post indicates, you will use a server streaming RPC. This is declared in the following manner:

```proto
service Longlived {
  rpc Subscribe(Request) returns (stream Response) {}
}
```

The server will handle Subscribe requests (hence the Request argument) and will return a Response. Let’s look at both of these objects:

```proto
message Request {
  int32 id = 1;
}
message Response {
  string data = 1;
}
```

The Request object holds an ID, this would be the client identifier. The Response object holds data – this is the data you will send from the server to the subscribed clients.

The proto file also includes an Unsubscribe unary RPC. This function won’t be used as part of this tutorial but exists to give you and example on how to unsubscribe:

```proto
rpc Unsubscribe(Request) returns (Response) {}
```

## Creating the server
See the full server file here: server.go

First off let’s take a look at the server struct:

```go
type longlivedServer struct {
    protos.UnimplementedLonglivedServer
    subscribers sync.Map // subscribers is a concurrent map that holds mapping from a client ID to it's subscriber
}
type sub struct {
    stream   protos.Longlived_SubscribeServer // stream is the server side of the RPC stream
    finished chan<- bool                      // finished is used to signal closure of a client subscribing goroutine
}
```

- An explanation on line #2 can be found in this thread and this README file.
- The subscribers struct will hold each client which subscribes to your server. It will hold a map from the client ID to a server stream, which you will soon see it’s creation and purpose.
- Since the server can send data to subscribers and receive subscribe requests in parallel, you need to ensure no conflicts, that is why a concurrent map is needed.
- The sub struct will be saved as the map value. It represents a subscriber which holds:
  - The server stream
  - A channel to signal closing the stream

## Server subscribe method
In order to implement the gRPC server interface defined in the proto file you need to implement the following method:

```go
func (s *longlivedServer) Subscribe(request *protos.Request, stream protos.Longlived_SubscribeServer) error
```

This is a method that has a separate context for each incoming subscribe request from a client (a dedicated goroutine). You will receive the client request and a corresponding stream which is used to stream data to the client.

An important note on this function is that the stream will be closed once this function returns. As long as the client is subscribed – this function scope needs to be kept alive.

## Subscribing a client
Since this function will be running as long as the client is subscribed in a separate goroutine, you need a way to access it’s stream in order to send data to the subscribed client.

You will also need a way to signal this goroutine to exit in case the stream is closing.

This is why we need to create a dedicated channel to each client stream (fin channel). The way I chose to implement this mechanism is by holding a map of client IDs to their corresponding channel and stream:

```go
fin := make(chan bool)
// Save the subscriber stream according to the given client ID
s.subscribers.Store(request.Id, sub{stream: stream, finished: fin})
```

And, as explained above, writing to this map (or reading from it) needs to be protected – that is why I’ve used a concurrent map.

Last thing left for you to do is to wait for one of 2 possible events on channels:

- A message that you will send to signal the channel needs to be closed in case of error (you will use the other side of this channel below).
- A message to ctx.Done – this is how the client disconnecting is communicated.

```go
for {
   select {
   case <-fin:
      log.Printf("Closing stream for client ID: %d", request.Id)
   case <- ctx.Done():
      log.Printf("Client ID %d has disconnected", request.Id)
      return nil
   }
}
```

##Creating the client
See the full client code here: client.go

First let’s go over the client struct:

```go
type longlivedClient struct {
   client protos.LonglivedClient // client is the long lived gRPC client
   conn   *grpc.ClientConn       // conn is the client gRPC connection
   id     int32                  // id is the client ID used for subscribing
}
```

- client represents the gRPC client, we will soon initialize it.
- conn will hold the gRPC connection (client <-> server)
- As explained on the server side, clients are subscribing with their unique ID. The id field is the one that holds this ID.

## Client subscribe method
In order to subscribe to server updates, the client must call the gRPC Subscribe() function. This is done as follows:

```go
c.client.Subscribe(context.Background(), &protos.Request{Id: c.id})
```

- The context can be set to carry deadlines, cancellation signals and so on. Since this is out of the scope of this blog post you can read more about it here.
- The second value passed to the server is the client Request which holds the client id.

## Starting the client
The stream is used to stream data from the server to the client after subscribing:

```go
var stream protos.Longlived_SubscribeClient
```

First thing the client should do is subscribe as explained in previous section:

```go
if stream == nil {
   if stream, err = c.subscribe(); err != nil {
      log.Printf("Failed to subscribe: %v", err)
      c.sleep()
      // Retry on failure
      continue
   }
}
```

As you can see, in case of an error we just sleep for a few seconds and then try to resubscribe. This is done to ensure that the client is resilient to server crashes. You need to ensure that the client keeps retrying to connect in case the server is unresponsive. You will see how this works in action later.

Next and last part of the client is to receive data from the stream:

```go
response, err := stream.Recv()
if err != nil {
   log.Printf("Failed to receive message: %v", err)
   // Clearing the stream will force the client to resubscribe on next iteration
   stream = nil
   c.sleep()
   // Retry on failure
   continue
}
log.Printf("Client ID %d got response: %q", c.id, response.Data)
```

## Creating the actual application
Now that the foundations are clear let’s run our application and see how it actually works.

## Server main function
First thing you need to do is to init the server:

```go
lis, err := net.Listen("tcp", "127.0.0.1:7070")
if err != nil {
   log.Fatalf("failed to listen: %v", err)
lis, err := net.Listen("tcp", "127.0.0.1:7070")
if err != nil {
   log.Fatalf("failed to listen: %v", err)
}
grpcServer := grpc.NewServer([]grpc.ServerOption{}...)

server := &longlivedServer{}
```

Next thing is to run a background task (a goroutine) to generate and send some data to the subscribing clients:

```go
// Start sending data to subscribers
go server.mockDataGenerator()
```

This function iterates over the subscribed clients and sends data on their corresponding stream :

```go
func (s *longlivedServer) mockDataGenerator() {
   log.Println("Starting data generation")
   for {
      time.Sleep(time.Second)

      // A list of clients to unsubscribe in case of error
      var unsubscribe []int32

      // Iterate over all subscribers and send data to each client
      s.subscribers.Range(func(k, v interface{}) bool {
         id, ok := k.(int32)
         if !ok {
            log.Printf("Failed to cast subscriber key: %T", k)
            return false
         }
         sub, ok := v.(sub)
         if !ok {
            log.Printf("Failed to cast subscriber value: %T", v)
            return false
         }
         // Send data over the gRPC stream to the client
         if err := sub.stream.Send(&protos.Response{Data: fmt.Sprintf("data mock for: %d", id)}); err != nil {
            log.Printf("Failed to send data to client: %v", err)
            select {
            case sub.finished <- true:
               log.Printf("Unsubscribed client: %d", id)
            default:
               // Default case is to avoid blocking in case client has already unsubscribed
            }
            // In case of error the client would re-subscribe so close the subscriber stream
            unsubscribe = append(unsubscribe, id)
         }
         return true
      })

      // Unsubscribe erroneous client streams
      for _, id := range unsubscribe {
         s.subscribers.Delete(id)
      }
   }
}
```

And all that’s left is to register and start the server:

```go
// Register the server
protos.RegisterLonglivedServer(grpcServer, server)

log.Printf("Starting server on address %s", lis.Addr().String())

// Start listening
if err := grpcServer.Serve(lis); err != nil {
   log.Fatalf("failed to listen: %v", err)
}
```

You can start the server now:

```sh
$ go build server.go
$ ./server
2021/03/04 08:48:09 Starting server on address 127.0.0.1:7070
2021/03/04 08:48:09 Starting data generation
```

## Client(s) main function

For the client side you are going to emulate several clients under the same process. This could easily be done separately:

```go
func main() {
   // Create multiple clients and start receiving data
   var wg sync.WaitGroup

   for i := 1; i <= 10; i++ {
      wg.Add(1)
      client, err := mkLonglivedClient(int32(i))
      if err != nil {
         log.Fatal(err)
      }
      go client.start()
      time.Sleep(time.Second*2)
   }

   // The wait group purpose is to avoid exiting, the clients do not exit
   wg.Wait()
}
```

The server is already up and running so let’s run the clients:

```sh
$ go build client.go
$ ./client
2021/03/04 09:19:29 Subscribing client ID: 1
 2021/03/04 09:19:29 Client ID 1 got response: "data mock for: 1"
 2021/03/04 09:19:30 Client ID 1 got response: "data mock for: 1"
 2021/03/04 09:19:31 Subscribing client ID: 2
 2021/03/04 09:19:31 Client ID 1 got response: "data mock for: 1"
 2021/03/04 09:19:31 Client ID 2 got response: "data mock for: 2"
 2021/03/04 09:19:32 Client ID 1 got response: "data mock for: 1"
 2021/03/04 09:19:32 Client ID 2 got response: "data mock for: 2"
 2021/03/04 09:19:33 Subscribing client ID: 3
 2021/03/04 09:19:33 Client ID 2 got response: "data mock for: 2"
 2021/03/04 09:19:33 Client ID 1 got response: "data mock for: 1"
 2021/03/04 09:19:33 Client ID 3 got response: "data mock for: 3"
 2021/03/04 09:19:34 Client ID 1 got response: "data mock for: 1"
 2021/03/04 09:19:34 Client ID 3 got response: "data mock for: 3"
 2021/03/04 09:19:34 Client ID 2 got response: "data mock for: 2"
 2021/03/04 09:19:35 Subscribing client ID: 4
 2021/03/04 09:19:35 Client ID 2 got response: "data mock for: 2"
 2021/03/04 09:19:35 Client ID 4 got response: "data mock for: 4"
```

As you can see clients are starting and are subscribing and receiving data as expected.

## Error handling

Both the client and the server handle errors on the opposite side by simply retrying to connect.

Let’s test this by stopping the server and viewing the client logs:

```log
2021/03/07 19:38:43 Failed to receive message: rpc error: code = Unavailable desc = transport is closing
2021/03/07 19:38:48 Failed to subscribe: rpc error: code = Unavailable desc = connection error: desc = "transport: Error while dialing dial tcp 127.0.0.1:7070: connect: connection refused"
```

As you can see, each of the clients would first get an RPC error since the transport stream is closing. The consequent error is failing to subscribe since the server did not respond – hence the transport error: connect: connection refused.

Now start the server again, and stop the clients – let’s view the server logs:

```log
2021/03/07 19:43:04 Failed to send data to client: rpc error: code = Unavailable desc = transport is closing
```

Same error as on the client side, but since we delete the subscribed client and retry we won’t get any other errors.

Starting the clients again would resubscribe the clients, and you can see that both the client(s) and the server are working as expected.

## Conclusion

Using gRPC for a long-lived stream might be a bit daunting at first, but as you can see from this example – it doesn’t have to be!

The gRPC project is a perfect fit for cloud native applications and has a great community and documentation.

Feel free to follow me on twitter for regular updates and please do comment below or contact me if you have any questions or comments.

See you on the next one!

## References and credits:

The official gRPC documentation and examples: https://grpc.io/docs/languages/go/basics/

This great talk by Eric Anderson is a place to start when exploring ideas presented in this blog post: https://www.youtube.com/watch?v=Naonb2XD_2Q&t=4s
