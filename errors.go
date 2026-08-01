package vrage

import (
	"errors"
)

var (
	// remote errors (invoked by the server)

	ErrConnectionFailed   = errors.New("failed to connect to the server: check if the server is running and reachable")
	ErrInvalidSecurityKey = errors.New("server returned StatusForbidden: security key is invalid or missing")
	ErrRequestTimeout     = errors.New("request timed out: the server did not respond in time")
)
