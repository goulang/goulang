package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Topic struct
type Topic struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title     string        `bson:"title,omitempty" json:"title"`
	Content   string        `bson:"content,omitempty" json:"content"`
	View      int           `bson:"view,omitempty" json:"view"`
	UserID    bson.ObjectId `bson:"user_id,omitempty" json:"userID"`
	CreatedAt time.Time     `bson:"created_at,omitempty" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty" json:"updatedAt"`
}
