package vrage

import "net/http"

// region /v1/server

// GetV1ServerStatus fetches the current status of the server and returns the HTTP response.
//
//	GET /v1/server
func (c *HTTPClient) GetV1ServerStatus() (*http.Response, error) {
	return c.DoErr(http.MethodGet, "/v1/server", jsonMap(nil), httpHeaders(nil))
}

// DeleteV1Server stops the server and returns the HTTP response.
//
// returns 200 OK with empty body if the server was successfully stopped
//
//	DELETE /v1/server
func (c *HTTPClient) DeleteV1Server() (*http.Response, error) {
	return c.DoErr(http.MethodDelete, "/v1/server", jsonMap(nil), httpHeaders(nil))
}

// region /v1/server/ping

// GetV1ServerPing fetches a ping response from the server and returns the HTTP response.
//
//	GET /v1/server/ping
func (c *HTTPClient) GetV1ServerPing() (*http.Response, error) {
	return c.DoErr(http.MethodGet, "/v1/server/ping", jsonMap(nil), httpHeaders(nil))
}
