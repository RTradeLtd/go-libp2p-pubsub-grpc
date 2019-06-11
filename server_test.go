package libpubsubgrpc

import (
	"context"
	"testing"
	"time"

	"sync"

	tutil "github.com/RTradeLtd/go-libp2p-pubsub-grpc/testutils"
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
	host, _ := tutil.SetupLibp2p(ctx, wg, logger.Desugar(), pk, nil, addrs, pstore, dstore, t)
	pubsub, err := ps.NewGossipSub(ctx, host)
	if err != nil {
		cancel()
		t.Fatal(err)
	}
	go func() {
		if err := NewServer(
			ctx,
			wg,
			pubsub,
			logger,
			true,
			serverProtocol,
			serverAddr,
		); err != nil {
			t.Fatal(err)
		}
	}()
	time.Sleep(time.Second * 10)
	cancel()
	wg.Wait()
}
