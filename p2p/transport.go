package p2p

// peer is an interface that represents remote nodes
type Peer interface{}

// transports is anything that handle communication
// between nodes in the network
// It can be used for TCP, UDP, WS...
type Transport interface {
	ListenAndAccept() error
}
