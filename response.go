package vrage

// APIResponseWithData represents an API with a data field.
type APIResponseWithData[T any] struct {
	Data T       `json:"data"`
	Meta APIMeta `json:"meta"`
}

// APIResponseWithoutData represents an API response without a data field.
type APIResponseWithoutData struct {
	Meta APIMeta `json:"meta"`
}

// APIMeta represents the metadata of the API response.
type APIMeta struct {
	// APIVersion is the version of the API.
	APIVersion string `json:"apiVersion"`
	// QueryTime is the time taken (in seconds) to process the request.
	QueryTime float64 `json:"queryTime"`
}
