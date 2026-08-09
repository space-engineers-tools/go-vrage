package vrage

import (
	"log"
)

// APIServer provides access to the /v1/server API routes.
type APIServer struct {
	http *HTTPClient
}

// region GET /v1/server/ping

// APIServerPingData represents the data returned by the GET /v1/server/ping endpoint.
type APIServerPingData struct {
	Result string `json:"result"`
}

// Ping returns a ping response from the server, which can be used to check if the server is reachable and responding.
func (s *APIServer) Ping() (APIResponseWithData[APIServerPingData], error) {
	var responseStruct APIResponseWithData[APIServerPingData]
	httpResponse, err := s.http.GetV1ServerPing()
	if err != nil {
		return responseStruct, err
	}
	defer func() {
		if err := httpResponse.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	responseStruct, err = parseResponse[APIResponseWithData[APIServerPingData]](httpResponse)
	if err != nil {
		return responseStruct, err
	}

	return responseStruct, nil
}

// region GET /v1/server

// APIServerStatusData represents the data returned by the GET /v1/server endpoint.
type APIServerStatusData struct {
	Game              string  `json:"Game"`
	IsReady           bool    `json:"IsReady"`
	PirateUsedPCU     int     `json:"PirateUsedPCU"`
	Players           int     `json:"Players"`
	ServerID          int64   `json:"ServerId"`
	ServerName        string  `json:"ServerName"`
	SimSpeed          float64 `json:"SimSpeed"`
	SimulationCPULoad float64 `json:"SimulationCpuLoad"`
	TotalTime         int     `json:"TotalTime"`
	UsedPCU           int     `json:"UsedPCU"`
	Version           string  `json:"Version"`
	WorldName         string  `json:"WorldName"`
}

// Status returns the current status of the server. This includes information about the performance and the world.
func (s *APIServer) Status() (APIResponseWithData[APIServerStatusData], error) {
	var responseStruct APIResponseWithData[APIServerStatusData]
	httpResponse, err := s.http.GetV1ServerStatus()
	if err != nil {
		return responseStruct, err
	}
	defer func() {
		if err := httpResponse.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	responseStruct, err = parseResponse[APIResponseWithData[APIServerStatusData]](httpResponse)
	if err != nil {
		return responseStruct, err
	}

	return responseStruct, nil
}

// region DELETE /v1/server

// Stop stops the server.
// Some hosting providers may restart the server instead of stopping it, depending on their configuration.
//
// Use with caution.
func (s *APIServer) Stop() (APIResponseWithoutData, error) {
	var responseStruct APIResponseWithoutData
	httpResponse, err := s.http.DeleteV1Server()
	if err != nil {
		return responseStruct, err
	}
	defer func() {
		if err := httpResponse.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	responseStruct, err = parseResponse[APIResponseWithoutData](httpResponse)
	if err != nil {
		return responseStruct, err
	}
	return responseStruct, nil
}
