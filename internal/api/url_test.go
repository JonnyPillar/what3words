package api_test

import (
	"fmt"
	"testing"

	"github.com/jonnypillar/what3words/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestURLURL(t *testing.T) {
	testCases := []struct {
		desc    string
		apiKey  string
		baseURL string
		route   string

		expectedURL string
		expectedErr error
	}{
		{
			desc:    "given an valid set of keys, url and routes, a valid URL is returned",
			apiKey:  "FooBar",
			baseURL: "https://example.com",
			route:   "convert-to-3wa",

			expectedURL: "https://example.com/convert-to-3wa?key=FooBar",
		},
		{
			desc:    "given that the API key is not provided, an error is returned",
			baseURL: "https://api.what3words.com/v3",
			route:   "convert-to-3wa",

			expectedErr: fmt.Errorf("invalid api key"),
		},
		{
			desc:    "given that the route is not provided, an error is returned",
			apiKey:  "FooBar",
			baseURL: "https://api.what3words.com/v3",

			expectedErr: fmt.Errorf("invalid w3w route"),
		},
		// {
		// 	desc:    "given an invalid API URL is provided, an error is returned",
		// 	apiKey:  "FooBar",
		// 	baseURL: "foo.html",
		// 	route:   "convert-to-3wa",

		// 	expectedErr: fmt.Errorf("invalid w3w route"),
		// },
		{
			desc:   "given that the API URL is not provided, the default W3W URL is set in the URL",
			apiKey: "FooBar",
			route:  "convert-to-3wa",

			expectedURL: "https://api.what3words.com/v3/convert-to-3wa?key=FooBar",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			url, err := api.NewURL(tt.apiKey, tt.baseURL, tt.route)

			if tt.expectedErr != nil {
				assert.NotNil(t, err)
				assert.EqualError(t, tt.expectedErr, err.Error())
				assert.Nil(t, url)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedURL, url.URL())
			}
		})
	}
}
