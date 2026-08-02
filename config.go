package vrage

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

// Default configuration values for the VRage Remote API client.
const (
	DefaultTimeout     time.Duration = 10 * time.Second
	DefaultBasePath    string        = "/vrageremote"
	DefaultContentType string        = "application/json"
	DefaultPort        uint32        = 8080
)

var validate = validator.New()

// ToPtr returns a pointer to the given value.
func ToPtr[T any](value T) *T {
	return &value
}

// ClientConfig holds the configuration settings for the VRage Remote API client.
type ClientConfig struct {
	// RemoteApiIP is the IP address or DNS name of the Space Engineers server.
	//
	// Corresponding Setting in SpaceEngineers-Dedicated.cfg:
	//  <RemoteApiIP>
	//
	// Examples:
	//  "127.0.0.1"
	//  "example.com"
	//  "play.cool-server.com"
	RemoteApiIP string `validate:"required,ip|fqdn"` //nolint:revive // so the name is closer to the .cfg file

	// RemoteSecurityKey is the security key used for authenticating API requests.
	//
	// Corresponding Setting in SpaceEngineers-Dedicated.cfg:
	// 	<RemoteSecurityKey>
	RemoteSecurityKey string `validate:"required"`

	// RemoteApiPort is the port of the Remote API on the Space Engineers server.
	//
	// Corresponding Setting in SpaceEngineers-Dedicated.cfg:
	//  <RemoteApiPort>
	//
	// Default:
	//  8080
	RemoteApiPort uint32 `validate:"port"` //nolint:revive // so the name is closer to the .cfg file

	// Timeout specifies the maximum duration for an API request before an vrage.ErrRequestTimeout error is returned.
	//
	// Default:
	//  vrage.DefaultTimeout
	Timeout time.Duration `validate:"gte=0"`

	// UseHTTPS indicates whether to use HTTPS for API requests.
	//
	// Note: While the Space Engineers server does not natively support HTTPS, this option
	// can be used when routing through a reverse proxy.
	//
	// Default:
	//  false
	UseHTTPS bool `validate:"-"`

	// BasePath is the base route path for API requests.
	//
	// Note: While this path is fixed by the Space Engineers server, this option
	// allows customization when routing through a reverse proxy.
	//
	// Example: "/custompath", "/", or ""
	//
	// Default:
	//  ToPtr(DefaultAPIEndpoint)
	BasePath *string `validate:"-"`

	// HTTPClient allows the use of a custom HTTP client for making requests.
	//
	// If not provided, a default client with the specified Timeout will be used.
	//
	// Warning: when using a custom HTTPClient the Timeout field in ClientConfig will be ignored.
	// In this case, ensure the custom HTTPClient has an appropriate timeout set to avoid hanging requests.
	HTTPClient *http.Client `validate:"-"`
}

// Validate checks the ClientConfig for required fields and valid values.
func (c *ClientConfig) Validate() error {
	return validate.Struct(c)
}

// SetDefaults initializes default values for ClientConfig fields that are not explicitly set.
func (c *ClientConfig) SetDefaults() {
	if c.RemoteApiPort == 0 {
		c.RemoteApiPort = DefaultPort
	}

	if c.Timeout == 0 {
		c.Timeout = DefaultTimeout
	}
	if c.BasePath == nil {
		c.BasePath = ToPtr(DefaultBasePath)
	}
	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{
			Timeout: c.Timeout,
		}
	}
}
