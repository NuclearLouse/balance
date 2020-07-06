package store

import "github.com/pkg/errors"

var (
	// ErrRecordNotFound ...
	ErrRecordNotFound = errors.New("запись не найдена")
)
