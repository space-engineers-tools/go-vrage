package vrage

import (
	"errors"
)

var (
	// remote errors (invoked by the server)

	ErrConnectionFailed   error = errors.New("failed to connect to the server: check if the server is running and reachable")
	ErrInvalidSecurityKey error = errors.New("server returned StatusForbidden: security key is invalid or missing")
	ErrRequestTimeout     error = errors.New("request timed out: the server did not respond in time")
)
