package store

import (
	"balance/internal/app/models"
	"context"
)

// Store ...
type Store interface {
	UserRepository() UserRepository
}

// UserRepository ...
type UserRepository interface {
	Create(context.Context, *models.User) error
	FindByEmail(context.Context, string) (*models.User, error)
}
