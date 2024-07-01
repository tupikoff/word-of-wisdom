package infrastructure

import (
	"context"
	"sync"

	"github.com/tupikoff/word-of-wisdom/internal/server/domain"
)

type StorageInMemoryRepository struct {
	storage map[string]domain.StorageRecord
	mutex   sync.RWMutex
}

func NewStorageInMemoryRepository() *StorageInMemoryRepository {
	return &StorageInMemoryRepository{
		storage: make(map[string]domain.StorageRecord),
	}
}

func (r *StorageInMemoryRepository) Save(
	ctx context.Context,
	storageRecord domain.StorageRecord,
) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, ok := r.storage[storageRecord.HashString]
	if ok {
		return domain.ErrRecordAlreadyExists
	}

	r.storage[storageRecord.HashString] = storageRecord

	return nil
}
