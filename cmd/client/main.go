package main

import (
	"fmt"
	"time"

	"github.com/tupikoff/word-of-wisdom/pkg/hashcash"
)

func main() {
	t := time.Now()
	//fmt.Println(hashcash.New("192.168.0.1", random.String(30), 25).String())
	fmt.Println(hashcash.New("192.168.0.1", "3PsqoXs4WJVARfEYbZCdVuWTePxtR4", 20).String())
	fmt.Printf("%.2f sec\n", time.Since(t).Seconds())
}
