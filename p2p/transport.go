package p2p

// Peer is an interface to represent the remote node
type Peer interface{}

// Transport is anything that handles communication between nodes
// TCP, UDP, Websockets, etc.
type Transport interface {
	ListenAndAccept() error
}
