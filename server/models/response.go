package models

// Response represents the REST response model
type Response struct {
	Data       interface{} `json:"data,omitempty"`     // primary data for the response
	Metadata   interface{} `json:"metadata,omitempty"` // associated metadata
	StatusCode uint16      `json:"statusCode"`         // status code eg. 200, 207, 400
	Message    string      `json:"message,omitempty"`  // a message associated with the response
}
