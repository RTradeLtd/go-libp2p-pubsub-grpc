# go-libp2p-pubsub-grpc (pubsubgrpc)

`pubsubgrpc` provides a LibP2P PubSub framework that can be used as a stand-alone gRPC pubsub server and API, or as a module/component of existing gRPC servers to provide LibP2P PubSub functionality. It uses `proto3` and borrows some ideas from [libp2p/go-libp2p-daemon](https://github.com/libp2p/go-libp2p-daemon/blob/master/pb/p2pd.proto) and [libp2p/go-libp2p-pubsub](https://github.com/libp2p/go-libp2p-pubsub/tree/master/pb).

# usage

There are two ways of using this library:

1) One is to construct a standalone pubsub server with a gRPC api (see `server.go`)
2) Provide pubsub functionality to already existing gRPC servers (see `pubsubgrpc.go`)

If using the standalone pubsub server, there is a golang client in `client.go` that can be used.

# todo:

* once go-libp2p-pubsub has discovery, use it.