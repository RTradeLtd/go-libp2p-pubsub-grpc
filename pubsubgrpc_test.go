package pubsubgrpc_test

import (
	"context"
	"fmt"
	"io"
	"sync"
	"testing"
	"time"

	"github.com/RTradeLtd/go-libp2p-pubsub-grpc/pb"

	pubsubgrpc "github.com/RTradeLtd/go-libp2p-pubsub-grpc"
	testutils "github.com/RTradeLtd/go-libp2p-testutils"
	discovery "github.com/libp2p/go-libp2p-discovery"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/multiformats/go-multiaddr"
)

const (
	serverAddr     = "127.0.0.1:9090"
	serverProtocol = "tcp"
)

func TestPubSubService(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	logger := testutils.NewLogger(t)
	pk := testutils.NewPrivateKey(t)
	dstore := testutils.NewDatastore(t)
	pstore := testutils.NewPeerstore(t)
	addrs := []multiaddr.Multiaddr{testutils.NewMultiaddr(t)}
	host, dht := testutils.NewLibp2pHostAndDHT(ctx, t, logger.Desugar(), dstore, pstore, pk, addrs, nil)
	pubsub, err := ps.NewGossipSub(ctx, host)
	if err != nil {
		cancel()
		t.Fatal(err)
	}
	pubsubService := pubsubgrpc.NewService(pubsub, discovery.NewRoutingDiscovery(dht), host, logger)
	if err != nil {
		cancel()
		t.Fatal(err)
	}
	go func() {
		if err := pubsubgrpc.NewServer(
			ctx, wg, pubsubService, "tcp", serverAddr,
		); err != nil {
			cancel()
			t.Fatal(err)
		}
	}()
	// wait for the gRPC server to spin up
	time.Sleep(time.Second * 5)
	client, err := pubsubgrpc.NewClient("", "", serverAddr)
	if err != nil {
		cancel()
		t.Fatal(err)
	}

	subStream, err := client.Subscribe(ctx, &pb.SubscribeRequest{Topic: "hello", Discover: true})
	if err != nil {
		cancel()
		t.Fatal(err)
	}

	defer subStream.CloseSend()
	go func() {
		for {
			msg, err := subStream.Recv()
			if err != nil && err == io.EOF {
				return
			} else if err != nil {
				fmt.Println("got error", err.Error())
				return
			}

			if msg.GetFrom() != nil {
				fmt.Println("from", string(msg.GetFrom()))
			}
			if msg.GetData() != nil {
				fmt.Println("data", string(msg.GetData()))
			}
			if msg.GetSeqno() != nil {
				fmt.Println("seqno", string(msg.GetSeqno()))
			}
			if msg.GetTopicIDs() != nil {
				fmt.Println("topicIDs", msg.GetTopicIDs())
			}
			if msg.GetSignature() != nil {
				fmt.Println("signature", msg.GetSignature())
			}
			if msg.GetKey() != nil {
				fmt.Println("key", msg.GetKey())
			}
		}
	}()

	pubStream, err := client.Publish(ctx)
	if err != nil {
		cancel()
		t.Fatal(err)
	}

	if err := pubStream.Send(&pb.PublishRequest{Topic: "hello", Data: []byte("world"), Advertise: true}); err != nil {
		cancel()
		t.Fatal(err)
	}
	if _, err := pubStream.CloseAndRecv(); err != nil && err != io.EOF {
		cancel()
		t.Fatal(err)
	}
	if _, err := client.ListPeers(ctx, &pb.ListPeersRequest{Topics: []string{"hello"}}); err != nil {
		cancel()
		t.Fatal(err)
	}

	var foundHelloTopic bool
	resp, err := client.GetTopics(ctx, &pb.Empty{})
	if err != nil {
		cancel()
		t.Fatal(err)
	}
	for _, topic := range resp.GetNames() {
		if topic == "hello" {
			foundHelloTopic = true
			break
		}
	}
	if !foundHelloTopic {
		cancel()
		t.Fatal("failed to find hello topic")
	}
	// we need to sleep to give enough time for our goroutine
	// that processes messages to have enough time to pick up the sent messages
	time.Sleep(time.Second * 15)
	// end publishing test
	client.Close()
	cancel()

	wg.Wait()
}
