package w3w

import (
	"fmt"

	"github.com/jonnypillar/what3words/internal/api"
)

var (
	// ErrNoAPIKey ...
	ErrNoAPIKey = fmt.Errorf("invalid or empty API Key provided")
	// ErrEmptyWord ...
	ErrEmptyWord = fmt.Errorf("an empty words was provided")
	// ErrInvalidNumberOfWords ...
	ErrInvalidNumberOfWords = fmt.Errorf("invalid number of words provided")
)

// Error ...
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error ...
func (w Error) Error() string {
	return fmt.Sprintf("%s: %s", w.Code, w.Message)
}

func newResponseError(err api.ErrorResponse) Error {
	return Error{
		Code:    err.Err.Code,
		Message: err.Err.Message,
	}
}
