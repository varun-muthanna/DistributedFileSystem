package p2p

type Peer interface{  // any remote node, server
	Close() error
}

type Transport interface{ // handles communication between the nodes in the network . (TCP ,UDP , websockets)
	ListenAndAccept() error
	Consume() <-chan RPC
} 