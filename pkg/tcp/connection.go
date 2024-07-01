package tcp

import (
	"log"
	"net"
)

const (
	receiveByteNum = 1024
)

type Connection struct {
	conn net.Conn
}

func NewConnection(
	conn net.Conn,
) *Connection {
	return &Connection{
		conn: conn,
	}
}

func (c Connection) Close() error {
	if c.conn != nil {
		err := c.conn.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c Connection) Read() (string, error) {
	buffer := make([]byte, receiveByteNum)
	n, err := c.conn.Read(buffer)
	if err != nil {
		return "", err
	}
	message := string(buffer[:n])
	log.Printf("read: %s", message)
	return message, nil
}

func (c Connection) Write(message string) error {
	log.Printf("write: %s", message)
	response := []byte(message)
	_, err := c.conn.Write(response)
	if err != nil {
		return err
	}
	return nil
}

func (c Connection) Conn() net.Conn {
	return c.conn
}
