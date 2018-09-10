package models

import (
	"encoding/gob"
	"time"

	"github.com/globalsign/mgo/bson"
)

func init() {
	gob.Register(User{})
}

// User struct
type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	Avatar      string        `bson:"avatar" json:"avatar"`
	Email       string        `bson:"email" json:"email"`
	Password    string        `bson:"password" json:"password"`
	CreatedAt   time.Time     `bson:"created_at" json:"createdAt"`
	UpdatedAt   time.Time     `bson:"updated_at" json:"updatedAt"`
}
