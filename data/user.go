package data

import (
	"context"
	"fmt"

	"github.com/a-berahman/auth-api/config"
	"github.com/a-berahman/auth-api/model"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func GetUserID(mobile string, password string) (string, error) {
	filter := bson.M{"mobile": mobile}
	u := config.Users.FindOne(context.Background(), filter)
	currUser := model.User{}
	if err := u.Decode(&currUser); err != nil {
		return "", fmt.Errorf("your authentication information didn't match")
	}
	err := bcrypt.CompareHashAndPassword(currUser.Password, []byte(password))
	if err != nil {
		return "", fmt.Errorf("your authentication information didn't match")
	}

	return currUser.ID.Hex(), nil
}
