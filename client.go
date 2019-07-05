package pubsubgrpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"fmt"

	"github.com/RTradeLtd/go-libp2p-pubsub-grpc/pb"
	"github.com/RTradeLtd/grpc/dialer"
)

// Client is used to communicate
// with the gRPC Server service
type Client struct {
	pb.PubSubServiceClient
	conn *grpc.ClientConn
}

// NewClient is used to instantiate a pubsub grpc client
func NewClient(certPath, authKey, url string) (*Client, error) {
	var (
		dialOpts []grpc.DialOption
	)
	if certPath != "" {
		creds, err := credentials.NewClientTLSFromFile(certPath, "")
		if err != nil {
			return nil, fmt.Errorf("could not load tls cert: %s", err)
		}
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(creds))
		if authKey != "" {
			dialOpts = append(dialOpts,
				grpc.WithPerRPCCredentials(dialer.NewCredentials(authKey, true)))
		}
	} else {
		dialOpts = append(dialOpts, grpc.WithInsecure())
		if authKey != "" {
			dialOpts = append(dialOpts,
				grpc.WithPerRPCCredentials(dialer.NewCredentials(authKey, false)))
		}
	}
	conn, err := grpc.Dial(url, dialOpts...)
	if err != nil {
		return nil, err
	}
	return &Client{
		conn:                conn,
		PubSubServiceClient: pb.NewPubSubServiceClient(conn),
	}, nil
}

// Close is used to terminate our connection
// to the grpc server
func (c *Client) Close() error {
	return c.conn.Close()
}
