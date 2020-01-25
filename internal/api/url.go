package api

import (
	"fmt"
	"net/url"
)

const (
	w3wAPIURL = "https://api.what3words.com/v3"
)

// URL ...
type URL struct {
	url    *url.URL
	params url.Values
}

// NewURL ...
func NewURL(apiKey, baseURL, route string) (*URL, error) {
	var apiURL string
	if baseURL == "" {
		apiURL = w3wAPIURL
	} else {
		apiURL = baseURL
	}

	if route == "" {
		return nil, fmt.Errorf("invalid w3w route")
	}

	u, err := url.Parse(fmt.Sprintf("%s/%s", apiURL, route))
	if err != nil {
		return nil, fmt.Errorf("invalid w3w API URL: %w", err)
	}

	params := url.Values{}

	if apiKey == "" {
		return nil, fmt.Errorf("invalid api key")
	}

	params.Add("key", apiKey)

	return &URL{
		url:    u,
		params: params,
	}, nil
}

// AddParam ...
func (u URL) AddParam(key, value string) {
	u.params.Add(key, value)
}

// URL ...
func (u URL) URL() string {
	u.url.RawQuery = u.params.Encode()

	return u.url.String()
}
