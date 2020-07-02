package teststore

import (
	"balance/internal/app/models"
	"context"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[int]*models.User
}

// Create ...
func (r *UserRepository) Create(ctx context.Context, u *models.User) error {
	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(context.Context, string) (*models.User, error) {
	return nil, nil
}
