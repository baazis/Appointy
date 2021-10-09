package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create Struct
type Users struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"mail" bson:"mail,omitempty"`
	Password string             `json:"pass" bson:"pass,omitempty"`
	// Posts    []string           `json:"posts" bson:"posts"`
}

type Post struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption   string             `json:"cap,omitempty" bson:"cap,omitempty"`
	URL       string             `json:"url" bson:"url,omitempty"`
	Timestamp time.Time          `json:"tstamp" bson:"tstamp,omitempty"`
}
