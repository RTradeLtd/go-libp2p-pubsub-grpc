package testutils

import (
	"testing"
	"context"
	"sync"
	"time"

	connmgr "github.com/RTradeLtd/go-libp2p-connmgr"
	datastore "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-ipns"
	"github.com/libp2p/go-libp2p"
	circuit "github.com/libp2p/go-libp2p-circuit"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	host "github.com/libp2p/go-libp2p-core/host"
	peerstore "github.com/libp2p/go-libp2p-core/peerstore"
	ipnet "github.com/libp2p/go-libp2p-core/pnet"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	dopts "github.com/libp2p/go-libp2p-kad-dht/opts"
	pnet "github.com/libp2p/go-libp2p-pnet"
	record "github.com/libp2p/go-libp2p-record"
	routedhost "github.com/libp2p/go-libp2p/p2p/host/routed"
	"github.com/multiformats/go-multiaddr"
	"go.uber.org/zap"
)

// SetupLibp2p is used to initialize
// a libp2p host, and dht on public
// or private networks. To use
// public networks, please provide
// and empty secret
//
// The DHT is not bootstrapped
func SetupLibp2p(
	ctx context.Context,
	wg *sync.WaitGroup,
	logger *zap.Logger,
	hostKey crypto.PrivKey,
	secret []byte,
	listenAddrs []multiaddr.Multiaddr,
	pstore peerstore.Peerstore,
	dstore datastore.Batching,
	t *testing.T,
) (host.Host, *dht.IpfsDHT) {

	var (
		prot ipnet.Protector
		opts []libp2p.Option
		err  error
	)

	// Create protector if we have a secret.
	if secret != nil && len(secret) > 0 {
		var key [32]byte
		copy(key[:], secret)
		prot, err = pnet.NewV1ProtectorFromBytes(&key)
		if err != nil {
			t.Fatal(err)
		}
		opts = append(opts, libp2p.PrivateNetwork(prot))
	}
	opts = append(opts,
		libp2p.Identity(hostKey),
		libp2p.ListenAddrs(listenAddrs...),
		// disabled because it is racy
		// see https://github.com/libp2p/go-nat/issues/11
		//	libp2p.NATPortMap(),
		libp2p.EnableRelay(circuit.OptHop),
		libp2p.ConnectionManager(
			connmgr.NewConnManager(
				ctx,
				wg,
				logger,
				200,
				600,
				time.Minute,
			),
		),
		libp2p.DefaultMuxers,
		libp2p.DefaultTransports,
		libp2p.DefaultSecurity,
	)
	h, err := libp2p.New(ctx, opts...)
	if err != nil {
		t.Fatal(err)
	}

	idht, err := dht.New(ctx, h,
		dopts.Validator(record.NamespacedValidator{
			"pk":   record.PublicKeyValidator{},
			"ipns": ipns.Validator{KeyBook: pstore},
		}),
		dopts.Datastore(dstore),
	)
	if err != nil {
		h.Close()
		t.Fatal(err)
	}
	rHost := routedhost.Wrap(h, idht)
	return rHost, idht
}
