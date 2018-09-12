package proxy

import (
	"os"

	"github.com/globalsign/mgo"
)

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
