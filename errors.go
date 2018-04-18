package gopeana

import (
	"fmt"
	"net/http"
)

// StatusError is used for error responses from the Europeana API endpoint.
type StatusError struct {
	StatusCode int
	StatusText string
}

func newStatusError(code int) StatusError {
	return StatusError{
		StatusCode: code,
		StatusText: http.StatusText(code),
	}
}

// Error returns a string representation of the StatusError.
func (e StatusError) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode, e.StatusText)
}
