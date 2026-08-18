package vrage

import "net/http"

// region /v1/session
// todo: implement remaining

// region /v1/session/players

// GetV1SessionPlayers fetches the list of players in the session and returns the HTTP response.
//
//	GET /v1/session/players
func (c *HTTPClient) GetV1SessionPlayers() (*http.Response, error) {
	return c.DoErr(http.MethodGet, "/v1/session/players", nil, httpHeaders(nil))
}

// region /v1/session/chat

// GetV1SessionChat fetches chat messages and returns the HTTP response.
//
//	GET /v1/session/chat
func (c *HTTPClient) GetV1SessionChat() (*http.Response, error) {
	return c.DoErr(http.MethodGet, "/v1/session/chat", nil, httpHeaders(nil))
}

// PostV1SessionChat sends a chat message and returns the HTTP response.
//
//	POST /v1/session/chat
func (c *HTTPClient) PostV1SessionChat(message string) (*http.Response, error) {
	return c.DoErr(http.MethodPost, "/v1/session/chat", message, httpHeaders(nil))
}
