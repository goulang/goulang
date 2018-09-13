package proxy

import (
	"os"
	"reflect"

	"github.com/globalsign/mgo"
	"github.com/goulang/goulang/models"
)

var User *userProxy
var Topic *topicProxy
var Comment *commentProxy

func init() {
	session, err := mgo.Dial(os.Getenv("MONGO"))
	if err != nil {
		panic(err)
	}
	User = &userProxy{baseProxy{session.DB("goulang").C("user"), reflect.TypeOf((*models.User)(nil)).Elem()}}
	Topic = &topicProxy{baseProxy{session.DB("goulang").C("topic"), reflect.TypeOf((*models.Topic)(nil)).Elem()}}
	Comment = &commentProxy{baseProxy{session.DB("goulang").C("comment"), reflect.TypeOf((*models.Comment)(nil)).Elem()}}
}
