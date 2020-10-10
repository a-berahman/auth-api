package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/a-berahman/auth-api/model"
)

var authTestHandler *AuthHandler

func init() {
	l := log.New(os.Stdout, "auth-api", log.LstdFlags)
	authTestHandler = NewAuthHandler(l)
}

var currentToken string

func Test_GetToken(t *testing.T) {
	tokenRequest, _ := json.Marshal(model.TokenRequest{Username: "+989120729713", Password: "123456"})
	req, err := http.NewRequest(http.MethodPost, "/v1/auth/token", bytes.NewBuffer(tokenRequest))
	req.Header.Add("Content-Type", "application/json")
	res := httptest.NewRecorder()
	if err != nil {
		t.Fatal(err)
	}

	curHandler := http.HandlerFunc(authTestHandler.GetToken)

	curHandler.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatal("status was expected to be OK")
	}
	tokenRes := model.TokenResponse{}
	if err := model.FromJSON(&tokenRes, res.Body); err != nil {
		t.Fatal(err)
	}
	currentToken = tokenRes.Token
	if currentToken == "" {
		t.Fatal("currentToken was expected not to be empty")
	}

}

func Test_ValidateToken(t *testing.T) {
	validationRequest, _ := json.Marshal(model.ValidateRequest{Token: currentToken})
	req, err := http.NewRequest(http.MethodPost, "/v1/auth/validate", bytes.NewBuffer(validationRequest))
	req.Header.Add("content-type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	curHandler := http.HandlerFunc(authTestHandler.ValidateToken)
	curHandler.ServeHTTP(res, req)
	if res.Code != http.StatusOK {
		t.Fatal("status was expected to be OK")

	}
	validateRes := model.ValidateResponse{}
	if err := model.FromJSON(&validateRes, res.Body); err != nil {
		t.Fatal(err)
	}

}
