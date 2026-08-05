package vrage

// Client is the main entry point for interacting with the VRage API. It provides access to various API endpoints through its sub-clients.
type Client struct {
	Config  ClientConfig
	HTTP    HTTPClient
	Session APISession
	Server  APIServer
	Admin   APIAdmin
}

// NewClient creates a new Client instance with the provided configuration.
func NewClient(config ClientConfig) (*Client, error) {
	config.SetDefaults()
	if err := config.Validate(); err != nil {
		return nil, err
	}
	// todo: use sentinels in errors.go

	httpClient := HTTPClient{config: &config}

	return &Client{
		Config:  config,
		HTTP:    httpClient,
		Session: APISession{http: &httpClient},
		Server:  APIServer{http: &httpClient},
		Admin:   APIAdmin{http: &httpClient},
	}, nil
}

// todo: migrate config validation to own solution to minimize dependencies
// todo: ping server to check if everything is setup correctly and the server is reachable.
