package main

import (
	"context"
	"log"
	"net"

	"github.com/tupikoff/word-of-wisdom/internal/server"
)

func main() {
	port := ":8080"

	ctx := context.Background()

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer func(l net.Listener) {
		err = l.Close()
		if err != nil {
			log.Printf("Error closing listener: %v", err)
		}
	}(l)

	log.Printf("Listening on port %s", port)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			return
		}
		connection := server.NewConnection(conn)
		go connection.Handle(ctx)
	}

	// TODO graceful shutdown? (OoS = Out of Scope)
}
