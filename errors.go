package w3w

import "fmt"

// Error ...
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error ...
func (w Error) Error() string {
	return fmt.Sprintf("%s: %s", w.Code, w.Message)
}
