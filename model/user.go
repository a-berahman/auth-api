package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User defines the structure for an API user
// swagger:model
type User struct {
	ID         primitive.ObjectID `json:"sid" bson:"_id,omitempty"`
	FirstName  string             `json:"firstName" bson:"firstName"`
	LastName   string             `json:"lastName" bson:"lastName"`
	Mobile     string             `json:"mobile" bson:"mobile"`
	Password   []byte             `json:"password" bson:"password"`
	Role       string             `json:"role" bson:"role"`
	Status     bool               `json:"status" bson:"status"`
	CreateDate string             `json:"-"`
}
