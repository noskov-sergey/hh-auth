package auth

import (
	"context"

	"github.com/noskov-sergey/hh-auth/internal/domain"
	"github.com/noskov-sergey/hh-auth/internal/domain/agregates"
)

type repository interface {
	Create(ctx context.Context, data agregates.Authorization) error
	GetExpired(ctx context.Context) (*agregates.Authorization, error)
	MarkInactive(ctx context.Context, id agregates.ID) error
}

type hh interface {
	RefreshAccessToken(ctx context.Context, rToken domain.RefreshToken, aToken domain.AccessToken) (*agregates.Refresh, error)
}

type authService struct {
	rep repository
	hh  hh
}

func NewAuthService(rep repository, hh hh) *authService {
	return &authService{rep: rep, hh: hh}
}
