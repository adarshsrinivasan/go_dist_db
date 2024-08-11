package p2p

import (
	"fmt"
	"io"
	"net"
	"sync"

	log "github.com/sirupsen/logrus"
)

// TCPPeer represents remote peer connection
type TCPPeer struct {
	conn net.Conn

	// isOutBound will be true if we dial a conn, false if we accept.
	isOutBound bool
}

func NewTCPPeer(conn net.Conn, isOutBound bool) *TCPPeer {
	return &TCPPeer{
		conn:       conn,
		isOutBound: isOutBound,
	}
}

func (t *TCPPeer) Close() error {
	return t.conn.Close()
}

type TCPTransportOpts struct {
	ListenAddress string
	HandShakeFunc HandShakeFunc
	Decoder       Decoder
	OnNewPeer     func(Peer) error
}

// TCPTransport implements Transport operations
type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rpcChan  chan RPC

	mu    sync.RWMutex // mutex above what it wants to protect
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcChan:          make(chan RPC),
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var (
		err error
	)
	t.listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil
}

// Consume returns read-only channel to consume incoming messages
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcChan
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			log.Errorf("exception while accepting connection: %v", err)
		}
		log.Infof("New incoming connection: %+v", conn)
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	var err error
	defer func() {
		conn.Close()
		log.Errorf("dropping peer connextion. %v", err)
	}()

	peer := NewTCPPeer(conn, true)

	if err := t.HandShakeFunc(peer); err != nil {
		return
	}

	if t.OnNewPeer != nil {
		if err := t.OnNewPeer(peer); err != nil {
			return
		}
	}

	// Read loop
	rpc := RPC{}
	for {
		if err = t.Decoder.Decode(conn, &rpc); err != nil {
			if err == io.EOF {
				err = fmt.Errorf("connection closed")
				return
			}
			log.Errorf("exception while decoding message: %v", err)
			continue
		}
		rpc.FromAddr = conn.RemoteAddr()
		t.rpcChan <- rpc
	}
}
