package models

import "time"

// Stock ...
type Stock struct {
	ID        int64
	Owner     string
	Name      string
	Status    bool
	CreatedAt time.Time
	Comment   string
}
