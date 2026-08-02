package vrage

// APIServer provides access to the /v1/server API routes.
type APIServer struct {
	http *HTTPClient
}

// region GET /v1/server/ping

type APIServerPingV1Data struct {
	Result string `json:"result"`
}

// PingV1 returns a ping response from the server, which can be used to check if the server is reachable and responding.
//
//	GET /v1/server/ping
func (s *APIServer) PingV1() (BaseResponse[APIServerPingV1Data], error) {
	var responseStruct BaseResponse[APIServerPingV1Data]
	httpResponse, err := s.http.V1ServerPing()
	if err != nil {
		return responseStruct, err
	}

	responseStruct, err = parseResponse[APIServerPingV1Data](httpResponse)
	if err != nil {
		return responseStruct, err
	}

	return responseStruct, nil
}

// region GET /v1/server

type APIServerStatusData struct {
	Game              string  `json:"Game"`
	IsReady           bool    `json:"IsReady"`
	PirateUsedPCU     int     `json:"PirateUsedPCU"`
	Players           int     `json:"Players"`
	ServerId          int64   `json:"ServerId"`
	ServerName        string  `json:"ServerName"`
	SimSpeed          float64 `json:"SimSpeed"`
	SimulationCpuLoad float64 `json:"SimulationCpuLoad"`
	TotalTime         int     `json:"TotalTime"`
	UsedPCU           int     `json:"UsedPCU"`
	Version           string  `json:"Version"`
	WorldName         string  `json:"WorldName"`
}

// StatusV1 returns the current status of the server. This includes information about the performance and the world.
//
//	GET /v1/server
func (s *APIServer) StatusV1() (BaseResponse[APIServerStatusData], error) {
	var responseStruct BaseResponse[APIServerStatusData]
	httpResponse, err := s.http.V1ServerStatus()
	if err != nil {
		return responseStruct, err
	}

	responseStruct, err = parseResponse[APIServerStatusData](httpResponse)
	if err != nil {
		return responseStruct, err
	}

	return responseStruct, nil
}

// func (s *APIServer) Status() (JSON, error) {
// 	return JSON{}, nil
// }

// region DELETE /v1/server

// // todo: stop
// func (s *APIServer) Stop() (JSON, error) {
// 	return JSON{}, nil
// }
