package vrage

// BaseResponse represents the base response structure of the API.
type BaseResponse[T any] struct {
	Data T    `json:"data"`
	Meta Meta `json:"meta"`
}

// Meta represents the metadata of the API response.
type Meta struct {
	APIVersion string  `json:"apiVersion"`
	QueryTime  float64 `json:"queryTime"`
}
