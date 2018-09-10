package routes

import (
	"os"

	"github.com/globalsign/mgo"
)

var userCollection *mgo.Collection
var topicCollection *mgo.Collection
var commentCollection *mgo.Collection

func init() {
	session, err := mgo.Dial(os.Getenv("MONGO"))
	if err != nil {
		panic(err)
	}
	userCollection = session.DB("goulang").C("user")
	topicCollection = session.DB("goulang").C("topic")
	commentCollection = session.DB("goulang").C("comment")
}
