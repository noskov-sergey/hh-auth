package pgsql

import (
	"time"

	"github.com/noskov-sergey/hh-auth/internal/domain"
	"github.com/noskov-sergey/hh-auth/internal/domain/agregates"
)

type entityAuthorization struct {
	ID           int32     `json:"id"`
	AccessToken  string    `json:"access_token"`
	ExpiredAt    time.Time `json:"expired_at"`
	RefreshToken string    `json:"refresh_token"`
	CreateAt     time.Time `json:"create_at"`
	Status       string    `json:"status"`
}

func fromEntityAuthorization(e entityAuthorization) (*agregates.Authorization, error) {
	params := agregates.AuthorizationParams{
		ID:           agregates.ID(e.ID),
		AccessToken:  domain.AccessToken(e.AccessToken),
		Expired:      e.ExpiredAt,
		RefreshToken: domain.RefreshToken(e.RefreshToken),
		Created:      e.CreateAt,
		Status:       domain.Status(e.Status),
	}

	return agregates.NewAuthorization(params)
}
