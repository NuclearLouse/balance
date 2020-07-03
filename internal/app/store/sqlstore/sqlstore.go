package sqlstore

import (
    "github.com/jackc/pgx/v4"
    "balance/internal/app/store"
)

// Store ...
type Store struct {
	db             *pgx.Conn
	repository *Repository
}

// Repository ...
type Repository struct {
	store *Store
}

// New ...
func New(db *pgx.Conn) *Store {
	return &Store{db: db}
}

// Repository ...
func (s *Store) Repository() store.Repository {
	if s.repository == nil {
		s.repository = &Repository{store: s}
	}
	return s.repository
}
