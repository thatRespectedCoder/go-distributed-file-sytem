package p2p

// hand shaker is like , connection has to be there ,
//if there is any error , print them
// else return true for P2P

type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error { return nil }
