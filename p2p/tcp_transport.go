package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	mu            sync.RWMutex // for protecting  nodes from concurrent writes
	peers         map[net.Addr]Peer
}

// here if we want to add new TCP connection , return the listenAddress
func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

// this function ListenAndAccept will look for a TCP
// connection request from the listen Address , returns error if it fails
// port blocked , port not accessible , probably I think it will be a service error
// might have crashed

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)

	if err != nil {
		return err
	}
	// start connection Loop for TCP
	go t.startAcceptLoop()
	return nil

}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()

		// this only when the TCP service , port issue are there
		if err != nil {
			fmt.Printf("TCP is not accepting error: %s\n", err)
		}
		// calling handle connection helper function
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	fmt.Printf("new incoming connection: %+v\n", conn)

}

// if the conenction has to be tested
// check tcp_transport_test.go for the custom testing function
// func Test() {
// 	t := NewTCPTransport((":4344"))
// 	// accpet that connection and listen
// }
