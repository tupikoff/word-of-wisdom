package main

import (
	"context"
	"log"

	"github.com/tupikoff/word-of-wisdom/internal/server/delivery"
	"github.com/tupikoff/word-of-wisdom/internal/server/infrastructure"
	"github.com/tupikoff/word-of-wisdom/internal/server/usecase"
)

const defaultPort = 8080
const defaultProtocol = "tcp4"

func main() {
	ctx := context.Background()

	wordsRepository := infrastructure.NewBibleVersesRepository()
	registerRepository := infrastructure.NewRegisterInMemoryRepository()
	protocolService := usecase.NewProtocolService(wordsRepository, registerRepository)
	server := delivery.NewServer(defaultProtocol, defaultPort, protocolService)
	err := server.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// TODO graceful shutdown? (OoS = Out of Scope)
}
