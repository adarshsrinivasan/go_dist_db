package main

import (
	"github.com/adarshsrinivasan/go_dist_db/p2p"
	log "github.com/sirupsen/logrus"
)

func main() {
	onNewPeer := func(p2p.Peer) error {
		log.Infof("doing something with new connection")
		return nil
	}
	transportOpts := p2p.TCPTransportOpts{
		ListenAddress: ":3000",
		HandShakeFunc: p2p.NOPHandShakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnNewPeer:     onNewPeer,
	}
	transport := p2p.NewTCPTransport(transportOpts)

	go func() {
		consumeChan := transport.Consume()
		for {
			rpc := <-consumeChan
			log.Infof("Received message: %+v", rpc)
		}
	}()

	if err := transport.ListenAndAccept(); err != nil {
		log.Fatalf("exception while listening: %v", err)
	}
	select {}
}
