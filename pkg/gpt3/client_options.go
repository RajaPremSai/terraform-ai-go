package gpt3

import (
	"net/http"
	"time"
)

// ClientOption are options that can be passed when creating a new gpt client.
type ClientOption func(*client) error

func WithAPIVersion(apiVersion string) ClientOption {
	return func(c *client) error {
		c.apiVersion = apiVersion
		return nil
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(c *client) error {
		c.userAgent = userAgent
		return nil
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *client) error {
		c.httpClient = httpClient
		return nil
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *client) error {
		c.httpClient.Timeout = timeout
		return nil
	}
}
