package teststore_test

import (
	"balance/internal/app/models"
	"balance/internal/app/store"
	"balance/internal/app/store/teststore"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	ctx := context.Background()

	s := teststore.New()
	u := models.TestUser(t)

	u.RepeatPassword = "fail"
	err := s.UserRepository().Create(ctx, u)
	assert.EqualError(t, err, "пароли не идентичны")

	u.RepeatPassword = "password"
	err = s.UserRepository().Create(ctx, u)
	assert.NoError(t, err)
	assert.NotEmpty(t, u.Status)

	err = s.UserRepository().Create(ctx, u)
	assert.EqualError(t, err, "пользователь с таким логином существует")
}

func TestUserRepository_FindByEmail(t *testing.T) {
	ctx := context.Background()
	s := teststore.New()

	email := "user@example.org"
	_, err := s.UserRepository().FindByEmail(ctx, email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := models.TestUser(t)
	s.UserRepository().Create(ctx, u)
	tu, err := s.UserRepository().FindByEmail(ctx, u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, tu)

}
