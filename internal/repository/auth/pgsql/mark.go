package pgsql

import (
	"context"
	"fmt"

	"github.com/noskov-sergey/hh-auth/internal/domain"
	"github.com/noskov-sergey/hh-auth/internal/domain/agregates"
)

func (r *Repository) MarkInactive(ctx context.Context, id agregates.ID) error {
	_, err := r.db.Exec(`UPDATE auth SET status = $1 WHERE id = $2`, domain.Inactive, id)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
