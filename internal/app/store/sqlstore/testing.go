package sqlstore

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/jackc/pgx/v4"
)

// TestDB ...
func TestDB(ctx context.Context, t *testing.T, testDB string) (*pgx.Conn, func(...string)) {
	t.Helper()
	conn, err := pgx.Connect(ctx, testDB)
	if err != nil {
		t.Fatal(err)
	}
	return conn, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := conn.Exec(ctx, fmt.Sprintf("TRUNCATE %s CASCADE;", strings.Join(tables, ","))); err != nil {
				t.Fatal(err)
			}
		}
		conn.Close(ctx)
	}
}
