package auth

import (
	"context"

	"go.uber.org/zap"

	"github.com/noskov-sergey/hh-auth/internal/domain/agregates"
)

type service interface {
	Refresh(ctx context.Context, authorization agregates.Authorization) (*agregates.Refresh, error)
	GetExpired(ctx context.Context) (*agregates.Authorization, error)
}

type auth struct {
	service service

	log *zap.Logger
}

func NewRefresh(service service, log *zap.Logger) *auth {
	return &auth{service: service, log: log}
}
