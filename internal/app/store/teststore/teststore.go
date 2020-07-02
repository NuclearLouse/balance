package teststore

import (
	"balance/internal/app/models"
	"balance/internal/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New  ...
func New() store.Store {
	return &Store{}
}

// UserRepository ...
func (s *Store) UserRepository() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
			users: make(map[string]*models.User),
		}
	}
	return s.userRepository
}
