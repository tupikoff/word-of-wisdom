package main

import (
	"log"
	"time"

	"github.com/tupikoff/word-of-wisdom/pkg/client"
)

func main() {
	c := client.New("localhost:8080", "tcp")

	for {
		wow, err := c.Request()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("WORDS OF WISDOM: %s", wow)

		time.Sleep(time.Second * 5)
	}
}
