package handler

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/a-berahman/auth-api/data"
	"github.com/a-berahman/auth-api/model"
	"github.com/dgrijalva/jwt-go"
)

//CreateToken is generated token
func createToken(sid string, userID string) (string, error) {
	cc := model.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(120 * time.Minute).Unix(),
		},
		SID: sid,
	}

	//create token with using claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &cc)
	// if currentKeyId is null generate and fill again
	if CurrentKid == "" {
		refreshCurrentKeyID()
	}
	token.Header["kid"] = CurrentKid

	key, err := data.GetKey(CurrentKid)
	if err != nil {
		return "", fmt.Errorf("couldn't sign token in createToken %w", err)
	}
	//Get the complete, signed token
	res, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("couldn't sign token in createToken %w", err)
	}

	return res, nil
}

func refreshCurrentKeyID() {
	l := log.New(os.Stdout, "auth-api", log.LstdFlags)
	keyHandler := NewKeyhHandler(l)
	keyHandler.GenerateKey()
}

//ParseToken is parsing token with claims
func parseToken(str string) (*model.CustomClaims, error) {
	afterValidate, err := jwt.ParseWithClaims(str, &model.CustomClaims{}, func(beforeValidate *jwt.Token) (interface{}, error) {
		if beforeValidate.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("your token is fake")
		}

		kid, ok := beforeValidate.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("Invalid Key Id")
		}
		k, err := data.GetKey(kid)

		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}

		return k, nil
	})

	if err != nil {
		return nil, fmt.Errorf("ParseWithClaims face on %w", err)
	}

	if !afterValidate.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	//The second segment of the token
	//Using assertion and taking sessionID
	return afterValidate.Claims.(*model.CustomClaims), nil
}
