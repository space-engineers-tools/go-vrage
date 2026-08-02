package vrage

import (
	"errors"
	"fmt"
	"strings"
)

// ErrConfig... are errors for configuration validation
var (
	ErrConfigIncomplete   = errors.New("invalid config: missing required fields")
	ErrConfigInvalid      = errors.New("invalid config: field validation failed")
	ErrConfigIncompatible = errors.New("invalid config: conflicting settings")
)

// newErrConfigIncomplete creates a new ErrConfigIncomplete.
func newErrConfigIncomplete(fieldNames ...string) error {
	joinedFields := strings.Join(fieldNames, ", ")
	return fmt.Errorf("%w: [%s]", ErrConfigIncomplete, joinedFields)
}

// newErrConfigInvalid creates a new ErrConfigInvalid.
func newErrConfigInvalid(fieldName string, reason string) error {
	return fmt.Errorf("%w: field '%s' is invalid: %s", ErrConfigInvalid, fieldName, reason)
}

// newErrConfigIncompatible creates a new ErrConfigIncompatible.
func newErrConfigIncompatible(field1, field2 string, reason string) error {
	return fmt.Errorf("%w: fields '%s' and '%s' conflict: %s", ErrConfigIncompatible, field1, field2, reason)
}

// todo: errors below are not final. they are just ideas.

// ErrAPI... are errors for API request failures
var (
	ErrAPIConnectionFailed   = errors.New("failed to connect to the server: check if the server is running and reachable")
	ErrAPIInvalidSecurityKey = errors.New("server returned StatusForbidden: security key is invalid or missing")
	ErrAPIRequestTimeout     = errors.New("request timed out: the server did not respond in time")
)
