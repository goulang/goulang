package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// User struct
type Comment struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Content   string        `bson:"content" json:"content"`
	UserID    bson.ObjectId `bson:"user_id" json:"userID"`
	CreatedAt time.Time     `bson:"created_at" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updatedAt"`
}
