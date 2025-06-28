package agregates

import (
	"errors"
	"fmt"

	"github.com/noskov-sergey/hh-auth/internal/domain"
)

var (
	ErrRefreshValidate = errors.New("refresh response validation error")
)

type Refresh struct {
	accessToken  domain.AccessToken
	expiresIn    domain.ExpiresIn
	refreshToken domain.RefreshToken
}

type RefreshParams struct {
	AccessToken  domain.AccessToken
	ExpiresIn    domain.ExpiresIn
	RefreshToken domain.RefreshToken
}

func NewRefreshResponse(p RefreshParams) (*Refresh, error) {
	err := errors.Join(validateAccessToken(p.AccessToken), validateRefreshToken(p.RefreshToken),
		validateExpired(p.ExpiresIn))
	if err != nil {
		return nil, errors.Join(ErrAuthorizationValidate, err)
	}

	return &Refresh{
		accessToken:  p.AccessToken,
		expiresIn:    p.ExpiresIn,
		refreshToken: p.RefreshToken,
	}, nil
}

func (a *Refresh) AccessToken() domain.AccessToken {
	return a.accessToken
}

func (a *Refresh) ExpiresIn() domain.ExpiresIn {
	return a.expiresIn
}

func (a *Refresh) RefreshToken() domain.RefreshToken {
	return a.refreshToken
}

func validateAccessToken(v domain.AccessToken) error {
	if !v.IsValid() {
		return fmt.Errorf("invalid access token: %s", v)
	}
	return nil
}

func validateRefreshToken(v domain.RefreshToken) error {
	if !v.IsValid() {
		return fmt.Errorf("invalid refresh token: %s", v)
	}
	return nil
}

func validateExpired(v domain.ExpiresIn) error {
	if !v.IsValid() {
		return fmt.Errorf("invalid expired: %d", v)
	}
	return nil
}
