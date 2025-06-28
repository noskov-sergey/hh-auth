package pgsql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/noskov-sergey/hh-auth/internal/domain"
	"github.com/noskov-sergey/hh-auth/internal/domain/agregates"
)

func (r *Repository) GetExpired(ctx context.Context) (*agregates.Authorization, error) {
	query, err := r.db.Prepare(`SELECT id, access_token, expired_at, refresh_token, created_at, status FROM auth WHERE expired_at < $1 AND status = $2 LIMIT 1`)
	if err != nil {
		return nil, fmt.Errorf("query prepare: %w", err)
	}

	var got entityAuthorization

	err = query.QueryRow(time.Now(), domain.Active).Scan(&got.ID, &got.AccessToken, &got.ExpiredAt, &got.RefreshToken, &got.CreateAt, &got.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoAuthorization
		}
		return nil, fmt.Errorf("row scan: %w", err)
	}

	return fromEntityAuthorization(got)
}
