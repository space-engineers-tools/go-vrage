package vrage

import (
	"net/http"
)

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

// Config returns a copy of the ClientConfig with the RemoteSecurityKey field cleared for security reasons.
func (c *Client) Config() ClientConfig {
	config := c.Sender.config
	config.RemoteSecurityKey = ""
	return config
}

// NewClient creates a new Client instance with the provided configuration.
func NewClient(config ClientConfig) (*Client, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	if config.HTTPClient == nil {
		config.HTTPClient = http.DefaultClient
	}

	return &Client{
		Sender: &Sender{
			config:     config,
			httpClient: config.HTTPClient,
		},
	}, nil
}
