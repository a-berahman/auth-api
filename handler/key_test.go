package handler

import (
	"log"
	"os"
	"testing"
)

var h *KeyHandler

func init() {
	l := log.New(os.Stdout, "auth-api-test-handler/key", log.LstdFlags)
	h = NewKeyhHandler(l)
}
func Test_generateKey(t *testing.T) {
	h.l.Println("[TEST] test handler/key with generateKey condition")
	if CurrentKid == "" {
		h.l.Println("current id is empty")
	}

	h.GenerateKey()
	h.l.Println(CurrentKid)
	if CurrentKid == "" {
		t.Fatal("currentid was expected not to ne empty")
	}
}
