package main

type Store struct {
	data map[string]any
}

func NewStore() *Store {
	store := Store{}
	store.data = make(map[string]any)
	return &store
}

func (s *Store) Add(key string, value any) {
	s.data[key] = value
}

func (s *Store) GetAll() map[string]any {
	return s.data
}
