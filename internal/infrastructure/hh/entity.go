package hh

import (
	"github.com/noskov-sergey/hh-auth/internal/domain"
	"github.com/noskov-sergey/hh-auth/internal/domain/agregates"
)

type ErrorResponse struct {
	Error       string `json:"error"`
	Description string `json:"error_description"`
}

type EntityRefreshResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

func responseToRefresh(r EntityRefreshResponse) (*agregates.Refresh, error) {
	params := agregates.RefreshParams{
		AccessToken:  domain.AccessToken(r.AccessToken),
		ExpiresIn:    domain.ExpiresIn(r.ExpiresIn),
		RefreshToken: domain.RefreshToken(r.RefreshToken),
	}

	return agregates.NewRefreshResponse(params)
}
