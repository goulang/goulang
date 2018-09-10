package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Topic struct
type Topic struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Title     string        `bson:"title" json:"title"`
	Content   string        `bson:"content" json:"content"`
	View      int           `bson:"view" json:"view"`
	UserID    bson.ObjectId `bson:"user_id" json:"userID"`
	CreatedAt time.Time     `bson:"created_at" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updatedAt"`
}
