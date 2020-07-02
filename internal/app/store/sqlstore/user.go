package sqlstore

import (
	"balance/internal/app/models"
	"balance/internal/app/store"
	"context"
	"strings"

	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(ctx context.Context, u *models.User) error {
	// Проверка на валидность всех данных
	if err := u.Validate(); err != nil {
		return errors.New("данные не действительны")
	}
	if u.Password != u.RepeatPassword {
		return errors.New("пароли не идентичны")
	}
	// Проверка на существование такого email
	if _, err := r.FindByEmail(ctx, u.Email); err != nil {
		if err != store.ErrRecordNotFound {
			return err
		}
		// если нет - то создать uuid, захэшировать пароль и создать юзера
		if u.Username == "" {
			u.Username = strings.Split(u.Email, "@")[0]
		}
		u.ID = uuid.New().String()
		if err := u.PasswordHashing(); err != nil {
			return err
		}
		if _, err := r.store.db.Exec(ctx,
			"INSERT INTO users VALUES ($1,$2,$3,$4);",
			u.ID,
			u.Email,
			u.HashPassword,
			u.Username,
		); err != nil {
			return err
		}
		return nil
	}
	return errors.New("пользователь с таким логином существует")
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	u := &models.User{}
	var comment pgtype.Varchar
	if err := r.store.db.QueryRow(ctx,
		"SELECT * FROM users WHERE email=$1;", email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.HashPassword,
		&u.Username,
		&u.IsAdmin,
		&u.Status,
		&u.CreatedAt,
		&comment,
	); err != nil {
		if err == pgx.ErrNoRows {
			err = store.ErrRecordNotFound
		}
		return nil, err
	}

	if comment.Status != pgtype.Null {
		u.Comment = comment.String
	}
	return u, nil
}
