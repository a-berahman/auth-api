package model

// ValidateResponse is a validation returns structure in response
// swagger:response validateResponse
type ValidateResponse struct {
	// the sessionID of user
	SID string `json:"sid"`
}
