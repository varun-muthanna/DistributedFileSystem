package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct { //remote node over tcp connection
	Conn net.Conn //underlying connection of the peer

	Outbound bool //we dial to peer outbound = true
	// we accept outbound = false
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		Conn:     conn,
		Outbound: outbound,
	}
}

type TCPTransportOpts struct { //fully customizable 
	ListenAddress string
	HandshakeFunc HandshakeFunc
	Dec           DefaultDecode
	Onpeer func(Peer) error
}

type TCPTransport struct {
	TCPTransportOpts
	Listener net.Listener
	rpcch  	chan RPC
	
	Mu   sync.RWMutex
	Peer map[net.Addr]Peer
}

func NewTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch: make(chan RPC),
	}
}

func (p *TCPPeer) Close() error{
	return p.Conn.Close()
}

//implements transport interface , return a read only channel
//reads messages from anothere peer 
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
}

//net.Listen , opens a port a accepts incoming connections
//net.Listener, returned by the net.Listen , blocks and waits for incoming connections , return net.Conn

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.Listener, err = net.Listen("tcp", t.ListenAddress)

	if err != nil {
		fmt.Printf("Error establishing connections %s", err)
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() error { //private

	for {
		con, err := t.Listener.Accept()

		if err != nil {
			fmt.Printf("Error accepting %s", err)
			return err
		}

		go t.handleCon(con)
	}

}

func (t *TCPTransport) handleCon(conn net.Conn) { //read or write
	peer := NewTCPPeer(conn, false)

	defer func(){
		fmt.Println("Dropping peer connection")
		peer.Conn.Close()
	}()

	if t.Onpeer !=nil{
		if e:=t.Onpeer(peer); e!=nil{
			return 
		}
	}

	fmt.Printf("New incoming connection %+v", peer)

	rpc := RPC{}

	for {
		if err := t.Dec.Decode(conn, &rpc); err != nil {
			fmt.Printf("TCP error,%s", err) //if conn closes infinte loop (if no return )
			return 
		}

		rpc.From = conn.RemoteAddr()

		t.rpcch <- rpc

		//fmt.Printf("message %+v\n", rpc)
	}

}
