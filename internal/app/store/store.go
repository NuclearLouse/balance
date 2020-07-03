package store

import (
	"balance/internal/app/models"
	"context"
)

// Store ...
type Store interface {
	Repository() Repository
}

// Repository ...
type Repository interface {
	CreateUser(context.Context, *models.User) error
	FindUser(context.Context, string, string) (*models.User, error)
	CreateStockDefault(context.Context, *models.Stock) error
	CreateStockTemp(context.Context, *models.Stock) error
}
