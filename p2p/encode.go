package p2p

import (
	"io"
)

type Decoder interface {
	Decode(io.Reader, any) error //any can take int,string,struct .. just an interface represent any type {}, all types implement empty interface
}

type DefaultDecode struct{}

func (d DefaultDecode) Decode(i io.Reader, rpc *RPC) error {

	buff := make([]byte, 1028)

	n, err := i.Read(buff)

	if err != nil {
		return err
	}

	rpc.Payload = buff[:n]

	return nil
}
