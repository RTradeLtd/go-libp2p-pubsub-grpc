# go-libp2p-pubsub-grpc

gRPC libp2p pubsub API that uses `proto3`.

Some of the code is borrowed from [libp2p/go-libp2p-daemon](https://github.com/libp2p/go-libp2p-daemon/blob/master/pb/p2pd.proto) and [libp2p/go-libp2p-pubsub](https://github.com/libp2p/go-libp2p-pubsub/tree/master/pb) modified to use proto3, and provide an API over gRPC.

# usage

There's a few ways to use this repository, one is as a stand-alone client+server pubsub system using the `server.go` and `client.go` files. Alternatively you may use the compiled protobufs in `pb` to implement the API capabilities without your own systems.

To do this you would import `"github.com/RTradeLtd/go-libp2p-pubsub-grpc/pb` into your code if using Golang, or compile the `.proto` file for whatever language you want.

The server expects you to provide the given libp2p pubsub setup. This allows you to leverage whatever underlying pubsub algorithm (floodsub, gossipsub, etc..) you want, as long as it conforms to the pubsub interface specified in `libp2p/go-libp2p-pubsub`. 


# todo:

* once go-libp2p-pubsub has discovery, use it.