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
	storageRepository := infrastructure.NewStorageInMemoryRepository()
	protocolService := usecase.NewProtocolService(wordsRepository, storageRepository)
	server := delivery.NewServer(defaultProtocol, defaultPort, protocolService)

	err := server.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
