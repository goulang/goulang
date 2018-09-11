package proxy

import (
	"reflect"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/goulang/goulang/models"
)

var _ Proxy = &UserProxy{}

type UserProxy struct {
	coll *mgo.Collection
}

func (p *UserProxy) Get(id string) (interface{}, error) {
	var user models.User
	err := p.coll.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (p *UserProxy) GetMany(query interface{}, page int, limit int) (interface{}, error) {
	var users []models.User
	err := p.coll.Find(query).Limit(limit).Skip((page - 1) * limit).All(&users)
	return users, err
}

func (p *UserProxy) Create(body interface{}) error {
	v := reflect.ValueOf(body).Elem()
	now := time.Now()
	v.FieldByName("CreatedAt").Set(reflect.ValueOf(now))
	v.FieldByName("UpdatedAt").Set(reflect.ValueOf(now))
	v.FieldByName("ID").Set(reflect.ValueOf(bson.NewObjectId()))
	err := p.coll.Insert(body)
	return err
}

func (p *UserProxy) Update(id string, body interface{}) error {
	v := reflect.ValueOf(body).Elem()
	now := time.Now()
	v.FieldByName("UpdatedAt").Set(reflect.ValueOf(now))
	err := p.coll.UpdateId(bson.ObjectIdHex(id), bson.M{
		"$set": body,
	})
	return err
}

func (p *UserProxy) Delete(id string) error {
	err := p.coll.RemoveId(bson.ObjectIdHex(id))
	return err
}
