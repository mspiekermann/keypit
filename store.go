package main

import "sync"

type Store struct {
	mu   sync.RWMutex
	data map[string]any
}

func NewStore() *Store {
	store := Store{}
	store.data = make(map[string]any)
	return &store
}

func (s *Store) Add(key string, value any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *Store) GetAll() map[string]any {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make(map[string]any, len(s.data))
	for key, value := range s.data {
		result[key] = value
	}

	return result
}

func (s *Store) Get(key string) (any, bool) {
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
