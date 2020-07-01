package sqlstore

import (
	"balance/internal/app/models"
	"context"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(ctx context.Context, u *models.User) error {
    // Если такого юзера нет, то создать
	if _, err := r.store.db.Exec(ctx, ""); err != nil {
		return err
	}
	return nil
}
