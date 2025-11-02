// Package gtvapi provides a client for the Gronkh.TV API.
//
// This package offers a comprehensive interface to interact with the Gronkh.TV API,
// including video information retrieval, comments, playlists, discovery features,
// and Twitch live status checking.
//
// Example usage:
//
//	client := gtvapi.NewClient(nil)
//	ctx := context.Background()
//	videoInfo, err := client.VideoInfo(ctx, 1000)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Video: %s\n", videoInfo.Title)
package gtvapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"resty.dev/v3"
)

const (
	// DefaultBaseURL is the default base URL for the Gronkh.TV API.
	DefaultBaseURL = "https://api.gronkh.tv/v1"

	// DefaultUserAgent is the default user agent string used for API requests.
	DefaultUserAgent = "gtvapi/2.0 (+https://github.com/wehmoen-dev/gtvapi)"

	// DefaultTimeout is the default timeout for API requests.
	DefaultTimeout = 30 * time.Second

	// DefaultMaxRetries is the default number of retry attempts for failed requests.
	DefaultMaxRetries = 3

	// DefaultRetryWaitMin is the minimum wait time between retries.
	DefaultRetryWaitMin = 1 * time.Second

	// DefaultRetryWaitMax is the maximum wait time between retries.
	DefaultRetryWaitMax = 30 * time.Second
)

// Client represents a Gronkh.TV API client.
type Client struct {
	httpClient *resty.Client
	baseURL    string
}

// ClientConfig holds configuration options for the Client.
type ClientConfig struct {
	// Debug enables debug logging for HTTP requests and responses.
	Debug bool

	// Timeout is the request timeout duration.
	// If not specified, DefaultTimeout is used.
	Timeout time.Duration

	// MaxRetries is the maximum number of retry attempts for failed requests.
	// If not specified, DefaultMaxRetries is used.
	MaxRetries int

	// RetryWaitMin is the minimum wait time between retries.
	// If not specified, DefaultRetryWaitMin is used.
	RetryWaitMin time.Duration

	// RetryWaitMax is the maximum wait time between retries.
	// If not specified, DefaultRetryWaitMax is used.
	RetryWaitMax time.Duration

	// BaseURL is the base URL for the API.
	// If not specified, DefaultBaseURL is used.
	BaseURL string

	// UserAgent is the user agent string for requests.
	// If not specified, DefaultUserAgent is used.
	UserAgent string

	// HTTPClient allows providing a custom HTTP client.
	// If nil, a default client will be created.
	HTTPClient *http.Client
}

// NewClient creates a new Gronkh.TV API client with the given configuration.
// If config is nil, default configuration values are used.
func NewClient(config *ClientConfig) *Client {
	if config == nil {
		config = &ClientConfig{}
	}

	// Set defaults
	if config.BaseURL == "" {
		config.BaseURL = DefaultBaseURL
	}
	if config.UserAgent == "" {
		config.UserAgent = DefaultUserAgent
	}
	if config.Timeout == 0 {
		config.Timeout = DefaultTimeout
	}
	if config.MaxRetries == 0 {
		config.MaxRetries = DefaultMaxRetries
	}
	if config.RetryWaitMin == 0 {
		config.RetryWaitMin = DefaultRetryWaitMin
	}
	if config.RetryWaitMax == 0 {
		config.RetryWaitMax = DefaultRetryWaitMax
	}

	// Create Resty client
	restyClient := resty.New().
		SetBaseURL(config.BaseURL).
		SetHeader("User-Agent", config.UserAgent).
		SetTimeout(config.Timeout).
		SetRetryCount(config.MaxRetries).
		SetRetryWaitTime(config.RetryWaitMin).
		SetRetryMaxWaitTime(config.RetryWaitMax).
		SetDebug(config.Debug)

	// Set custom HTTP client if provided
	if config.HTTPClient != nil {
		restyClient.SetTransport(config.HTTPClient.Transport)
	}

	return &Client{
		httpClient: restyClient,
		baseURL:    config.BaseURL,
	}
}

// get performs a GET request to the specified endpoint with optional query parameters.
func (c *Client) get(ctx context.Context, endpoint string, query map[string]string) ([]byte, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	requestURL, err := url.Parse(fmt.Sprintf("%s/%s", c.baseURL, endpoint))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	if query != nil {
		q := requestURL.Query()
		for key, value := range query {
			q.Set(key, value)
		}
		requestURL.RawQuery = q.Encode()
	}

	resp, err := c.httpClient.R().
		SetContext(ctx).
		Get(requestURL.String())

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), string(resp.Bytes()))
	}

	return resp.Bytes(), nil
}

// unmarshalResponse unmarshals the JSON response into the provided target.
func unmarshalResponse(data []byte, target interface{}) error {
	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return nil
}
