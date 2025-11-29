package starage

import (
	"context"
	"sync"
	"time"

	"redis-like-golang/internal/domain/entity"
	"redis-like-golang/internal/domain/repository"
)

type Store struct {
	data        map[string]*entity.Item
	mu          sync.RWMutex
	stopCleanup chan struct{}
}

func NewStore() repository.KeyValuerepository {
	return &Store{
		data:        map[string]*entity.Item{},
		stopCleanup: make(chan struct{}),
	}
}

func (s *Store) Del(ctx context.Context, key string) int {
	if ctx.Err() != nil {
		return 0
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	_, found := s.data[key]
	if found {
		delete(s.data, key)
		return 1
	}

	return 0
}

func (s *Store) Exists(ctx context.Context, key string) bool {
	panic("unimplemented")
}

func (s *Store) Expire(ctx context.Context, key string, seconds int) bool {
	if ctx.Err() != nil {
		return false
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	item, found := s.data[key]
	if !found {
		return false
	}

	expiration := time.Now().Unix() + int64(seconds)
	item.ExpiresAt = &expiration
	return true
}

func (s *Store) Get(ctx context.Context, key string) (string, bool) {
	if ctx.Err() != nil {
		return "", false
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	item, found := s.data[key]
	if !found {
		return "", false
	}

	if item.IsExpired(time.Now().Unix()) {
		return "", false
	}

	return item.Value, true
}

func (s *Store) Keys(ctx context.Context, pattern string) bool {
	panic("unimplemented")
}

func (s *Store) Persist(ctx context.Context, key string) bool {
	panic("unimplemented")
}

func (s *Store) Set(ctx context.Context, key string, value string) {
	if ctx.Err() != nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = &entity.Item{
		Value:     value,
		ExpiresAt: nil,
	}
}

func (s *Store) Size(ctx context.Context) int {
	panic("unimplemented")
}

func (s *Store) StartCleanup(intervalMs int64) {
	panic("unimplemented")
}

func (s *Store) StopClenup() {
	panic("unimplemented")
}

func (s *Store) TTl(ctx context.Context, key string) int64 {
	if ctx.Err() != nil {
		return -1
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	item, found := s.data[key]
	if !found {
		return -1
	}

	if item.ExpiresAt == nil {
		return -1
	}

	now := time.Now().Unix()
	remaining := *item.ExpiresAt - now
	if remaining <= 0 {
		return -2
	}
	return remaining
}
