package agregates

import (
	"errors"
	"fmt"
	"time"

	"github.com/noskov-sergey/hh-auth/internal/domain"
)

var (
	ErrAuthorizationValidate = errors.New("refresh response validation error")
)

type (
	ID int32
)

type Authorization struct {
	id           ID
	accessToken  domain.AccessToken
	expired      time.Time
	refreshToken domain.RefreshToken
	status       domain.Status
	createdAt    time.Time
}

type AuthorizationParams struct {
	ID           ID
	AccessToken  domain.AccessToken
	RefreshToken domain.RefreshToken
	Status       domain.Status
	Expired      time.Time
	Created      time.Time
}

func NewAuthorization(p AuthorizationParams) (*Authorization, error) {
	err := errors.Join(validateAuthID(p.ID), validateAuthAccessToken(p.AccessToken), validateAuthRefreshToken(p.RefreshToken),
		validateAuthExpired(p.Expired), validateAuthStatus(p.Status), validateAuthCreated(p.Created))
	if err != nil {
		return nil, errors.Join(ErrAuthorizationValidate, err)
	}

	return &Authorization{
		accessToken:  p.AccessToken,
		expired:      p.Expired,
		refreshToken: p.RefreshToken,
		status:       p.Status,
	}, nil
}

func (a *Authorization) AccessToken() domain.AccessToken {
	return a.accessToken
}

func (a *Authorization) ExpiresIn() time.Time {
	return a.expired
}

func (a *Authorization) RefreshToken() domain.RefreshToken {
	return a.refreshToken
}

func validateAuthID(v ID) error {
	if v <= 0 {
		return fmt.Errorf("invalid id: %d", v)
	}
	return nil
}

func validateAuthAccessToken(v domain.AccessToken) error {
	if !v.IsValid() {
		return fmt.Errorf("invalid access token: %s", v)
	}
	return nil
}

func validateAuthRefreshToken(v domain.RefreshToken) error {
	if !v.IsValid() {
		return fmt.Errorf("invalid refresh token: %s", v)
	}
	return nil
}

func validateAuthExpired(v time.Time) error {
	if v.IsZero() {
		return fmt.Errorf("invalid expired: %s", v.String())
	}
	return nil
}

func validateAuthCreated(v time.Time) error {
	if v.IsZero() {
		return fmt.Errorf("invalid created: %s", v.String())
	}
	return nil
}

func validateAuthStatus(v domain.Status) error {
	if !v.IsValid() {
		return fmt.Errorf("invalid status: %s", v)
	}
	return nil
}
