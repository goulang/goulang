package models

import (
	"encoding/gob"
	"time"

	"github.com/globalsign/mgo/bson"
)

func init() {
	gob.Register(QFile{})
}

type QFile struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"key" json:"key"`
	Hash      string        `bson:"hash" json:"hash"`
	FSize     int64         `bson:"fsize" json:"fsize"`
	Bucket    string        `bson:"bucket" json:"bucket"`
	CreatedAt time.Time     `bson:"created_at" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updatedAt"`
}
