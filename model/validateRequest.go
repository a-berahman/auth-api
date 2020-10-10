package model

// ValidateRequest is a token validate structure
// swagger:model
type ValidateRequest struct {
	// the token for the validate
	//
	// required: true
	Token string `json:"token"`
}
