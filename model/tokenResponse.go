package model

// TokenResponse is token response structure
// swagger:response TokenResponse
type TokenResponse struct {
	// the genration token
	Token string `json:"token"`
}
