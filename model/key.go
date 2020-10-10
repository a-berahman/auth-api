package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Key structure
type Key struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Key     []byte             `bson:"key"`
	Created time.Time          `bson:"created"`
}
