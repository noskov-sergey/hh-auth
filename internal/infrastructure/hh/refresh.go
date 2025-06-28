package hh

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/noskov-sergey/hh-auth/internal/domain"
	"github.com/noskov-sergey/hh-auth/internal/domain/agregates"
)

func (c *Client) RefreshAccessToken(ctx context.Context, rToken domain.RefreshToken, aToken domain.AccessToken) (*agregates.Refresh, error) {
	uri, err := url.Parse(c.addr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}
	uri.Path = pathToken

	data := url.Values{}
	data.Set(grantType, c.grantType)
	data.Add(refreshToken, rToken.String())

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), strings.NewReader(data.Encode()))
	req.Header.Add(headerAuthorization, bearer+aToken.String())
	req.Header.Add(headerContentType, applicationXFormUrlEncoded)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusBadRequest:
		return nil, fmt.Errorf("bad request: %s", resp.Status)
	case http.StatusForbidden:
		return nil, fmt.Errorf("forbidden: %s", resp.Status)
	}

	r, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	var got EntityRefreshResponse
	err = json.Unmarshal(r, &got)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return responseToRefresh(got)
}
