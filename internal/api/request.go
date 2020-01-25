package api

import (
	"encoding/json"
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
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse

		err = json.Unmarshal(body, &errResp)
		if err != nil {
			return nil, err
		}

		return nil, errResp
	}

	var wResp Response
	err = json.Unmarshal(body, &wResp)
	if err != nil {
		return nil, err
	}

	return &wResp, nil
}
