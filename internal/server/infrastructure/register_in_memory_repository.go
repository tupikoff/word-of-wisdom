package infrastructure

import (
	"context"
	"sync"

	"github.com/tupikoff/word-of-wisdom/internal/server/domain"
)

type RegisterInMemoryRepository struct {
	storage map[string]domain.RegisterRecord
	mutex   sync.RWMutex
}

func NewRegisterInMemoryRepository() *RegisterInMemoryRepository {
	return &RegisterInMemoryRepository{
		storage: make(map[string]domain.RegisterRecord),
	}
}

func (r *RegisterInMemoryRepository) Save(
	ctx context.Context,
	registerRecord domain.RegisterRecord,
) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, ok := r.storage[registerRecord.HashString]
	if ok {
		return domain.ErrRecordAlreadyExists
	}

	r.storage[registerRecord.HashString] = registerRecord

	return nil
}
