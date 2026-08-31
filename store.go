package main

import (
	"maps"
	"sync"
	"time"
)

type Item struct {
	Value     any
	ExpiresAt *time.Time
}

type Store struct {
	mu   sync.RWMutex
	data map[string]Item
}

func NewStore() *Store {
	store := Store{}
	store.data = make(map[string]Item)
	return &store
}

func (s *Store) Add(key string, value Item) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *Store) GetAll() map[string]Item {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make(map[string]Item, len(s.data))
	maps.Copy(result, s.data)

	return result
}

func (s *Store) Get(key string) (Item, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, exists := s.data[key]
	return value, exists
}

func (s *Store) CountItems() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

func (s *Store) Delete(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.data[key]; !exists {
		return false
	}

	delete(s.data, key)
	return true
}
