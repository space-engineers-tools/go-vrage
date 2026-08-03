package vrage

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

// httpMethod represents an HTTP method.
// https://pkg.go.dev/net/http#MethodGet
type httpMethod = string

// jsonMap represents a jsonMap object returned by the API.
type jsonMap = map[string]any

// httpHeaders represents HTTP headers to be sent with the request.
type httpHeaders = map[string]string

// todo: HTTPClient is not responsible for parsing the response body to structs.
// todo: HTTPClient sends requests and returns the response body as a map and error.
// todo: parsing the responses should be done by in the Client methods.

// HTTPClient handles all HTTP requests.
type HTTPClient struct {
	config *ClientConfig
}

// todo HTTPClient needs methods:
// .Do()
// handle the space engineers hmac processing
// .SendRequest(method Method, path string, json, headers) (map[string]any, error)

// buildAuthHeaders creates VRage custom auth headers.
// Signature payload format:
//
//	"{endpoint}\r\n{nonce}\r\n{date}\r\n"

func attemptUnmarshal[T any](data []byte) (T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	return result, err
}

func uuidv4WithoutDashes() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func buildAuthHeaders(remotesecuritykey string, endpoint string) (dateHeader, authHeader string) {
	uuidFirst := uuidv4WithoutDashes()
	uuidSecond := uuidv4WithoutDashes()
	nonce := uuidFirst + uuidSecond

	date := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05")
	payload := endpoint + "\r\n" + nonce + "\r\n" + date + "\r\n"

	key := []byte(remotesecuritykey)
	if decoded, err := base64.StdEncoding.DecodeString(remotesecuritykey); err == nil {
		key = decoded
	}

	mac := hmac.New(sha1.New, key)
	_, _ = mac.Write([]byte(payload))

	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return date, nonce + ":" + signature
}

func parseResponse[T any](response *http.Response) (BaseResponse[T], error) {
	var target BaseResponse[T]

	defer response.Body.Close()

	// todo: improve error handling for HTTP status codes; map to custom error types
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		if response.Status != "" {
			return target, fmt.Errorf("request failed with status %s", response.Status)
		}
		return target, fmt.Errorf("request failed with status %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return target, err
	}

	target, err = attemptUnmarshal[BaseResponse[T]](data)
	if err != nil {
		return target, err
	}

	return target, nil
}

func buildURL(config *ClientConfig, endpoint string) string {
	var protocol string
	if config.UseHTTPS {
		protocol = "https://"
	} else {
		protocol = "http://"
	}

	return fmt.Sprintf("%s%s:%d%s%s",
		protocol,
		config.RemoteApiIP,
		config.RemoteApiPort,
		fromPtr(config.BaseEndpoint),
		endpoint,
	)
}

func (c *HTTPClient) Do(method httpMethod, endpoint string, jsonPayload jsonMap, headers httpHeaders) (*http.Response, error) {
	url := buildURL(c.config, endpoint)

	// Build a request without dispatching it yet
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	if jsonPayload != nil {
		payloadBytes, err := json.Marshal(jsonPayload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal JSON payload: %w", err) // todo: add sentinel
		}
		request.Body = io.NopCloser(strings.NewReader(string(payloadBytes)))
		request.Header.Set("Content-Type", "application/json")
	}

	requestPath := fromPtr(c.config.BaseEndpoint) + endpoint
	dateHeader, authHeader := buildAuthHeaders(c.config.RemoteSecurityKey, requestPath)
	request.Header.Set("Date", dateHeader)
	request.Header.Set("Authorization", authHeader)
	request.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := c.config.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
API routes below.

Naming convention:
func {Method}{Version}{PathWithoutSlash}

Example:
func GetV1ServerPing()
*/

// region /v1/server

// GetV1ServerPing performs a request and returns the HTTP response
//
//	GET /v1/server/ping
func (c *HTTPClient) GetV1ServerPing() (*http.Response, error) {
	return c.Do(http.MethodGet, "/v1/server/ping", jsonMap(nil), httpHeaders(nil))
}

// GetV1ServerStatus performs a request and returns the HTTP response
//
//	GET /v1/server
func (c *HTTPClient) GetV1ServerStatus() (*http.Response, error) {
	return c.Do(http.MethodGet, "/v1/server", jsonMap(nil), httpHeaders(nil))
}

// DeleteV1Server performs a request and returns the HTTP response
//
//	DELETE /v1/server
func (c *HTTPClient) DeleteV1Server() (*http.Response, error) {
	return c.Do(http.MethodDelete, "/v1/server", jsonMap(nil), httpHeaders(nil))
}
