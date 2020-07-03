package teststore

import (
	"balance/internal/app/models"
	"balance/internal/app/store"
)

// Store ...
type Store struct {
	repository *Repository
}

// Repository ...
type Repository struct {
	store  *Store
	users  map[string]*models.User
	stocks map[int64]*models.Stock
}

// New  ...
func New() store.Store {
	return &Store{}
}

// Repository ...
func (s *Store) Repository() store.Repository {
	if s.repository == nil {
		s.repository = &Repository{
			store:  s,
			users:  make(map[string]*models.User),
			stocks: make(map[int64]*models.Stock),
		}
	}
	return s.repository
}
