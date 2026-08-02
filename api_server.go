package vrage

// apiServer provides access to the /v1/server API routes.
type apiServer struct {
	http *httpClient
}

// region GET /v1/server/ping

type APIServerPingV1Data struct {
	Result string `json:"result"`
}

func (s *apiServer) PingV1() (BaseResponse[APIServerPingV1Data], error) {
	var pingV1Response BaseResponse[APIServerPingV1Data]
	httpResponse, err := s.http.V1ServerPing()
	if err != nil {
		return pingV1Response, err
	}

	pingV1Response, err = parseResponse[APIServerPingV1Data](httpResponse)
	if err != nil {
		return pingV1Response, err
	}

	return pingV1Response, nil
}

// region GET /v1/server

type ApiServerStatusResponse = BaseResponse[APIServerStatusData]
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

// func (s *APIServer) Status() (JSON, error) {
// 	return JSON{}, nil
// }

// region DELETE /v1/server

// // todo: stop
// func (s *APIServer) Stop() (JSON, error) {
// 	return JSON{}, nil
// }
