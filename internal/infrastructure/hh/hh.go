package hh

import "errors"

const (
	grantType    = "grant_type"
	refreshToken = "refresh_token"
	pathToken    = "/token"
	schema       = "https"

	headerAuthorization        = "Authorization"
	headerContentType          = "Content-Type"
	applicationXFormUrlEncoded = "application/x-www-form-urlencoded"
	bearer                     = "Bearer "
)

var (
	ErrTokenNotExpired = errors.New("token not expired")
)

type Client struct {
	addr      string
	grantType string
}

func NewClient(addr, grantType string) *Client {
	return &Client{
		addr:      addr,
		grantType: grantType,
	}
}
