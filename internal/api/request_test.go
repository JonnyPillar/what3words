package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jonnypillar/what3words/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	testCases := []struct {
		desc    string
		url     string
		handler func(w http.ResponseWriter, r *http.Request)

		expectedResp *api.Response
		expectedErr  error
	}{
		{
			desc: "given API returns a bad request, error response returned",
			handler: func(w http.ResponseWriter, r *http.Request) {
				resp := api.Response{
					Coordinates: struct {
						Lng float64 "json:\"lng\""
						Lat float64 "json:\"lat\""
					}{
						Lat: 1,
						Lng: 2,
					},
				}

				b, _ := json.Marshal(resp)
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			},

			expectedResp: &api.Response{
				Coordinates: struct {
					Lng float64 "json:\"lng\""
					Lat float64 "json:\"lat\""
				}{
					Lat: 1,
					Lng: 2,
				},
			},
		},
		{
			desc: "given API returns a bad request, error response returned",
			handler: func(w http.ResponseWriter, r *http.Request) {
				err := api.ErrorResponse{
					Err: struct {
						Code    string `json:"code"`
						Message string `json:"message"`
					}{
						Code:    "BadWords",
						Message: "Invalid or non-existent 3 word address",
					},
				}

				b, _ := json.Marshal(err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(b)
			},

			expectedErr: api.ErrorResponse{
				Err: struct {
					Code    string `json:"code"`
					Message string `json:"message"`
				}{
					Code:    "BadWords",
					Message: "Invalid or non-existent 3 word address",
				},
			},
		},
		{
			desc: "given the http client returns an error, error returned",
			url:  "invalidurl",

			expectedErr: fmt.Errorf(`error occurred performing get request Get invalidurl: unsupported protocol scheme ""`),
		},
		{
			desc: "given the the API returns invalid JSON, error returned",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.Write(nil)
			},
			expectedErr: fmt.Errorf("invalid JSON returned from API unexpected end of JSON input"),
		},
		{
			desc: "given the the API returns invalid error JSON, error returned",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(nil)
			},
			expectedErr: fmt.Errorf("invalid error JSON returned from API unexpected end of JSON input"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			s := testServer(tt.handler)
			defer s.Close()

			var url string
			if tt.url != "" {
				url = tt.url
			} else {
				url = s.URL
			}
			resp, err := api.Get(url)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}

func testServer(h func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	server := httptest.NewServer(
		http.HandlerFunc(h),
	)
	return server
}
