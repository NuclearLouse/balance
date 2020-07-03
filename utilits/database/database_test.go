package database

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestDatabase_New(t *testing.T) {
	if err := godotenv.Load("C:\\Users\\android\\go\\balance\\.env");err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	conn, err := New(ctx, os.Getenv("DB_TEST"))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(ctx)

	err = conn.Ping(ctx)
	assert.NoError(t, err)
}