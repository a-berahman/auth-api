package handler

import (
	"log"

	"github.com/a-berahman/auth-api/data"
)

// KeyHandler is a handler
type KeyHandler struct {
	l *log.Logger
}

// NewKeyhHandler makes Key handler object
func NewKeyhHandler(l *log.Logger) *KeyHandler {
	return &KeyHandler{l}
}

//CurrentKid is a current key id for token message cryptography
var CurrentKid = ""

//GenerateKey generates rotation key algorithm
func (h *KeyHandler) GenerateKey() {
	currentKid, err := data.GenerateNewKey()
	if err != nil {
		h.l.Println("[Debug]", "KeyHandler", err)
		return
	}
	CurrentKid = currentKid
}
