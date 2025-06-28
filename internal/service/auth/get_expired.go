package auth

import (
	"context"
	"fmt"

	"github.com/noskov-sergey/hh-auth/internal/domain/agregates"
)

func (a *authService) GetExpired(ctx context.Context) (*agregates.Authorization, error) {
	data, err := a.rep.GetExpired(ctx)
	if err != nil {
		return nil, fmt.Errorf("get expired: %w", err)
	}

	return data, nil
}
