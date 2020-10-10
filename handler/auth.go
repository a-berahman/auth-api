package handler

import (
	"log"
	"net/http"

	"github.com/a-berahman/auth-api/data"
	"github.com/a-berahman/auth-api/model"
)

// AuthHandler us a Handler
type AuthHandler struct {
	l *log.Logger
}

// NewAuthHandler creates new object of authHandler
func NewAuthHandler(l *log.Logger) *AuthHandler {
	return &AuthHandler{l}
}

// swagger:route POST /v1/auth/token auth getToken
// Generate a token
//
// responses:
//	200: getTokenResponse
//	500: errorResponse
//	400: errorResponse
//	403: errorResponse

//GetToken handles POST requests to generate token
func (h *AuthHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	h.l.Println("[DEBUG] GetToken was ran")

	if r.Method != http.MethodPost {
		h.l.Println("[ERROR] method was not post")
		w.WriteHeader(http.StatusMethodNotAllowed)
		model.ToJSON(&model.GenericError{Message: "your method was not post"}, w)
		return
	}
	reqModel := &model.TokenRequest{}
	err := model.FromJSON(reqModel, r.Body)
	if err != nil {
		h.l.Println("[ERROR] deserialization face on error")
		w.WriteHeader(http.StatusBadRequest)
		model.ToJSON(&model.GenericError{Message: "post structure wasn't correct"}, w)
		return
	}

	h.l.Println("[DEBUG] GetUserID with username and password")
	uid, err := data.GetUserID(reqModel.Username, reqModel.Password)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		model.ToJSON(&model.GenericError{Message: err.Error()}, w)
		return
	}

	h.l.Println("[DEBUG] Generate new UUID as sessionID")
	sid, err := data.CreateSession(uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		model.ToJSON(&model.GenericError{Message: "generating unique id face on problem"}, w)
		return
	}

	//check kon bebin to jadvale user id be session in user chand ta session dare agar az hade mojaz kharej bud session ro delete kon

	h.l.Println("[DEBUG] createToken is fired")
	token, err := createToken(sid, uid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		model.ToJSON(&model.GenericError{Message: err.Error()}, w)
		return
	}

	h.l.Println("[DEBUG] takeing token and passing response to client")
	w.WriteHeader(http.StatusOK)
	model.ToJSON(&model.TokenResponse{Token: token}, w)
}

// swagger:route POST /v1/auth/validate auth validateToken
// Validation token
//
// responses:
//	200: validateResponse
//	400: errorResponse
//	403: errorResponse

//ValidateToken handles POST requests to validate token
func (h *AuthHandler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	h.l.Println("[DEBUG] ValidateToken was ran")
	if r.Method != http.MethodPost {
		h.l.Println("[ERROR] method was not post")
		w.WriteHeader(http.StatusMethodNotAllowed)
		model.ToJSON(&model.GenericError{Message: "your method was not post"}, w)
		return
	}
	reqModel := model.ValidateRequest{}
	err := model.FromJSON(&reqModel, r.Body)
	if err != nil {
		h.l.Println("[ERROR] deserialization face on error")
		w.WriteHeader(http.StatusBadRequest)
		model.ToJSON(&model.GenericError{Message: "post structure wasn't correct"}, w)
		return
	}

	c, err := parseToken(reqModel.Token)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		model.ToJSON(&model.GenericError{Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	model.ToJSON(&model.ValidateResponse{SID: c.SID}, w)

}
