package vrage

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)

const testSecurityKey = "dGVzdA=="

var errDialTCPConnectionRefused = errors.New("connection refused")

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return fn(request)
}

func TestParseResponseRejectsExtraJSONData(t *testing.T) {
	body := []byte(`{"data":{"result":"ok"},"meta":{"apiVersion":"v1","queryTime":1.23}}{"extra":true}`)
	response := &http.Response{Body: io.NopCloser(bytes.NewReader(body))}

	_, err := parseResponse[APIServerPingData](response)
	if err == nil {
		t.Fatal("expected parseResponse to reject extra JSON data")
	}
}

func TestParseResponseReturnsErrorForHTTPStatusError(t *testing.T) {
	body := []byte(`<html><body>Not Found</body></html>`)
	response := &http.Response{
		StatusCode: http.StatusNotFound,
		Status:     "404 Not Found",
		Body:       io.NopCloser(bytes.NewReader(body)),
	}

	_, err := parseResponse[APIServerPingData](response)
	if err == nil {
		t.Fatal("expected parseResponse to return an error for HTTP status errors")
	}
	if !errors.Is(err, ErrAPIUnexpectedCode) {
		t.Fatalf("expected ErrAPIUnexpectedCode, got %v", err)
	}
	// Ensure the error message includes the status code and the response body.
	if got := err.Error(); !strings.Contains(got, "status code 404") || !strings.Contains(got, "Not Found") {
		t.Fatalf("expected error to mention status code and body, got %q", got)
	}
}

func TestDoUsesBaseEndpointForAuthHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(resWriter http.ResponseWriter, req *http.Request) {
		auth := req.Header.Get("Authorization")
		parts := strings.SplitN(auth, ":", 2)
		if len(parts) != 2 {
			t.Fatalf("expected authorization header to contain nonce and signature, got %q", auth)
		}

		nonce, signature := parts[0], parts[1]
		date := req.Header.Get("Date")
		payload := "/vrageremote/v1/server/ping\r\n" + nonce + "\r\n" + date + "\r\n"

		key, err := base64.StdEncoding.DecodeString(testSecurityKey)
		if err != nil {
			t.Fatalf("decode test key: %v", err)
		}

		mac := hmac.New(sha1.New, key)
		_, _ = mac.Write([]byte(payload))
		expectedSignature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
		if signature != expectedSignature {
			t.Fatalf("expected signature %q for payload %q, got %q", expectedSignature, payload, signature)
		}

		resWriter.WriteHeader(http.StatusOK)
		_, _ = resWriter.Write([]byte(`{"data":{"result":"ok"},"meta":{"apiVersion":"v1","queryTime":1.23}}`))
	}))
	defer server.Close()

	parsedURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("parse server URL: %v", err)
	}
	port, err := strconv.Atoi(parsedURL.Port())
	if err != nil {
		t.Fatalf("parse server port: %v", err)
	}

	config := ClientConfig{
		RemoteApiIP:       parsedURL.Hostname(),
		RemoteApiPort:     uint32(port),
		RemoteSecurityKey: testSecurityKey,
		BaseEndpoint:      toPtr(DefaultBaseEndpoint),
		HTTPClient:        server.Client(),
	}
	config.SetDefaults()

	client := HTTPClient{config: &config}
	response, err := client.Do(http.MethodGet, "/v1/server/ping", nil, nil)
	if err != nil {
		t.Fatalf("expected request to succeed, got %v", err)
	}

	if err := response.Body.Close(); err != nil {
		log.Printf("failed to close response body: %v", err)
	}
}

func TestRouteMethodsMapDeadlineExceededToAPIRequestTimeout(t *testing.T) {
	config := ClientConfig{
		RemoteApiIP:       "203.0.113.10",
		RemoteApiPort:     8080,
		RemoteSecurityKey: testSecurityKey,
		Timeout:           123 * time.Millisecond,
		BaseEndpoint:      toPtr(DefaultBaseEndpoint),
		HTTPClient: &http.Client{
			Transport: roundTripperFunc(func(_ *http.Request) (*http.Response, error) {
				return nil, context.DeadlineExceeded
			}),
		},
	}
	config.SetDefaults()

	client := HTTPClient{config: &config}

	response, err := client.Do(http.MethodGet, "/v1/server/ping", nil, nil)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected Do to return context deadline exceeded, got %v", err)
	}
	if response != nil {
		if err := response.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}

	response, err = client.GetV1ServerPing()
	if err == nil {
		t.Fatal("expected route method to return an error")
	}
	if response != nil {
		if err := response.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}
	if !errors.Is(err, ErrAPIRequestTimeout) {
		t.Fatalf("expected route method to map deadline exceeded to ErrAPIRequestTimeout, got %v", err)
	}
	if got, want := err.Error(), "request timed out: the server did not respond in time: exceeded 123ms"; got != want {
		t.Fatalf("expected error %q, got %q", want, got)
	}
}

func TestRouteMethodsMapTransportErrorsToAPIConnectionFailed(t *testing.T) {
	config := ClientConfig{
		RemoteApiIP:       "203.0.113.10",
		RemoteApiPort:     8080,
		RemoteSecurityKey: testSecurityKey,
		BaseEndpoint:      toPtr(DefaultBaseEndpoint),
		HTTPClient: &http.Client{
			Transport: roundTripperFunc(func(_ *http.Request) (*http.Response, error) {
				return nil, fmt.Errorf("dial tcp 203.0.113.10:8080: connect: %w", errDialTCPConnectionRefused)
			}),
		},
	}
	config.SetDefaults()

	client := HTTPClient{config: &config}

	response, err := client.GetV1ServerPing()
	if response != nil {
		if err := response.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}
	if !errors.Is(err, ErrAPIConnectionFailed) {
		t.Fatalf("expected route method to map transport failures to ErrAPIConnectionFailed, got %v", err)
	}
	if got := err.Error(); !strings.Contains(
		got, "failed to connect to the server: connection refused or host not available") ||
		!strings.Contains(got, "connect: connection refused") {
		t.Fatalf("expected error to include connection-failure context and dial failure, got %q", got)
	}
}
