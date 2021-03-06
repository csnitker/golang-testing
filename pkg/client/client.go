package client

import (
	"context"
	"net/http"
)

func StrToPtr(v string) *string {
	return &v
}

type Getter interface {
	Get(url string) (*http.Response, error)
}

type Client struct {
	HTTP Getter
}

func (c *Client) Do(ctx context.Context) (*http.Response, error) {
	return c.HTTP.Get("https://google.com")
}
