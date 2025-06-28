package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/noskov-sergey/hh-auth/internal/domain"
	"github.com/noskov-sergey/hh-auth/internal/domain/agregates"
)

func (a *authService) Refresh(ctx context.Context, agr agregates.Authorization) (*agregates.Refresh, error) {
	got, err := a.hh.RefreshAccessToken(ctx, agr.RefreshToken(), agr.AccessToken())
	if err != nil {
		return nil, fmt.Errorf("refresh access token: %w", err)
	}

	params := agregates.AuthorizationParams{
		ID:           1,
		AccessToken:  got.AccessToken(),
		Expired:      time.Now().Add(time.Second * time.Duration(got.ExpiresIn())),
		RefreshToken: got.RefreshToken(),
		Status:       domain.Active,
		Created:      time.Now(),
	}

	data, err := agregates.NewAuthorization(params)
	if err != nil {
		return nil, fmt.Errorf("new authorization: %w", err)
	}

	err = a.rep.Create(ctx, *data)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	err = a.rep.MarkInactive(ctx, agr.ID())

	return got, nil
}
