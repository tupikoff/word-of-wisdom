package main

import (
	"log"
	"os"
	"time"

	"github.com/tupikoff/word-of-wisdom/pkg/client"
)

func main() {
	serverAddress := os.Getenv("SERVER_URL")
	if serverAddress == "" {
		serverAddress = "localhost:8080"
	}

	c := client.New(serverAddress, "tcp")

	for {
		wow, err := c.Request()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 1)
			continue
		}
		log.Printf("WORDS OF WISDOM: %s\n\n", wow)

		time.Sleep(time.Second * 5)
	}
}
