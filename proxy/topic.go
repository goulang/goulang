package proxy

import (
	"github.com/globalsign/mgo/bson"
)

var _ proxy = &topicProxy{}

type topicProxy struct {
	baseProxy
}

func (p *topicProxy) View(id string) error {
	err := p.coll.UpdateId(bson.ObjectIdHex(id), bson.M{
		"$inc": bson.M{
			"view": 1,
		},
	})
	return err
}
