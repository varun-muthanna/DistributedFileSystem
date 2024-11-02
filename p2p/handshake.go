package p2p

type HandshakeFunc func (Peer) error 

func NOHandshakeFunc(Peer) error { return nil}
