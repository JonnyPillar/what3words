// +build e2e

package e2e_test

import (
	"os"
	"testing"

	"github.com/jonnypillar/what3words/pkg/w3w"
	"github.com/stretchr/testify/assert"
)

func TestE2EWords(t *testing.T) {
	testCases := []struct {
		desc  string
		words w3w.Coordinates
		opts  w3w.WordOptions

		expectedErr    error
		expectedResult w3w.Result
	}{
		{
			desc: "given Lat & Lon coordinates, results returned",
			words: w3w.Coordinates{
				Lng: -0.195521,
				Lat: 51.520847,
			},
			opts: w3w.WordOptions{},

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
		{
			desc: "given Lat & Lon coordinates & FR language set, french results returned",
			words: w3w.Coordinates{
				Lng: -0.195521,
				Lat: 51.520847,
			},
			opts: w3w.WordOptions{
				Language: "fr",
			},

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
				NearestPlace: "Bayswater, Grand Londres",
				Coordinates: w3w.Coords{
					Lng: -0.195521,
					Lat: 51.520847,
				},
				Words:    "conduite.richissime.empâter",
				Language: "fr",
				Map:      "https://w3w.co/conduite.richissime.emp%C3%A2ter",
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			apiKey := os.Getenv("W3W_INTEGRATION_API_KEY")
			c, _ := w3w.New(apiKey)

			res, err := c.GetWords(tt.words, tt.opts)

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
