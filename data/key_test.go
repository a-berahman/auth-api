package data

import (
	"log"
	"os"
	"testing"
)

var testLog = log.New(os.Stdout, "auth-api-test-data/key", log.LstdFlags)

func Test_GenerateKey(t *testing.T) {
	testLog.Println("[TEST] test data/key by GenerateKey and GetKey")
	str, err := GenerateNewKey()
	if err != nil {
		t.Fatal(err)
	}
	if str == "" {
		t.Fatal("the result of generateKet was expected not to be empty")
	}
	key, err := GetKey(str)
	if err != nil {
		t.Fatal(err)
	}
	if key == nil {
		t.Fatal("key was expected not to be null")
	}
}
