package main

import (
	"fmt"
	"log"

	"github.com/varun-muthanna/filesystem/p2p"
)

func main() {

	opts := p2p.TCPTransportOpts{
		ListenAddress: ":4000",
		HandshakeFunc: p2p.NOHandshakeFunc,
		Dec: p2p.DefaultDecode{},
		Onpeer: func(peer p2p.Peer) error{
			fmt.Println("Logic for peer outside tcptransport")
			peer.Close()
			return nil
		},
	}

	tr := p2p.NewTransport(opts)

	go func(){
		for{
			msg := <-tr.Consume()
			fmt.Printf("%+v\n",msg)
		}
	}()

	err := tr.ListenAndAccept()

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Succesfull")
	select {} //blocking purpose

}
