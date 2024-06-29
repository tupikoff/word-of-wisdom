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
		return domain.ErrAlreadyExists
	}

	r.storage[registerRecord.HashString] = registerRecord

	return nil
}

func (r *RegisterInMemoryRepository) Get(
	ctx context.Context,
	hashString string,
) (*domain.RegisterRecord, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	record, ok := r.storage[hashString]
	if !ok {
		return nil, domain.ErrHashStringNotRegistered
	}

	return &record, nil
}
