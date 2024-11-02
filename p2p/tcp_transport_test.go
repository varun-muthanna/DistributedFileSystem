package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) { //testing object
	opts := TCPTransportOpts{
		ListenAddress: ":4000",
		HandshakeFunc: NOHandshakeFunc,
		Dec:           DefaultDecode{},
	}

	tr := NewTransport(opts)

	assert.Equal(t, tr.ListenAddress, tr.ListenAddress)
}
