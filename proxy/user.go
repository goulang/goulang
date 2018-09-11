package proxy

import (
	"github.com/globalsign/mgo/bson"
)

var _ proxy = &userProxy{}

type userProxy struct {
	baseProxy
}

func (p *userProxy) Login(name, password string) bool {
	err := p.coll.Find(bson.M{
		"name":     name,
		"password": password,
	}).One(nil)
	if err != nil {
		return false
	}
	return true
}
