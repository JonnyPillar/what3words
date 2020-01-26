package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	requestTimeout = 30 * time.Second
)

// Get ...
func Get(url string) (*Response, error) {
	http.DefaultClient.Timeout = requestTimeout

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing get request %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error occurred reading response body %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse

		err = json.Unmarshal(body, &errResp)
		if err != nil {
			return nil, fmt.Errorf("invalid error JSON returned from API %w", err)
		}

		return nil, errResp
	}

	var wResp Response
	err = json.Unmarshal(body, &wResp)
	if err != nil {
		return nil, fmt.Errorf("invalid JSON returned from API %w", err)
	}

	return &wResp, nil
}
