package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/tupikoff/word-of-wisdom/pkg/hashcash"
	"github.com/tupikoff/word-of-wisdom/pkg/random"
)

const (
	defaultDifficulty            = 3
	defaultChallengeStringLength = 20
	receiveByteNum               = 256
)

type Connection struct {
	conn net.Conn
}

func NewConnection(conn net.Conn) *Connection {
	return &Connection{
		conn: conn,
	}
}

func (c Connection) Handle(context context.Context) {
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

	var response string

	requestData := strings.Split(request, "|")

	switch requestData[0] {
	case "request":
		response = fmt.Sprintf("challenge|%s:%d",
			random.String(defaultChallengeStringLength), defaultDifficulty)
		break
	case "response":
		if len(requestData) > 1 {
			solution := requestData[1]
			hc, err := hashcash.NewFromString(solution)
			if err != nil {
				log.Printf("HashCash read error: %s", err.Error())
				break
			}
			if !hc.IsHashValid() {
				log.Printf("HashCash not valid")
				break
			}
			response = fmt.Sprintf("granted|%s", random.FromSlice(bibleVerses))
		}
	default:
		log.Printf("Unknown command: `%s`", requestData[0])
	}

	if response == "" {
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
