package p2p

import "net"

// RPC holds the received message from a remote peer
type RPC struct {
	FromAddr net.Addr
	Payload  []byte
}
