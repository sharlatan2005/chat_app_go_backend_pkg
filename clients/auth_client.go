package clients

import (
	"net/http"
	"time"
)

type AuthClient struct {
	baseURL string
	http    *http.Client
}

func NewAuthClient(baseURL string) *AuthClient {
	return &AuthClient{
		baseURL: baseURL,
		http:    &http.Client{Timeout: 5 * time.Second},
	}
}

func (c *AuthClient) ValidateToken(token string) (int, error) {
	return 0, nil
}
