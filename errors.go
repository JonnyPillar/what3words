package w3w

import "fmt"

// ErrorResponse ...
type ErrorResponse struct {
	Err struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (w ErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", w.Err.Code, w.Err.Message)
}
