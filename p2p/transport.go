package p2p

// Peer represents the operations preformed on the remote nodes
type Peer interface {
	Close() error
}

// Transport handles the communication operations among peers
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
