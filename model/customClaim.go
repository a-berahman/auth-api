package model

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//CustomClaims model structure
type CustomClaims struct {
	jwt.StandardClaims
	SID string
}

func (cc *CustomClaims) Valid() error {
	if !cc.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("Token has expired")
	}

	if cc.SID == "" {
		return fmt.Errorf("Invalid session ID")
	}
	//boro to jadvale ssesions ha bebein in hanuz active ya na
	return nil
}
