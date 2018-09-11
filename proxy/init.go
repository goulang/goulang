package proxy

import (
	"os"

	"github.com/globalsign/mgo"
)

type proxy interface {
	Get(id string) (interface{}, error)
	GetMany(query interface{}, page int, limit int) (interface{}, error)
	Create(body interface{}) error
	Update(id string, body interface{}) error
	Delete(id string) error
}

var User *userProxy
var Topic *topicProxy
var Comment *commentProxy

func init() {
	session, err := mgo.Dial(os.Getenv("MONGO"))
	if err != nil {
		panic(err)
	}
	User = &userProxy{baseProxy{session.DB("goulang").C("user")}}
	Topic = &topicProxy{baseProxy{session.DB("goulang").C("user")}}
	Comment = &commentProxy{baseProxy{session.DB("goulang").C("user")}}
}
