package auth

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"

	"github.com/noskov-sergey/hh-auth/internal/infrastructure/hh"
	"github.com/noskov-sergey/hh-auth/internal/repository/auth/pgsql"
)

func (a *auth) Refresh(ctx context.Context) {
	log := a.log.Named("controller.refresh")
	for {
		expired, err := a.service.GetExpired(ctx)
		if err != nil {
			if !errors.Is(pgsql.ErrNoAuthorization, err) {
				log.Error("get expired", zap.Error(err))
				time.Sleep(60 * time.Second)
			}
			continue
		}

		log.Info("got expired token, start refresh",
			zap.String("access token", expired.AccessToken().String()),
			zap.Time("expired at", expired.ExpiresAt()),
			zap.String("refresh token", expired.RefreshToken().String()),
			zap.String("status", expired.Status().String()))

		refreshed, err := a.service.Refresh(ctx, *expired)
		if err != nil {
			if !errors.Is(hh.ErrTokenNotExpired, err) {
				log.Warn("token not expired", zap.Error(err),
					zap.String("refresh token", expired.RefreshToken().String()),
					zap.String("access token", expired.AccessToken().String()))
				time.Sleep(60 * time.Second)
				continue
			}
			log.Error("refresh token",
				zap.Error(err),
				zap.Int32("id token", int32(expired.ID())),
				zap.String("refresh token", expired.RefreshToken().String()),
				zap.String("access token", expired.AccessToken().String()))
			time.Sleep(60 * time.Second)
		}

		log.Info("saved refreshed token",
			zap.String("access token", refreshed.AccessToken().String()),
			zap.Int64("expired in", int64(refreshed.ExpiresIn())),
			zap.String("refresh token", refreshed.RefreshToken().String()))
	}
}
