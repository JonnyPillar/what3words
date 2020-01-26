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

	paramWords       = "words"
	paramCoordinates = "coordinates"
	paramFormat      = "format"
	paramLanguage    = "language"
)

// Client defines the W3W Client
type Client struct {
	key string
}

// New initalises a new Client instance
// The `key` parameter sets the API Key used to authenticate against W3W APIs.
// For information on how to get this value see https://accounts.what3words.com/en/account/developer
//
// If no API Key is provided, ErrNoAPIKey is returned
func New(key string) (*Client, error) {
	if key == "" || strings.TrimSpace(key) == "" {
		return nil, ErrNoAPIKey
	}

	return &Client{
		key: key,
	}, nil
}

// GetCoordinates converts a 3 word address into a Longitude and Latitude along with the country,
// the bounds of the grid square, a nearby place and a link to the W3W site
func (c Client) GetCoordinates(req Words, options *CoordinateOptions) (*Result, error) {
	url, err := c.coordinatesURL(req, options)
	if err != nil {
		return nil, err
	}

	resp, err := api.Get(url)
	if err != nil {
		var apiErr api.ErrorResponse

		if errors.As(err, &apiErr) {
			return nil, newResponseError(apiErr)
		}

		return nil, err
	}

	return newResponse(resp), nil
}

// GetWords converts a Longitude and Latitude into a 3 word address along with the country,
// the bounds of the grid square, a nearby place and a link to the W3W site
func (c Client) GetWords(req Coordinates, opts *WordOptions) (*Result, error) {
	url, err := c.wordsURL(req, opts)
	if err != nil {
		return nil, err
	}

	resp, err := api.Get(url)
	if err != nil {
		var apiErr api.ErrorResponse

		if errors.As(err, &apiErr) {
			return nil, newResponseError(apiErr)
		}

		return nil, err
	}

	return newResponse(resp), nil
}

func (c Client) coordinatesURL(req Words, opts *CoordinateOptions) (string, error) {
	url, err := api.NewURL(c.key, opts.APIURL, convertToCoordinatesRoute)
	if err != nil {
		return "", err
	}

	err = validateWordParam(req)
	if err != nil {
		return "", err
	}

	url.AddParam(paramWords, strings.Join(req[:], wordsDelimiter))

	switch opts.Format {
	case formatGeoJSON:
		url.AddParam(paramFormat, formatGeoJSON)
	default:
		url.AddParam(paramFormat, formatJSON)
	}

	return url.URL(), nil
}

func (c Client) wordsURL(req Coordinates, opts *WordOptions) (string, error) {
	url, err := api.NewURL(c.key, opts.APIURL, convertToWordsRoute)
	if err != nil {
		return "", err
	}

	url.AddParam(paramCoordinates, fmt.Sprintf("%f,%f", req.Lat, req.Lng))

	if opts.Language != "" {
		url.AddParam(paramLanguage, opts.Language)
	}

	switch opts.Format {
	case formatGeoJSON:
		url.AddParam(paramFormat, formatGeoJSON)
	default:
		url.AddParam(paramFormat, formatJSON)
	}

	return url.URL(), nil
}

func validateWordParam(req Words) error {
	var count int
	for _, word := range req {
		if word == "" || strings.TrimSpace(word) == "" {
			return ErrEmptyWord
		}
		count++
	}

	return nil
}
