package sqlstore

import (
    "github.com/jackc/pgx/v4"
    "balance/internal/app/store"
)

// Store ...
type Store struct {
	db             *pgx.Conn
	userRepository *UserRepository
}

// New ...
func New(db *pgx.Conn) *Store {
	return &Store{db: db}
}

// UserRepository ...
func (s *Store) UserRepository() store.UserRepository {
    if s.userRepository == nil {
        s.userRepository = &UserRepository{store: s}
    }
    return s.userRepository
}