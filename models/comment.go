package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// User struct
type Comment struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Content   string        `bson:"content,omitempty" json:"content"`
	TopicID   bson.ObjectId `bson:"topic_id,omitempty" json:"topicID"`
	UserID    bson.ObjectId `bson:"user_id,omitempty" json:"userID"`
	CreatedAt time.Time     `bson:"created_at,omitempty" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty" json:"updatedAt"`
}
