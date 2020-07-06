package sqlstore

import (
	"balance/internal/app/models"
	"balance/internal/app/store"
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

// CreateUser ...
func (r *Repository) CreateUser(ctx context.Context, u *models.User) error {
	if err := u.Validate(); err != nil {
		return errors.New("данные не действительны")
	}
	if u.Password != u.RepeatPassword {
		return errors.New("пароли не идентичны")
	}
	if _, err := r.FindUser(ctx, "email",u.Email); err != nil {
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
		
		if _, err := r.store.db.Exec(ctx,
			"INSERT INTO users VALUES ($1,$2,$3,$4);",
			u.ID,
			u.Email,
			u.HashPassword,
			u.Username,
		); err != nil {
			return err
		}
		s := &models.Stock{
			Name:    u.Username + "_основной",
			Owner:   u.ID,
			Comment: "основной склад",
		}
		if err := r.CreateStockDefault(ctx, s); err != nil {
			return err
		}
		
		return nil
	}
	return errors.New("пользователь с таким логином существует")
}

// FindUser ...
func (r *Repository) FindUser(ctx context.Context, field, value string) (*models.User, error) {
	u := &models.User{}
	var comment pgtype.Varchar
	query := fmt.Sprintf("SELECT * FROM users WHERE status=true AND %s=$1;", field)
	if err := r.store.db.QueryRow(ctx, query, value).
	Scan(
		&u.ID,
		&u.Email,
		&u.HashPassword,
		&u.Username,
		&u.Admin,
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

