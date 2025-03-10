package p2p

import "net"

// Message is a struct that represents the data that is sent between nodes
type Message struct {
	From    net.Addr
	Payload []byte
}
