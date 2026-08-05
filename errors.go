package vrage

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// ErrConfig... are errors for configuration validation.
var (
	ErrConfigIncomplete   = errors.New("invalid config: missing required fields")
	ErrConfigInvalid      = errors.New("invalid config: field validation failed")
	ErrConfigIncompatible = errors.New("invalid config: conflicting settings")
)

// newErrConfigIncomplete creates a new ErrConfigIncomplete.
func newErrConfigIncomplete(fieldNames ...string) error { //nolint:unused // todo
	joinedFields := strings.Join(fieldNames, ", ")
	return fmt.Errorf("%w: [%s]", ErrConfigIncomplete, joinedFields)
}

// newErrConfigInvalid creates a new ErrConfigInvalid.
func newErrConfigInvalid(fieldName string, reason string) error { //nolint:unused // todo
	return fmt.Errorf("%w: field '%s' is invalid: %s", ErrConfigInvalid, fieldName, reason)
}

// newErrConfigIncompatible creates a new ErrConfigIncompatible.
func newErrConfigIncompatible(field1, field2 string, reason string) error { //nolint:unused // todo
	return fmt.Errorf("%w: fields '%s' and '%s' conflict: %s", ErrConfigIncompatible, field1, field2, reason)
}

// ErrAPI... are errors for API request failures.
var (
	ErrAPIConnectionFailed   = errors.New("failed to connect to the server: connection refused or host not available")
	ErrAPIInvalidSecurityKey = errors.New("server returned StatusForbidden: security key is invalid or missing")
	ErrAPIRequestTimeout     = errors.New("request timed out: the server did not respond in time")
	ErrAPIUnexpectedCode     = errors.New("unexpected status code from the server")
	ErrAPIUnexpectedBody     = errors.New("unexpected response body from the server")
)

func newErrAPIConnectionFailed(err error) error {
	return fmt.Errorf("%w: %s", ErrAPIConnectionFailed, err.Error())
}

func newErrAPIInvalidSecurityKey() error {
	return ErrAPIInvalidSecurityKey
}

func newErrAPIRequestTimeout(timeout time.Duration) error {
	return fmt.Errorf("%w: exceeded %s", ErrAPIRequestTimeout, timeout)
}

func newErrAPIUnexpectedCode(statusCode int, body string) error {
	return fmt.Errorf("%w: status code %d, body: %s", ErrAPIUnexpectedCode, statusCode, body)
}

func newErrAPIUnexpectedBody(body string) error {
	return fmt.Errorf("%w: body: %s", ErrAPIUnexpectedBody, body)
}
