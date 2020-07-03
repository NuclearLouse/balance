package models

import (
	"testing"
)

// TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email:          "user@example.org",
		Password:       "password",
		RepeatPassword: "password",
	}
}

// TestStock ...
func TestStock(t *testing.T) *Stock {
	return &Stock{
		// Owner:   "b626d490-74d3-444d-9bac-6c1b1b950f79",
		Name:    "user_основной",
		Comment: "основной склад",
	}
}
