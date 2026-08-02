package vrage_test

import (
	"errors"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	vrage "github.com/space-engineers-tools/go-vrage"
	"github.com/stretchr/testify/assert"
)

func strPtr(s string) *string {
	return &s
}

func validConfig() vrage.ClientConfig {
	return vrage.ClientConfig{
		RemoteApiIP:       "203.0.113.10",
		RemoteApiPort:     8080,
		RemoteSecurityKey: "top-secret",
		Timeout:           2 * time.Second,
		UseHTTPS:          false,
		BasePath:          strPtr("/vrageremote"),
	}
}

func TestClientConfigValidate_ValidIPv4(t *testing.T) {
	cfg := validConfig()

	err := cfg.Validate()

	assert.NoError(t, err)
}

func TestClientConfigValidate_ValidFQDN(t *testing.T) {
	cfg := validConfig()
	cfg.RemoteApiIP = "example.com"

	err := cfg.Validate()

	assert.NoError(t, err)
}

func TestClientConfigValidate_MissingRequiredFields(t *testing.T) {
	cfg := validConfig()
	cfg.RemoteApiIP = ""
	cfg.RemoteApiPort = 0
	cfg.RemoteSecurityKey = ""

	err := cfg.Validate()

	assert.Error(t, err)

	var verr validator.ValidationErrors
	assert.True(t, errors.As(err, &verr))
	assert.Len(t, verr, 3)

	fields := make(map[string]string, len(verr))
	for _, fe := range verr {
		fields[fe.Field()] = fe.Tag()
	}

	assert.Equal(t, "required", fields["RemoteApiIP"])
	assert.Equal(t, "port", fields["RemoteApiPort"])
	assert.Equal(t, "required", fields["RemoteSecurityKey"])
}

func TestClientConfigValidate_InvalidPort(t *testing.T) {
	cfg := validConfig()
	cfg.RemoteApiPort = 70000

	err := cfg.Validate()

	assert.Error(t, err)

	var verr validator.ValidationErrors
	assert.True(t, errors.As(err, &verr))
	assert.Equal(t, "RemoteApiPort", verr[0].Field())
	assert.Equal(t, "port", verr[0].Tag())
}

func TestClientConfigValidate_NegativeTimeout(t *testing.T) {
	cfg := validConfig()
	cfg.Timeout = -1 * time.Second

	err := cfg.Validate()

	assert.Error(t, err)

	var verr validator.ValidationErrors
	assert.True(t, errors.As(err, &verr))
	assert.Equal(t, "Timeout", verr[0].Field())
	assert.Equal(t, "gte", verr[0].Tag())
}

func TestClientConfigValidate_EndpointCanBeNilOrEmpty(t *testing.T) {
	cfg := validConfig()
	cfg.BasePath = nil

	err := cfg.Validate()
	assert.NoError(t, err)

	cfg.BasePath = strPtr("")
	err = cfg.Validate()
	assert.NoError(t, err)
}
