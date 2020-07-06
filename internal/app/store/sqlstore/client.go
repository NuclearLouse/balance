package sqlstore

import (
	"balance/internal/app/models"
	"context"
)

// CreateClient ...
func (r *Repository) CreateClient(ctx context.Context, c models.Client) error {
	_, err := r.store.db.Exec(ctx,
		`INSERT INTO clients (name,"user",type,markup,status,comment) VALUES ($1,$2,$3,$4,$5,$6);`,
		c.Name,
		c.User,
		c.Type,
		c.Markup,
		c.Status,
		c.Comment,
	)
	return err
}
