package usecase

import (
	"context"

	"github.com/tupikoff/word-of-wisdom/internal/server/domain"
)

type wisdomRepository interface {
	Read() string
}

type storageRepository interface {
	Save(
		ctx context.Context,
		storageRecord domain.StorageRecord,
	) error
}
