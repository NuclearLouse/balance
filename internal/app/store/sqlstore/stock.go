package sqlstore

import (
	"balance/internal/app/models"
	"context"
)

// CreateStockDefault ...
func (r *Repository) CreateStockDefault(ctx context.Context, s *models.Stock) error {
	if _, err := r.store.db.Exec(ctx,
		"INSERT INTO stocks (owner,name,comment) VALUES ($1,$2,$3);",
		s.Owner,
		s.Name,
		s.Comment,
	); err != nil {
		return err
	}
	return nil
}

// CreateStockTemp ...
func (r *Repository) CreateStockTemp(ctx context.Context, s *models.Stock) error {
	return nil
}
