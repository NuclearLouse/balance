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

// CreateUser ...
func (r *Repository) CreateUser(ctx context.Context, u *models.User) error {
	if err := u.Validate(); err != nil {
		return errors.New("данные не действительны")
	}
	if u.Password != u.RepeatPassword {
		return errors.New("пароли не идентичны")
	}
	if _, err := r.FindUser(ctx, "email", u.Email); err != nil {
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
		//TODO: тут надо создать склад по дефолту
		
		u.Admin = false
		u.Status = true
		u.CreatedAt = time.Now()
		r.users[u.ID] = u

		s := &models.Stock{
			ID: time.Now().Unix(),
			Owner: u.ID,
			Name: u.Username+"_основной",
			CreatedAt: time.Now(),
			Comment: "основной склад",
			Status: true,
		}
		if err := r.CreateStockDefault(ctx, s); err != nil {
			return err
		}
		return nil
	}
	return errors.New("пользователь с таким логином существует")

}

// FindUser ...
func (r *Repository) FindUser(ctx context.Context, key, value string) (*models.User, error) {
	for _, u := range r.users {
		if u.Email == value {
			return u, nil
		}
	}
	return nil, store.ErrRecordNotFound
}
