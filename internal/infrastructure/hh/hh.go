package hh

const (
	grantType    = "grant_type"
	refreshToken = "refresh_token"
	pathToken    = "/token"

	headerAuthorization        = "Authorization"
	headerContentType          = "Content-Type"
	applicationXFormUrlEncoded = "application/x-www-form-urlencoded"
	bearer                     = "Bearer "
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
