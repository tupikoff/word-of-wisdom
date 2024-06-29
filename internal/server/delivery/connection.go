package delivery

import (
	"context"
	"log"
	"net"
)

const (
	receiveByteNum = 256
)

type Connection struct {
	conn            net.Conn
	protocolService protocolServiceInterface
}

func NewConnection(
	conn net.Conn,
	protocolServiceInterface protocolServiceInterface,
) *Connection {
	return &Connection{
		conn:            conn,
		protocolService: protocolServiceInterface,
	}
}

func (c Connection) Handle(ctx context.Context) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}(c.conn)

	log.Printf("New connection from %v", c.conn.RemoteAddr())
	request, err := c.receive()
	if err != nil {
		log.Printf("Error receive: %v", err)
		return
	}
	log.Printf("Request: %s", request)

	response, err := c.protocolService.Execute(ctx, request)
	if err != nil {
		log.Printf("protocol error: %v", err)
		return
	}

	log.Printf("Response: %s", response)
	err = c.response(response)
	if err != nil {
		log.Printf("Error responce: %v", err)
		return
	}
}

func (c Connection) receive() (string, error) {
	buffer := make([]byte, receiveByteNum)
	n, err := c.conn.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}

func (c Connection) response(message string) error {
	response := []byte(message)
	_, err := c.conn.Write(response)
	if err != nil {
		return err
	}
	return nil
}
