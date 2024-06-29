package usecase

import (
	"context"

	"github.com/tupikoff/word-of-wisdom/internal/server/domain"
)

type wisdomRepository interface {
	Read() string
}

type registerRepository interface {
	Save(
		ctx context.Context,
		registerRecord domain.RegisterRecord,
	) error
}
