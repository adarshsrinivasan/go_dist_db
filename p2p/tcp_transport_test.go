package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	tcpTransportOpts := TCPTransportOpts{
		ListenAddress: ":3000",
		HandShakeFunc: NOPHandShakeFunc,
		Decoder:       DefaultDecoder{},
	}
	transport := NewTCPTransport(tcpTransportOpts)
	assert.Equal(t, transport.ListenAddress, tcpTransportOpts.ListenAddress)
	assert.Equal(t, transport.Decoder, tcpTransportOpts.Decoder)

	// Server test
	assert.Nil(t, transport.ListenAndAccept())
}
