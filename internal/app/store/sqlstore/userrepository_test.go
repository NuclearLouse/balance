package sqlstore_test

import (
	"balance/internal/app/models"
	"balance/internal/app/store"
	"balance/internal/app/store/sqlstore"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	ctx := context.Background()
	db, teardown := sqlstore.TestDB(ctx, t, testDB)
	defer teardown("users")

	s := sqlstore.New(db)

	u := models.TestUser(t)
	assert.NoError(t, s.UserRepository().Create(ctx, u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	ctx := context.Background()
	db, teardown := sqlstore.TestDB(ctx, t, testDB)
	defer teardown("users")

	s := sqlstore.New(db)

	email := "fail@mail.org"
	_, err := s.UserRepository().FindByEmail(ctx, email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := models.TestUser(t)
	s.UserRepository().Create(ctx, u)
	tu, err := s.UserRepository().FindByEmail(ctx, u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, tu)
}
