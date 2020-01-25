package api

import "fmt"

// Response defines the response body for Coordinates request
type Response struct {
	Country string `json:"country"`
	Square  struct {
		Southwest struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"southwest"`
		Northeast struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"northeast"`
	} `json:"square"`
	NearestPlace string `json:"nearestPlace"`
	Coordinates  struct {
		Lng float64 `json:"lng"`
		Lat float64 `json:"lat"`
	} `json:"coordinates"`
	Words    string `json:"words"`
	Language string `json:"language"`
	Map      string `json:"map"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Err struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// Error ...
func (w ErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", w.Err.Code, w.Err.Message)
}
