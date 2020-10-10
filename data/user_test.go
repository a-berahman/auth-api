package data

import (
	"log"
	"os"
	"testing"
)

var userLog = log.New(os.Stdout, "auth-api-test-data/user", log.LstdFlags)

func Test_getUser_userExist(t *testing.T) {
	userLog.Println("[TEST] test data/user with userExist condition")

	oid, err := GetUserID("+989120729713", "123456")

	if err != nil {
		t.Fatal(err)
		return
	}
	if oid == "" {
		t.Fatal("the user was not expected not to be found")
		return
	}

}
func Test_getUser_userNotExist(t *testing.T) {
	userLog.Println("[TEST] test data/user with userNotExist condition")
	oid, err := GetUserID("usernotexist", "1234567")
	if oid != "" {
		t.Fatal("the user was not expected to be found")
		return
	}
	if err != nil {
		return
	}

}
