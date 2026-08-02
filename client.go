package vrage

type Client struct {
	Config  ClientConfig
	HTTP    httpClient
	Session apiSession
	Server  apiServer
	Admin   apiAdmin
}

// NewClient creates a new Client instance with the provided configuration.
func NewClient(config ClientConfig) (*Client, error) {
	config.SetDefaults()
	if err := config.Validate(); err != nil {
		return nil, err
	}

	httpClient := httpClient{config: &config}

	return &Client{
		Config:  config,
		HTTP:    httpClient,
		Session: apiSession{http: &httpClient},
		Server:  apiServer{http: &httpClient},
		Admin:   apiAdmin{http: &httpClient},
	}, nil
}

// todo: migrate config validation to own solution to minimize dependencies
// todo: ping server to check if everything is setup correctly and the server is reachable.
