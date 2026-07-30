package vrage

import (
	"errors"
	"net/http"
	"time"
)

const (
	DefaultTimeout     time.Duration = 10 * time.Second
	DefaultAPIEndpoint string        = "/vrageremote"
)

var (
	ErrInvalidConfig = errors.New("invalid client configuration: missing required fields")
)

// ClientConfig holds the configuration settings for the VRage Remote API client.
type ClientConfig struct {
	// RemoteApiIP is the IP address of the Space Engineers server for API requests.
	//
	// This corresponds to the <RemoteApiIP> tag in SpaceEngineers-Dedicated.cfg.
	RemoteApiIP string

	// RemoteApiPort is the port number of the Space Engineers server for API requests.
	//
	// This corresponds to the <RemoteApiPort> tag in SpaceEngineers-Dedicated.cfg.
	RemoteApiPort int

	// RemoteSecurityKey is the security key used for authenticating API requests.
	//
	// This corresponds to the <RemoteSecurityKey> tag in SpaceEngineers-Dedicated.cfg.
	RemoteSecurityKey string

	// Timeout specifies the maximum duration for an individual API request.
	// If a request exceeds this duration, context.DeadlineExceeded is returned.
	//
	// Default: vrage.DefaultTimeout
	Timeout time.Duration

	// APIEndpoint is the base route path for API requests.
	//
	// Note: While this path is fixed by the Space Engineers server, this option
	// allows customization when routing through a reverse proxy.
	//
	// Default: vrage.DefaultAPIEndpoint
	APIEndpoint string

	// HttpClient allows the use of a custom HTTP client for making requests.
	//
	// If not provided, a default client with the specified Timeout will be used.
	//
	// Warning: when using a custom HttpClient the Timeout field in ClientConfig will be ignored.
	// In this case, ensure the custom HttpClient has an appropriate timeout set to avoid hanging requests.
	HttpClient *http.Client
}

// setDefaults initializes default values for ClientConfig fields that are not explicitly set.
func (c *ClientConfig) setDefaults() {
	if c.Timeout == 0 {
		c.Timeout = DefaultTimeout
	}
	if c.APIEndpoint == "" {
		c.APIEndpoint = DefaultAPIEndpoint
	}
	if c.HttpClient == nil {
		c.HttpClient = &http.Client{
			Timeout: c.Timeout,
		}
	}
}

// Sender handles low-level HTTP request execution against the server.
//
// It can be invoked directly if raw communication or custom response handling is needed.
//
// Use at your own risk, as this bypasses the high-level abstractions provided by the Client.
type Sender struct {
	config     ClientConfig
	httpClient *http.Client
}

// Client is the high-level API client for the Space Engineers VRage Remote API.
//
// It provides typed methods for interacting with server endpoints.
// Before using this client, ensure <RemoteApiEnabled> is set to true in
// SpaceEngineers-Dedicated.cfg and the target server is running.
type Client struct {
	// Sender is the underlying request executor for the Client.
	//
	// While it can be used directly for raw requests, it's recommended to use the
	// high-level methods provided by Client for type safety and convenience.
	Sender *Sender
}

// GetConfig returns the ClientConfig of the current Client instance.
//
// Be careful not to leak sensitive information, such as the SecurityKey, when using this function.
func (c *Client) Config() ClientConfig {
	return c.Sender.config
}

// NewClient creates a new Client instance with the provided configuration.
func NewClient(config ClientConfig) *Client {
	// todo: add validation of config
	config.setDefaults()

	return &Client{
		Sender: &Sender{
			config:     config,
			httpClient: config.HttpClient,
		},
	}
}
