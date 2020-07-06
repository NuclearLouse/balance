package teststore

import (
	"balance/internal/app/models"
	"context"

	"github.com/pkg/errors"
)

// CreateClient ...
func (r *Repository) CreateClient(ctx context.Context, client models.Client) error {
	id := len(client.Name)
	r.clients[id]=client
	if len(r.clients) == 0 {
		return errors.New("не смог создать клиента")
	}
	return nil
}