package data

import (
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
)

var sessionLog = log.New(os.Stdout, "auth-api-test-data/session", log.LstdFlags)

func Test_CreateSession(t *testing.T) {
	sessionLog.Println("[TEST] test data/session by createSession ")
	uid := uuid.New()
	oid, err := CreateSession(uid.String())
	if err != nil {
		t.Fatal(err)
	}
	if oid == "" {
		t.Fatal("oid was expected not to be empty")
	}
}

func Test_CreateSessionFail(t *testing.T) {
	sessionLog.Println("[TEST] test data/session by createSessionFail")
	oid, err := CreateSession("")
	if err == nil {
		t.Fatal("error was expected not to be null")
	}
	if oid != "" {
		t.Fatal("oid was expected not to be fill")
	}
}
