package w3w

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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
func (c Client) GetCoordinates(req *Words, options *Options) (*Coordinates, error) {
	url, err := c.coordinatesURL(req, options)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var cResp ErrorResponse
		json.Unmarshal(body, &cResp)

		return nil, cResp
	}

	var cResp Response
	json.Unmarshal(body, &cResp)

	return &Coordinates{
		Lat: cResp.Coordinates.Lat,
		Lng: cResp.Coordinates.Lng,
	}, nil
}

// GetWords ...
func (c Client) GetWords(req *Coordinates, options *Options) (*Words, error) {
	url, err := c.wordsURL(req, options)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var cResp ErrorResponse
		json.Unmarshal(body, &cResp)

		return nil, cResp
	}

	var wResp Response
	json.Unmarshal(body, &wResp)

	words := strings.Split(wResp.Words, wordsDelimiter)

	w := Words{}
	copy(w[:], words[:3])

	return &w, nil
}

func (c Client) coordinatesURL(req *Words, options *Options) (string, error) {
	baseURL, err := url.Parse(
		fmt.Sprintf("%s/%s", options.APIURL, convertToCoordinatesRoute),
	)
	if err != nil {
		return "", fmt.Errorf("invalid w3w API URL")
	}

	params := url.Values{}

	if c.key == "" {
		return "", fmt.Errorf("invalid api key")
	}

	params.Add("key", c.key)
	params.Add("words", strings.Join(req[:], wordsDelimiter))

	switch options.Format {
	case formatGeoJSON:
		params.Add("format", formatGeoJSON)
	default:
		params.Add("format", formatJSON)
	}

	baseURL.RawQuery = params.Encode()

	return baseURL.String(), nil
}

func (c Client) wordsURL(req *Coordinates, options *Options) (string, error) {
	baseURL, err := url.Parse(
		fmt.Sprintf("%s/%s", options.APIURL, convertToWordsRoute),
	)
	if err != nil {
		return "", fmt.Errorf("invalid w3w API URL")
	}

	params := url.Values{}

	if c.key == "" {
		return "", fmt.Errorf("invalid api key")
	}

	params.Add("key", c.key)

	params.Add("coordinates", fmt.Sprintf("%f,%f", req.Lat, req.Lng))

	// switch options.Format {
	// case formatGeoJSON:
	// 	params.Add("format", formatGeoJSON)
	// default:
	// 	params.Add("format", formatJSON)
	// }

	baseURL.RawQuery = params.Encode()

	return baseURL.String(), nil
}
