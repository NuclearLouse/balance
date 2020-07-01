package models

import "time"

// User ...
type User struct {
	ID        int
	Email     string
	Password  string
	IsAdmin   bool
	Status    bool
	CreatedAt time.Time
	Comment   string
}

