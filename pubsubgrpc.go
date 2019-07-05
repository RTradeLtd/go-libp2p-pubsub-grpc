package pubsubgrpc

import (
	"context"
	"io"

	pb "github.com/RTradeLtd/go-libp2p-pubsub-grpc/pb"
	"github.com/libp2p/go-libp2p-core/host"
	discovery "github.com/libp2p/go-libp2p-discovery"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"go.uber.org/zap"
)

// Service is a libp2p pubsub service component
type Service struct {
	ps *ps.PubSub
	sd *discovery.RoutingDiscovery
	h  host.Host
	l  *zap.SugaredLogger
}

// NewService creates and returns a new libp2p pubsub service
func NewService(
	ps *ps.PubSub,
	sd *discovery.RoutingDiscovery,
	h host.Host,
	l *zap.SugaredLogger,
) Service {
	return Service{
		ps: ps,
		sd: sd,
		h:  h,
		l:  l,
	}
}

// GetTopics is used to return a list of all known topics the pubsub instance is subscribed to.
func (s *Service) GetTopics(ctx context.Context, req *pb.Empty) (*pb.TopicsResponse, error) {
	return &pb.TopicsResponse{Names: s.ps.GetTopics()}, nil
}

// ListPeers is used to return a list of peers subscribed to a given topic or topics
func (s *Service) ListPeers(ctx context.Context, req *pb.ListPeersRequest) (*pb.ListPeersResponse, error) {
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
func (s *Service) Subscribe(req *pb.SubscribeRequest, stream pb.PubSubService_SubscribeServer) error {
	sub, err := s.ps.Subscribe(req.GetTopic())
	if err != nil {
		return err
	}
	// we're joining a pubsub room
	// so we should ensure that we're discovering
	// peers that are also on this room
	if req.GetDiscover() {
		go s.handleDiscover(stream.Context(), req.GetTopic())
	}
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
func (s *Service) Publish(stream pb.PubSubService_PublishServer) error {
	// defer stream closure
	defer stream.SendAndClose(&pb.Empty{})
	var sent = make(map[string]bool)
	for {
		msg, err := stream.Recv()
		if err != nil && err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		// prevent multiple goroutines for the same topic
		// only advertise if specified
		if !sent[msg.GetTopic()] && msg.GetAdvertise() {
			sent[msg.GetTopic()] = true
			go s.handleAnnounce(stream.Context(), msg.GetTopic())
		}
		if err := s.ps.Publish(msg.GetTopic(), msg.GetData()); err != nil {
			return err
		}
	}
}

func (s *Service) handleAnnounce(ctx context.Context, ns string) error {
	_, err := s.sd.Advertise(ctx, ns)
	return err
}

// handleDiscover is used to trigger discovery of
// peers that are part of a particular pubsub room
func (s *Service) handleDiscover(ctx context.Context, ns string) error {
	peerChan, err := s.sd.FindPeers(ctx, ns)
	if err != nil {
		return err
	}
	for peer := range peerChan {
		go s.h.Connect(ctx, peer)
	}
	return nil
}
