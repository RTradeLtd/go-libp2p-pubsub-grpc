package libpubsubgrpc

import (
	"context"
	"net"
	"time"

	"sync"

	"io"

	"github.com/RTradeLtd/go-libp2p-pubsub-grpc/pb"
	discovery "github.com/libp2p/go-libp2p-discovery"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Server is used to run the libp2p pubsub grpc server
// it enables communicating pubsub messages over grpc
type Server struct {
	pb pb.PubSubServiceServer
	ps *ps.PubSub
	sd *discovery.RoutingDiscovery
}

// NewServer is used to intiialize a pubsub grpc server and run it
func NewServer(ctx context.Context, wg *sync.WaitGroup, pubsub *ps.PubSub, sd *discovery.RoutingDiscovery, logger *zap.SugaredLogger, insecure bool, protocol, url string) error {
	lis, err := net.Listen(protocol, url)
	if err != nil {
		return err
	}
	var serverOpts []grpc.ServerOption
	if !insecure {
		serverOpts, err = options("", "", "", logger)
		if err != nil {
			return err
		}
	}
	srv := &Server{ps: pubsub, sd: sd}
	gServer := grpc.NewServer(serverOpts...)
	pb.RegisterPubSubServiceServer(gServer, srv)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				logger.Info("shutting server down")
				gServer.GracefulStop()
				return
			}
		}
	}()
	return gServer.Serve(lis)
}

// GetTopics is used to return a list of all known topics the pubsub instance is subscribed to.
func (s *Server) GetTopics(ctx context.Context, req *pb.Empty) (*pb.TopicsResponse, error) {
	return &pb.TopicsResponse{Names: s.ps.GetTopics()}, nil
}

// ListPeers is used to return a list of peers subscribed to a given topic or topics
func (s *Server) ListPeers(ctx context.Context, req *pb.ListPeersRequest) (*pb.ListPeersResponse, error) {
	var peers []*pb.ListPeersResponse_Peer
	for _, topic := range req.GetTopics() {
		pids := s.ps.ListPeers(topic)
		for _, pid := range pids {
			peers = append(peers, &pb.ListPeersResponse_Peer{Topic: topic, PeerID: pid.String()})
		}
	}
	return &pb.ListPeersResponse{Peers: peers}, nil
}

// Subscribe is used to subscribe to a topic and receive messages
func (s *Server) Subscribe(req *pb.SubscribeRequest, stream pb.PubSubService_SubscribeServer) error {
	sub, err := s.ps.Subscribe(req.GetTopic())
	if err != nil {
		return err
	}
	// we're joining a pubsub room
	// so we should ensure that we're discovering
	// peers that are also on this room
	go s.handleDiscover(req.GetTopic())
	for {
		proto2Msg, err := sub.Next(stream.Context())
		if err != nil {
			return err
		}
		if proto2Msg == nil {
			continue
		}
		// since libp2p-pubsub is using proto2 and we are using proto3
		// we need to copy fields, and format as needed
		proto3Msg := &pb.PubSubMessageResponse{
			From:      []byte(proto2Msg.GetFrom().String()),
			Data:      proto2Msg.GetData(),
			Seqno:     proto2Msg.GetSeqno(),
			TopicIDs:  proto2Msg.GetTopicIDs(),
			Signature: proto2Msg.GetSignature(),
			Key:       proto2Msg.GetKey(),
		}
		if err := stream.Send(proto3Msg); err != nil {
			return err
		}
	}
}

// Publish is used to send a stream of messages to a pubsub topic.
func (s *Server) Publish(stream pb.PubSubService_PublishServer) error {
	// defer stream closure
	defer stream.SendAndClose(&pb.Empty{})
	var sent map[string]bool
	for {
		msg, err := stream.Recv()
		if err != nil && err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		// prevent multiple goroutines for the same topic
		if !sent[msg.GetTopic()] {
			sent[msg.GetTopic()] = true
			go s.handleAnnounce(msg.GetTopic())
		}
		if err := s.ps.Publish(msg.GetTopic(), msg.GetData()); err != nil {
			return err
		}
	}
}

// handleAnnounce is used to handle announcing
// that we are joining a particular pubsub room
// or that we're broadcasting messages for a topic
func (s *Server) handleAnnounce(ns string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	_, err := s.sd.Advertise(ctx, ns)
	return err
}

// handleDiscover is used to trigger discovery of
// peers that are part of a particular pubsub room
func (s *Server) handleDiscover(ns string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	_, err := s.sd.Advertise(ctx, ns)
	return err
}
