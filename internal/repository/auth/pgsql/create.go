package pgsql

import (
	"context"
	"fmt"

	"github.com/noskov-sergey/hh-auth/internal/domain/agregates"
)

func (r *Repository) Create(ctx context.Context, data agregates.Authorization) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO auth (access_token, refresh_token, expired_at, status)
		VALUES ($1, $2, $3, $4)`,
		data.AccessToken(), data.RefreshToken(), data.ExpiresAt(), data.Status())
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
