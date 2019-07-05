package pubsubgrpc

import (
	"context"
	"net"

	"sync"

	"github.com/RTradeLtd/go-libp2p-pubsub-grpc/pb"
	"google.golang.org/grpc"
)

// Server is used to run the libp2p pubsub grpc server
// it enables communicating pubsub messages over grpc
type Server struct {
	Service
}

// NewServer is used to intiialize a pubsub grpc server and run it
func NewServer(ctx context.Context, wg *sync.WaitGroup, service Service, protocol, url string, serverOpts ...grpc.ServerOption) error {
	lis, err := net.Listen(protocol, url)
	if err != nil {
		return err
	}
	srv := &Server{service}
	gServer := grpc.NewServer(serverOpts...)
	pb.RegisterPubSubServiceServer(gServer, srv)
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			service.l.Info("shutting server down")
			gServer.GracefulStop()
			return
		}
	}()
	return gServer.Serve(lis)
}
