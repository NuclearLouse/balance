package sqlstore_test

import (
	"os"
	"testing"
)

var testDB string

// TestMain ...
func TestMain(m *testing.M) {
	var ok bool
	testDB, ok = os.LookupEnv("DB_TEST")
	if !ok {
		testDB = "postgres://postgres:postgres@localhost/balance?sslmode=disable"
	}
	os.Exit(m.Run())
}