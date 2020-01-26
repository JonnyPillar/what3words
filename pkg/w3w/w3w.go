package w3w

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jonnypillar/what3words/internal/api"
)

const (
	wordsDelimiter            = "."
	convertToWordsRoute       = "convert-to-3wa"
	convertToCoordinatesRoute = "convert-to-coordinates"

	formatJSON    = "json"
	formatGeoJSON = "geojson"
)

// Client ...
type Client struct {
	key string
}

// New ...
func New(apiKey string) *Client {
	return &Client{
		key: apiKey,
	}
}

// GetCoordinates ...
func (c Client) GetCoordinates(req *Words, options *Options) (*Result, error) {
	url, err := c.coordinatesURL(req, options)
	if err != nil {
		return nil, err
	}

	resp, err := api.Get(url)
	if err != nil {
		var apiErr api.ErrorResponse

		if errors.As(err, &apiErr) {
			return nil, Error{
				Code:    apiErr.Err.Code,
				Message: apiErr.Err.Message,
			}
		}

		return nil, err
	}

	return newResponse(resp), nil
}

// GetWords ...
func (c Client) GetWords(req *Coordinates, options *Options) (*Result, error) {
	url, err := c.wordsURL(req, options)
	if err != nil {
		return nil, err
	}

	resp, err := api.Get(url)
	if err != nil {
		var apiErr api.ErrorResponse

		if errors.As(err, &apiErr) {
			return nil, Error{
				Code:    apiErr.Err.Code,
				Message: apiErr.Err.Message,
			}
		}

		return nil, err
	}

	return newResponse(resp), nil
}

func (c Client) coordinatesURL(req *Words, options *Options) (string, error) {
	url, err := api.NewURL(c.key, options.APIURL, convertToCoordinatesRoute)
	if err != nil {
		return "", err
	}

	url.AddParam("words", strings.Join(req[:], wordsDelimiter))

	switch options.Format {
	case formatGeoJSON:
		url.AddParam("format", formatGeoJSON)
	default:
		url.AddParam("format", formatJSON)
	}

	return url.URL(), nil
}

func (c Client) wordsURL(req *Coordinates, options *Options) (string, error) {
	url, err := api.NewURL(c.key, options.APIURL, convertToWordsRoute)
	if err != nil {
		return "", err
	}

	url.AddParam("coordinates", fmt.Sprintf("%f,%f", req.Lat, req.Lng))

	// switch options.Format {
	// case formatGeoJSON:
	// 	url.AddParam("format", formatGeoJSON)
	// default:
	// 	url.AddParam("format", formatJSON)
	// }

	return url.URL(), nil
}
