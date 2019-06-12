package testutils

import (
	"context"
	"fmt"
	"sync"

	config "github.com/ipfs/go-ipfs-config"
	libcore "github.com/libp2p/go-libp2p-core"
	peer "github.com/libp2p/go-libp2p-core/peer"

	host "github.com/libp2p/go-libp2p-core/host"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"go.uber.org/zap"
)

// TemporalBootstrapPeers are the temporal
// production nodes public facing
// adressess. Designed to be used
// as a bootstrapping method
var TemporalBootstrapPeers = []string{
	"/ip4/172.218.49.115/tcp/4002/ipfs/QmPvnFXWAz1eSghXD6JKpHxaGjbVo4VhBXY2wdBxKPbne5",
	"/ip4/172.218.49.115/tcp/4003/ipfs/QmXow5Vu8YXqvabkptQ7HddvNPpbLhXzmmU53yPCM54EQa",
	"/ip4/35.203.44.77/tcp/4001/ipfs/QmUMtzoRfQ6FttA7RygL8jJf7TZJBbdbZqKTmHfU6QC5Jm",
}

// DefaultBootstrapPeers returns the default lsit
// of bootstrap peers used by go-ipfs, updated
// with the Temporal bootstrap nodes
func DefaultBootstrapPeers() []libcore.PeerAddrInfo {
	// conversion copied from go-ipfs
	defaults, _ := config.DefaultBootstrapPeers()
	tPeers, _ := config.ParseBootstrapPeers(TemporalBootstrapPeers)
	defaults = append(defaults, tPeers...)
	pinfos := make(map[peer.ID]*libcore.PeerAddrInfo)
	for _, bootstrap := range defaults {
		pinfo, ok := pinfos[bootstrap.ID]
		if !ok {
			pinfo = new(libcore.PeerAddrInfo)
			pinfos[bootstrap.ID] = pinfo
			pinfo.ID = bootstrap.ID
		}

		pinfo.Addrs = append(pinfo.Addrs, bootstrap.Addrs...)
	}

	var peers []libcore.PeerAddrInfo
	for _, pinfo := range pinfos {
		peers = append(peers, *pinfo)
	}
	return peers
}

// Bootstrap is an optional helper to connect to the given peers and bootstrap
// the Peer DHT (and Bitswap). This is a best-effort function. Errors are only
// logged and a warning is printed when less than half of the given peers
// could be contacted. It is fine to pass a list where some peers will not be
// reachable.
func Bootstrap(ctx context.Context, logger *zap.Logger, dt *dht.IpfsDHT, hst host.Host, peers []libcore.PeerAddrInfo) {
	connected := make(chan struct{})

	var wg sync.WaitGroup
	for _, pinfo := range peers {
		//h.Peerstore().AddAddrs(pinfo.ID, pinfo.Addrs, peerstore.PermanentAddrTTL)
		wg.Add(1)
		go func(pinfo libcore.PeerAddrInfo) {
			defer wg.Done()
			err := hst.Connect(ctx, pinfo)
			if err != nil {
				logger.Error("failed to connect to peer", zap.String("peer.id", pinfo.ID.String()), zap.Error(err))
				return
			}
			logger.Info("successfully conneted to peer", zap.String("peer.id", pinfo.ID.String()))
			connected <- struct{}{}
		}(pinfo)
	}

	go func() {
		wg.Wait()
		close(connected)
	}()

	i := 0
	for range connected {
		i++
	}
	if nPeers := len(peers); i < nPeers/2 {
		logger.Warn(fmt.Sprintf("only connected to %d bootstrap peers out of %d\n", i, nPeers))
	}

	err := dt.Bootstrap(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
}
