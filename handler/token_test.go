package handler

import (
	"log"
	"os"
	"testing"

	"github.com/a-berahman/auth-api/data"
)

var token string
var sessionID string
var tokenLog = log.New(os.Stdout, "auth-api-test-handler/token", log.LstdFlags)

func Test_createToken(t *testing.T) {
	tokenLog.Println("[TEST] test handler/token with createToken condition")
	uid, err := data.GetUserID("+989120729713", "123456")
	if err != nil {
		t.Fatal(err)
	}
	sid, err := data.CreateSession(uid)
	sessionID = sid
	if err != nil {
		t.Fatal(err)
	}
	token, err = createToken(sid, uid)
	if err != nil {
		t.Fatal(err)
		return
	}
	if token == "" {
		t.Fatal("token was excpected not to be empty")
	}

}

func Test_parseToken(t *testing.T) {
	tokenLog.Println("[TEST] test handler/token with parseToken condition")
	cc, err := parseToken(token)
	if err != nil {
		t.Fatal(err)
	}
	if cc.SID == "" {
		t.Fatal("session id was expected not to be empty")
	}
	if cc.SID != sessionID {
		t.Fatal("session id was expected to be equals by ", sessionID)
	}
}

func Test_parseTokenWillBeFailed(t *testing.T) {
	tokenLog.Println("[TEST] test handler/token with parseTokenWillBeFailed condition")
	_, err := parseToken("")
	if err == nil {
		t.Fatal("error was expected not to be null")
	}
}
