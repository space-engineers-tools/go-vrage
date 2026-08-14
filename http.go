package vrage

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

// httpMethod represents an HTTP method.
// https://pkg.go.dev/net/http#MethodGet
type httpMethod = string

// httpHeaders represents HTTP headers to be sent with the request.
type httpHeaders = map[string]string

// HTTPClient handles all HTTP requests.
type HTTPClient struct {
	config *ClientConfig
}

// region Helper functions

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

// IsResponseSuccessful checks if the HTTP response starts with 2.
func IsResponseSuccessful(resp *http.Response) bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

// parseResponse parses the HTTP response and returns the API response structure or an error.
func parseResponse[T any](response *http.Response) (T, error) {
	var target T

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return target, fmt.Errorf("failed to read response body: %w", err)
	}

	if response.StatusCode == http.StatusForbidden {
		return target, newErrAPIInvalidSecurityKey()
	}

	if !IsResponseSuccessful(response) {
		return target, newErrAPIUnexpectedCode(response.StatusCode, string(bodyBytes))
	}

	target, err = attemptUnmarshal[T](bodyBytes)
	var unmarshalTypeErr *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeErr) {
		return target, newErrAPIUnexpectedBody(string(bodyBytes))
	}
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

// region HTTPClient methods

// Do sends an HTTP request to the API with the specified
// method, endpoint, JSON payload, and headers
// and returns the pure HTTP response and error without any wrapping.
func (c *HTTPClient) Do(
	method httpMethod,
	endpoint string,
	payload any,
	headers httpHeaders,
) (*http.Response, error) {
	url := buildURL(c.config, endpoint)

	// Build a request without dispatching it yet
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	if payload != nil {
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, newErrAPIUnexpectedBody(fmt.Sprintf("failed to marshal payload: %v", err))
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

// DoErr sends an HTTP request to the API
// and returns the HTTP response or sentinel errors defined in errors.go.
func (c *HTTPClient) DoErr(
	method httpMethod,
	endpoint string,
	payload any,
	headers httpHeaders,
) (*http.Response, error) {
	response, err := c.Do(method, endpoint, payload, headers)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, newErrAPIRequestTimeout(c.config.Timeout)
		}

		return nil, newErrAPIConnectionFailed(err)
	}

	return response, nil
}

/*
API routes are implementd in:
- http_server.go
- http_admin.go
- http_session.go

Naming convention for methods:
func {Method}{Version}{PathWithoutSlash}

Example:
GET /v1/api/server/ping = func GetV1ServerPing()
GET /v1/session/asteroids/{entityId} = func GetV1SessionAsteroidsEntityId(entityId string)
*/
