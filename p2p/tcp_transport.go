package p2p

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// TCPPeer represents the remote node established over TCP
type TCPPeer struct {
	// conn is the underlying connection
	conn net.Conn

	// if we dial and retreieve a conn ==> outbound = true
	// if we accept and retreieve conn ==> outbound = false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	log.Printf("TCP transport listening on port: %s\n", t.ListenAddr)

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		fmt.Printf("new incoming connection %+v\n", conn)

		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP handshake error: %s\n", err)
		return
	}

	// Read loop
	//msg := &Message{}
	msg := make([]byte, 1024)
	for {
		/*if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}*/
		n, err := conn.Read(msg)
		if err != nil {
			fmt.Printf("TCP error: %s\n", err)
		}
		fmt.Printf("TCP received message: %+v\n", msg[:n])
	}
}
