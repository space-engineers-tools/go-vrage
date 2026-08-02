package vrage

// BaseResponse represents the base response structure of the API.
type BaseResponse[T any] struct {
	Data T    `json:"data"`
	Meta Meta `json:"meta"`
}

// Meta represents the metadata of the API response.
type Meta struct {
	// APIVersion is the version of the API.
	APIVersion string `json:"apiVersion"`
	// QueryTime is the time taken (in seconds) to process the request.
	QueryTime float64 `json:"queryTime"`
}
