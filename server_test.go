package libpubsubgrpc

import (
	"context"
	"fmt"
	"testing"
	"time"

	//"time"

	"io"
	"sync"

	pb "github.com/RTradeLtd/go-libp2p-pubsub-grpc/pb"
	tutil "github.com/RTradeLtd/go-libp2p-pubsub-grpc/testutils"
	discovery "github.com/libp2p/go-libp2p-discovery"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/multiformats/go-multiaddr"
)

var (
	serverAddr     = "127.0.0.1:9090"
	serverProtocol = "tcp"
)

func Test_Server(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	logger := tutil.NewLogger(t)
	pk := tutil.NewPrivateKey(t)
	pstore := tutil.NewPeerstore(t)
	dstore := tutil.NewDatastore(t)
	addrs := []multiaddr.Multiaddr{tutil.NewMultiaddr(t)}
	host, dht := tutil.SetupLibp2p(ctx, wg, logger.Desugar(), pk, nil, addrs, pstore, dstore, t)
	pubsub, err := ps.NewGossipSub(ctx, host)
	if err != nil {
		cancel()
		t.Fatal(err)
	}
	tutil.Bootstrap(ctx, logger.Desugar(), dht, host, tutil.DefaultBootstrapPeers())
	go func() {
		if err := NewServer(
			ctx,
			wg,
			pubsub,
			discovery.NewRoutingDiscovery(dht),
			host,
			logger,
			true,
			serverProtocol,
			serverAddr,
		); err != nil {
			t.Fatal(err)
		}
	}()

	client, err := NewClient("", "", "127.0.0.1:9090")
	if err != nil {
		cancel()
		t.Fatal(err)
	}

	subStream, err := client.Subscribe(ctx, &pb.SubscribeRequest{Topic: "hello"})
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
	// test publishing
	pubStream, err := client.Publish(ctx)
	if err != nil {
		cancel()
		t.Fatal(err)
	}
	if err := pubStream.Send(&pb.PublishRequest{Topic: "hello", Data: []byte("world")}); err != nil {
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
	time.Sleep(time.Second * 30)
	// end publishing test
	client.Close()
	cancel()
	wg.Wait()
}
