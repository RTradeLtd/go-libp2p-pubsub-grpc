package testutils

import (
	"testing"
	"sync"
	"context"
	"github.com/multiformats/go-multiaddr"
)

func Test_NewPrivateKey(t *testing.T) {
	if pk := NewPrivateKey(t); pk == nil {
		t.Fatal("should not be nil")
	}
}

func Test_NewSecret(t *testing.T) {
	if secret := NewSecret(t); secret == nil {
		t.Fatal("should not be nil")
	}
}

func Test_NewPeerstore(t *testing.T) {
	if peerstore := NewPeerstore(t); peerstore == nil {
		t.Fatal("should not be nil")
	}
}

func Test_NewDatastore(t *testing.T) {
	if datastore := NewDatastore(t); datastore == nil {
		t.Fatal("should not be nil")
	}
}

func Test_NewMultiaddr(t *testing.T) {
	if addr := NewMultiaddr(t); addr == nil {
		t.Fatal("should not be nil")
	}
}

func Test_NewKeystore(t *testing.T) {
	if keystore := NewKeystore(t); keystore == nil {
		t.Fatal("should not be nil")
	}
}


func Test_SetupLibp2p(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	logger := NewLogger(t)
	pk := NewPrivateKey(t)
	pstore := NewPeerstore(t)
	dstore := NewDatastore(t)
	addrs := []multiaddr.Multiaddr{NewMultiaddr(t)}
	host, dht := SetupLibp2p(ctx, wg, logger.Desugar(), pk, nil, addrs, pstore, dstore, t)
	host.Close()
	dht.Close()
	cancel()
	wg.Wait()
}