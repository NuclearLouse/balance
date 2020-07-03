package sqlstore_test

import (
	"balance/internal/app/models"
	"balance/internal/app/store"
	"balance/internal/app/store/sqlstore"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_CreateUser(t *testing.T) {
	ctx := context.Background()
	db, teardown := sqlstore.TestDB(ctx, t, testDB)
	defer teardown("users")

	s := sqlstore.New(db)

	u := models.TestUser(t)
	assert.NoError(t, s.Repository().CreateUser(ctx, u))
}

func TestRepository_FindUserByEmail(t *testing.T) {
	ctx := context.Background()
	db, teardown := sqlstore.TestDB(ctx, t, testDB)
	defer teardown("users")

	s := sqlstore.New(db)

	email := "fail@mail.org"
	_, err := s.Repository().FindUser(ctx,"email", email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := models.TestUser(t)
	s.Repository().CreateUser(ctx, u)
	tu, err := s.Repository().FindUser(ctx,"email", u.Email)
	assert.NoError(t, err)
	assert.NotEmpty(t, tu.Status)
}

func TestRepository_CreateStockDefault(t *testing.T) {
	ctx := context.Background()
	db, teardown := sqlstore.TestDB(ctx, t, testDB)
	defer teardown("users", "stocks")

	s := sqlstore.New(db)
	u := models.TestUser(t)
	u.ID = "b626d490-74d3-444d-9bac-6c1b1b950f79"
	s.Repository().CreateUser(ctx, u)
	stock := models.TestStock(t)
	stock.Owner = u.ID
	assert.NoError(t, s.Repository().CreateStockDefault(ctx, stock))

}
