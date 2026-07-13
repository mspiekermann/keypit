package main

type Store struct {
	data map[string]any
}

func NewStore() *Store {
	store := Store{}
	store.data = make(map[string]any)
	return &store
}
