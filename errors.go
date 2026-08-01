package vrage

import (
	"errors"
)

// Predefined errors for common failure scenarios in API client.
var (
	ErrConnectionFailed   = errors.New("failed to connect to the server: check if the server is running and reachable")
	ErrInvalidSecurityKey = errors.New("server returned StatusForbidden: security key is invalid or missing")
	ErrRequestTimeout     = errors.New("request timed out: the server did not respond in time")
)
