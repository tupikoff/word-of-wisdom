package main

import (
	"fmt"
	"time"

	"github.com/tupikoff/word-of-wisdom/pkg/hashcash"
	"github.com/tupikoff/word-of-wisdom/pkg/random"
)

func main() {
	t := time.Now()
	resource := "192.168.0.1"
	randString := random.String(20)
	difficulty := 22
	fmt.Println(hashcash.New(resource, randString, difficulty).String())
	fmt.Printf("%.2f sec\n", time.Since(t).Seconds())
}
