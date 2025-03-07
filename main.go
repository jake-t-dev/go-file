package main

import (
	"fmt"
	"log"

	"github.com/jake-t-dev/go-file/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("failed to listen and accept: %v", err)
	}

	fmt.Println("gay")

	select {}
}
