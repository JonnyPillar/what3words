package w3w_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jonnypillar/what3words/internal/api"
	"github.com/jonnypillar/what3words/pkg/w3w"
	"github.com/stretchr/testify/assert"
)

const (
	apiKey = "foobar"
)

func TestGetCoordinates(t *testing.T) {
	tests := []struct {
		desc   string
		apiKey string
		words  w3w.Words

		apiResponse   interface{}
		apiStatusCode int

		expectedAPIURL string
		expectedCoords *w3w.Result
		expectedErr    error
	}{
		{
			desc:   "given a slice of words, coordinates returned",
			apiKey: apiKey,
			words:  w3w.Words{"one", "two", "three"},

			apiResponse: api.Response{
				Coordinates: struct {
					Lng float64 "json:\"lng\""
					Lat float64 "json:\"lat\""
				}{
					Lat: 1,
					Lng: 2,
				},
			},
			apiStatusCode: http.StatusOK,

			expectedAPIURL: "/convert-to-coordinates?format=json&key=foobar&words=one.two.three",
			expectedCoords: &w3w.Result{
				Coordinates: w3w.Coords{
					Lat: 1,
					Lng: 2,
				},
			},
		},
		{
			desc: "given an invalid API Key is set, error returned",

			expectedErr: fmt.Errorf("invalid api key"),
		},
		{
			desc:   "given the W3W API returns an error, error returned",
			apiKey: apiKey,
			words:  w3w.Words{"one", "two", "three"},

			apiResponse: api.ErrorResponse{
				Err: struct {
					Code    string `json:"code"`
					Message string `json:"message"`
				}{
					Code:    "BadWords",
					Message: "Invalid or non-existent 3 word address",
				},
			},
			apiStatusCode: http.StatusBadRequest,

			expectedAPIURL: "/convert-to-coordinates?format=json&key=foobar&words=one.two.three",
			expectedErr:    fmt.Errorf("BadWords: Invalid or non-existent 3 word address"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			s := testServer(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tt.expectedAPIURL, r.URL.String())

				b, _ := json.Marshal(tt.apiResponse)

				w.WriteHeader(tt.apiStatusCode)
				w.Write(b)
			})

			defer s.Close()

			c := w3w.New(tt.apiKey)
			coords, err := c.GetCoordinates(
				&tt.words,
				&w3w.Options{
					APIURL: s.URL,
				},
			)

			if tt.expectedErr != nil {
				assert.NotNil(t, err)
				assert.EqualError(t, tt.expectedErr, err.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedCoords, coords)
			}
		})
	}
}

func TestGetWords(t *testing.T) {
	tests := []struct {
		desc   string
		apiKey string
		coords w3w.Coordinates

		apiResponse   interface{}
		apiStatusCode int

		expectedAPIURL string
		expectedWords  *w3w.Result
		expectedErr    error
	}{
		{
			desc:   "given a coordinates, words returned",
			apiKey: apiKey,
			coords: w3w.Coordinates{
				Lat: 51.432393,
				Lng: -0.348023,
			},

			apiResponse: api.Response{
				Words: "one.two.three",
			},
			apiStatusCode: http.StatusOK,

			expectedAPIURL: "/convert-to-3wa?coordinates=51.432393%2C-0.348023&key=foobar",
			expectedWords: &w3w.Result{
				// "one", "two", "three"
				Words: "one.two.three",
			},
		},
		{
			desc: "given an invalid API Key is set, error returned",

			expectedErr: fmt.Errorf("invalid api key"),
		},
		{
			desc:   "given the W3W API returns an error, error returned",
			apiKey: apiKey,
			coords: w3w.Coordinates{
				Lat: 51.432393,
				Lng: -0.348023,
			},

			apiResponse: api.ErrorResponse{
				Err: struct {
					Code    string `json:"code"`
					Message string `json:"message"`
				}{
					Code:    "BadCoordinates",
					Message: "coordinates must be two comma separated lat,lng coordinates",
				},
			},
			apiStatusCode: http.StatusBadRequest,

			expectedAPIURL: "/convert-to-3wa?coordinates=51.432393%2C-0.348023&key=foobar",
			expectedErr:    fmt.Errorf("BadCoordinates: coordinates must be two comma separated lat,lng coordinates"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			s := testServer(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tt.expectedAPIURL, r.URL.String())

				b, _ := json.Marshal(tt.apiResponse)

				w.WriteHeader(tt.apiStatusCode)
				w.Write(b)
			})
			defer s.Close()

			c := w3w.New(tt.apiKey)

			words, err := c.GetWords(
				&tt.coords,
				&w3w.Options{
					APIURL: s.URL,
				},
			)

			if tt.expectedErr != nil {
				assert.NotNil(t, err)
				assert.EqualError(t, tt.expectedErr, err.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedWords, words)
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
