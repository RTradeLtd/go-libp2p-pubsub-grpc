package testutils

import (
	"encoding/hex"
	"github.com/ipfs/go-datastore"
	dssync "github.com/ipfs/go-datastore/sync"
	"github.com/ipfs/go-ipfs-keystore"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/libp2p/go-libp2p-peerstore/pstoremem"
	"github.com/multiformats/go-multiaddr"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"math/rand"
	"testing"
)

var (
	// EncodedPK is a hex encoded key
	// to be reused across tests
	EncodedPK = "0801124018c93db89bc9614d463003dab59eb9f8028b27835d4b42abe0b707770cbfc6bd9873de48ab48d753e6be17bc50e821e09f50959da17e45448074fdecccf3e7c0"
)

// NewPrivateKey is used to create a new private key
// for testing purposes
func NewPrivateKey(t *testing.T) crypto.PrivKey {
	pkBytes, err := hex.DecodeString(EncodedPK)
	if err != nil {
		t.Fatal(err)
	}
	pk, err := crypto.UnmarshalPrivateKey(pkBytes)
	if err != nil {
		t.Fatal(err)
	}
	return pk
}

// NewSecret is used to generate a
// secret used to secure private libp2p connections
func NewSecret(t *testing.T) []byte {
	data := make([]byte, 32)
	if _, err := rand.Read(data); err != nil {
		t.Fatal(err)
	}
	return data
}

// NewPeerstore is ued to generate an in-memory peerstore
func NewPeerstore(t *testing.T) peerstore.Peerstore {
	return pstoremem.NewPeerstore()
}

// NewDatastore is used to create a new in memory datastore
func NewDatastore(t *testing.T) datastore.Batching {
	return dssync.MutexWrap(datastore.NewMapDatastore())
}

// NewMultiaddr is used to create a new multiaddress
func NewMultiaddr(t *testing.T) multiaddr.Multiaddr {
	addr, err := multiaddr.NewMultiaddr("/ip4/0.0.0.0/tcp/4005")
	if err != nil {
		t.Fatal(err)
	}
	return addr
}

// NewLogger is used to return a test zap logger
func NewLogger(t *testing.T) *zap.SugaredLogger {
	return zaptest.NewLogger(t).Sugar()
}

// NewKeystore is used to return a new in memory keystore
func NewKeystore(t *testing.T) keystore.Keystore {
	return keystore.NewMemKeystore()
}