package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Session structure
type Session struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Status  bool
	UserID  string
	Created time.Time
}
