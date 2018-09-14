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
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name        string        `bson:"name,omitempty" json:"name"`
	Description string        `bson:"description,omitempty" json:"description"`
	Email       string        `bson:"email,omitempty" json:"email"`
	Password    string        `bson:"password,omitempty" json:"password"`
	Status      int           `bson:"status,omitempty" json:"status"`
	Admin       bool          `bson:"admin,omitempty" json:"admin"`
	Avatar      string        `bson:"avatar,omitempty" json:"avatar"`
	CreatedAt   time.Time     `bson:"created_at,omitempty" json:"createdAt"`
	UpdatedAt   time.Time     `bson:"updated_at,omitempty" json:"updatedAt"`
}

type Password struct {
	Password          string `json:"password" binding:"required,min=1,max=16"`
	RePassword        string `json:"re_password" binding:"required,min=8,max=16"`
	ReConfirmPassword string `json:"re_confirm_password" binding:"eqfield=RePassword"`
}

type Update struct {
	Name        string        `json:"name" binding:"min=4,max=12"`
	Description string        `json:"description" binding:"min=1,max=255"`
	Status      int           `json:"status"`
	UpdatedAt   time.Time     `bson:"updated_at,omitempty" json:"updatedAt"`
}
