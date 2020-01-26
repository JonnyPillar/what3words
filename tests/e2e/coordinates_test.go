// +build e2e

package e2e_test

import (
	"os"
	"testing"

	"github.com/jonnypillar/what3words/pkg/w3w"
	"github.com/stretchr/testify/assert"
)

func TestE2ECoordinates(t *testing.T) {
	testCases := []struct {
		desc  string
		words w3w.Words
		opts  w3w.CoordinateOptions

		expectedErr    error
		expectedResult w3w.Result
	}{
		{
			desc: "given three words, results returned",
			words: w3w.Words{
				"filled",
				"count",
				"soap",
			},
			opts: w3w.CoordinateOptions{},

			expectedResult: w3w.Result{
				Country: "GB",
				Square: w3w.Square{
					Southwest: w3w.Southwest{
						Lng: -0.195543,
						Lat: 51.520833,
					},
					Northeast: w3w.Northeast{
						Lng: -0.195499,
						Lat: 51.52086,
					},
				},
				NearestPlace: "Bayswater, London",
				Coordinates: w3w.Coords{
					Lng: -0.195521,
					Lat: 51.520847,
				},
				Words:    "filled.count.soap",
				Language: "en",
				Map:      "https://w3w.co/filled.count.soap",
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			apiKey := os.Getenv("W3W_INTEGRATION_API_KEY")
			c, _ := w3w.New(apiKey)

			res, err := c.GetCoordinates(tt.words, tt.opts)

			if tt.expectedErr != nil {
				assert.NotNil(t, err)
				assert.EqualError(t, tt.expectedErr, err.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedResult, res)
			}
		})
	}
}
