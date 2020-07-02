package teststore

import (
	"balance/internal/app/models"
	"balance/internal/app/store"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[string]*models.User
}

// Create ...
func (r *UserRepository) Create(ctx context.Context, u *models.User) error {
	if err := u.Validate(); err != nil {
		return errors.New("данные не действительны")
	}
	if u.Password != u.RepeatPassword {
		return errors.New("пароли не идентичны")
	}
	if _, err := r.FindByEmail(ctx, u.Email); err != nil {
		if err != store.ErrRecordNotFound {
			return err
		}
		if u.Username == "" {
			u.Username = strings.Split(u.Email, "@")[0]
		}
		u.ID = uuid.New().String()
		if err := u.PasswordHashing(); err != nil {
			return err
		}
		u.Admin = false
		u.Status = true
		u.CreatedAt = time.Now()
		r.users[u.ID] = u
		return nil
	}
	return errors.New("пользователь с таким логином существует")

}

// FindByEmail ...
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, store.ErrRecordNotFound
}
