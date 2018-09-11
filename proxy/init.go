package proxy

import (
	"os"

	"github.com/globalsign/mgo"
)

type Proxy interface {
	Get(id string) (interface{}, error)
	GetMany(query interface{}, page int, limit int) (interface{}, error)
	Create(body interface{}) error
	Update(id string, body interface{}) error
	Delete(id string) error
}

var User *UserProxy

// var topicCollection *mgo.Collection
// var commentCollection *mgo.Collection

func init() {
	session, err := mgo.Dial(os.Getenv("MONGO"))
	if err != nil {
		panic(err)
	}
	User = &UserProxy{session.DB("goulang").C("user")}

	// topicCollection = session.DB("goulang").C("topic")
	// commentCollection = session.DB("goulang").C("comment")
}
