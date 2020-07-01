package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// New ...
func New(ctx context.Context, databaseURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
