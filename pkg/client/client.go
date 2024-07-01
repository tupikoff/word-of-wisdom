package client

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/tupikoff/word-of-wisdom/pkg/hashcash"
	"github.com/tupikoff/word-of-wisdom/pkg/tcp"
)

type Client struct {
	*tcp.Connection

	Address  string
	Protocol string
}

func New(address string, protocol string) *Client {
	return &Client{
		Address:  address,
		Protocol: protocol,
	}
}

func (c *Client) Request() (string, error) {
	err := c.connect()
	if err != nil {
		return "", err
	}
	defer func(c *Client) {
		err := c.Close()
		if err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}(c)

	// 1. request service
	err = c.Write("request")
	if err != nil {
		return "", err
	}
	res, err := c.Read()
	if err != nil {
		return "", err
	}
	// 3. solve
	responses := strings.Split(res, "|")
	if responses[0] != "challenge" {
		return "", fmt.Errorf("invalid response protocol: %s (should be `challenge`)", responses[0])
	}
	payloads := strings.Split(responses[1], ":")
	randString := payloads[0]
	difficulty, err := strconv.Atoi(payloads[1])
	if err != nil {
		return "", errors.New("invalid response payload difficulty")
	}
	resource := c.Conn().LocalAddr().(*net.TCPAddr).IP.String()
	// 4. solve
	hc := hashcash.New(resource, randString, difficulty).String()
	// 5. response
	err = c.Write(fmt.Sprintf("response|%s", hc))
	if err != nil {
		return "", err
	}
	// 7. grant service
	response, err := c.Read()
	if err != nil {
		return "", err
	}
	responses = strings.Split(response, "|")

	return responses[1], nil
}

func (c *Client) connect() error {
	conn, err := net.Dial(c.Protocol, c.Address)
	if err != nil {
		return err
	}
	c.Connection = tcp.NewConnection(conn)

	return nil
}
