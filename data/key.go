package data

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"time"

	"github.com/a-berahman/auth-api/config"
	"github.com/a-berahman/auth-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetKey returns key from Keys
func GetKey(kid string) ([]byte, error) {
	oid, _ := primitive.ObjectIDFromHex(kid)
	filter := bson.M{"_id": oid}
	key := model.Key{}
	result := config.Keys.FindOne(context.Background(), filter)
	err := result.Decode(&key)
	if err != nil {
		return nil, fmt.Errorf("face on error while is got key from keys")
	}
	return key.Key, nil
}

//GenerateNewKey prepares rotation key algorithm
func GenerateNewKey() (string, error) {
	newKey := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, newKey)
	if err != nil {
		return "", fmt.Errorf("Error in generateNewKey while generating key: %w", err)
	}

	// uid, err := uuid.NewV4()
	// if err != nil {
	// 	return fmt.Errorf("Error in generateNewKey while generating kid: %w", err)
	// }

	key := model.Key{
		ID:      primitive.NewObjectID(),
		Key:     newKey,
		Created: time.Now(),
	}

	result, err := config.Keys.InsertOne(context.Background(), &key)
	if err != nil {
		return "", fmt.Errorf("insertKey face on error : %w", err)
	}

	// currentKid = uid.String()
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("Can not convert to OID while key was generated")
	}
	return oid.Hex(), nil
}
