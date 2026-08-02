package vrage

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestParseResponseRejectsExtraJSONData(t *testing.T) {
	body := []byte(`{"data":{"result":"ok"},"meta":{"apiVersion":"v1","queryTime":1.23}}{"extra":true}`)
	response := &http.Response{Body: io.NopCloser(bytes.NewReader(body))}

	_, err := parseResponse[APIServerPingV1Data](response)
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

	_, err := parseResponse[APIServerPingV1Data](response)
	if err == nil {
		t.Fatal("expected parseResponse to return an error for HTTP status errors")
	}
	if got, want := err.Error(), "request failed with status 404 Not Found"; got != want {
		t.Fatalf("expected error %q, got %q", want, got)
	}
}

func TestDoUsesBaseEndpointForAuthHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		parts := strings.SplitN(auth, ":", 2)
		if len(parts) != 2 {
			t.Fatalf("expected authorization header to contain nonce and signature, got %q", auth)
		}

		nonce, signature := parts[0], parts[1]
		date := r.Header.Get("Date")
		payload := "/vrageremote/v1/server/ping\r\n" + nonce + "\r\n" + date + "\r\n"

		key, err := base64.StdEncoding.DecodeString("dGVzdA==")
		if err != nil {
			t.Fatalf("decode test key: %v", err)
		}

		mac := hmac.New(sha1.New, key)
		_, _ = mac.Write([]byte(payload))
		expectedSignature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
		if signature != expectedSignature {
			t.Fatalf("expected signature %q for payload %q, got %q", expectedSignature, payload, signature)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":{"result":"ok"},"meta":{"apiVersion":"v1","queryTime":1.23}}`))
	}))
	defer server.Close()

	u, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("parse server URL: %v", err)
	}
	port, err := strconv.Atoi(u.Port())
	if err != nil {
		t.Fatalf("parse server port: %v", err)
	}

	config := ClientConfig{
		RemoteApiIP:       u.Hostname(),
		RemoteApiPort:     uint32(port),
		RemoteSecurityKey: "dGVzdA==",
		BaseEndpoint:      toPtr(DefaultBaseEndpoint),
		HTTPClient:        server.Client(),
	}
	config.SetDefaults()

	client := httpClient{config: &config}
	response, err := client.Do(http.MethodGet, "/v1/server/ping", nil, nil)
	if err != nil {
		t.Fatalf("expected request to succeed, got %v", err)
	}
	defer response.Body.Close()
}
