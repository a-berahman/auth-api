package model

// TokenRequest is token request structure
// swagger:model
type TokenRequest struct {
	// the username for the user
	//
	// required: true
	Username string `json:"username"`
	// the password for the user
	//
	// required: true
	Password string `json:"password"`
}
