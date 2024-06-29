package delivery

import (
	"context"
	"fmt"
	"log"
	"net"
)

type Server struct {
	Protocol string
	Port     int

	protocolService protocolServiceInterface
}

func NewServer(
	protocol string,
	port int,
	protocolService protocolServiceInterface,
) *Server {
	return &Server{
		Protocol:        protocol,
		Port:            port,
		protocolService: protocolService,
	}
}

func (s *Server) Start(ctx context.Context) error {
	l, err := net.Listen(s.Protocol, fmt.Sprintf(":%d", s.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer func(l net.Listener) {
		err = l.Close()
		if err != nil {
			log.Printf("Error closing listener: %v", err)
		}
	}(l)

	log.Printf("Listening on port %d", s.Port)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			return err
		}
		connection := NewConnection(conn, s.protocolService)
		go connection.Handle(ctx)
	}
}
