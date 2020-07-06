package teststore_test

import (
	"balance/internal/app/models"
	"balance/internal/app/store"
	"balance/internal/app/store/teststore"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_CreateUser(t *testing.T) {
	ctx := context.Background()

	s := teststore.New()
	u := models.TestUser(t)

	u.RepeatPassword = "fail"
	err := s.Repository().CreateUser(ctx, u)
	assert.EqualError(t, err, "пароли не идентичны")

	u.RepeatPassword = "password"
	err = s.Repository().CreateUser(ctx, u)
	assert.NoError(t, err)
	assert.NotEmpty(t, u.Status)

	err = s.Repository().CreateUser(ctx, u)
	assert.EqualError(t, err, "пользователь с таким логином существует")
}

func TestRepository_FindUserByEmail(t *testing.T) {
	ctx := context.Background()
	s := teststore.New()

	email := "user@example.org"
	_, err := s.Repository().FindUser(ctx, "email", email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := models.TestUser(t)
	s.Repository().CreateUser(ctx, u)
	tu, err := s.Repository().FindUser(ctx, "email", u.Email)
	assert.NoError(t, err)
	assert.NotEmpty(t, tu.Status)
}

func TestRepository_CreateStockDefault(t *testing.T) {
	ctx := context.Background()
	s := teststore.New()

	stock := models.TestStock(t)

	assert.NoError(t, s.Repository().CreateStockDefault(ctx, stock))
}

func TestRepository_CreateClient(t *testing.T) {
	ctx := context.Background()
	s := teststore.New()

	client := models.TestClient(t)

	assert.NoError(t, s.Repository().CreateClient(ctx, client))
}
