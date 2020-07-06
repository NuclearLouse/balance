package teststore

import (
	"balance/internal/app/models"
	"context"
	"github.com/pkg/errors"
)

// CreateStockDefault ...
func (r *Repository) CreateStockDefault(ctx context.Context, s *models.Stock) error {
	r.stocks[s.ID] = s
	if len(r.stocks) == 0 {
		return errors.New("не смог создать основной склад")
	}
	return nil
}

// CreateStockTemp ...
func (r *Repository) CreateStockTemp(ctx context.Context, s *models.Stock) error {
	return nil
}
