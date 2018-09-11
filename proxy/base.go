package proxy

import (
	"reflect"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var _ proxy = &baseProxy{}

type baseProxy struct {
	coll *mgo.Collection
}

func (p *baseProxy) Get(id string) (interface{}, error) {
	var i interface{}
	err := p.coll.FindId(bson.ObjectIdHex(id)).One(&i)
	return i, err
}

func (p *baseProxy) GetMany(query interface{}, page int, limit int) (interface{}, error) {
	var i []interface{}
	err := p.coll.Find(query).Limit(limit).Skip((page - 1) * limit).All(&i)
	return i, err
}

func (p *baseProxy) Create(body interface{}) error {
	v := reflect.ValueOf(body).Elem()
	now := time.Now()
	v.FieldByName("CreatedAt").Set(reflect.ValueOf(now))
	v.FieldByName("UpdatedAt").Set(reflect.ValueOf(now))
	v.FieldByName("ID").Set(reflect.ValueOf(bson.NewObjectId()))
	err := p.coll.Insert(body)
	return err
}

func (p *baseProxy) Update(id string, body interface{}) error {
	v := reflect.ValueOf(body).Elem()
	now := time.Now()
	v.FieldByName("UpdatedAt").Set(reflect.ValueOf(now))
	err := p.coll.UpdateId(bson.ObjectIdHex(id), bson.M{
		"$set": body,
	})
	return err
}

func (p *baseProxy) Delete(id string) error {
	err := p.coll.RemoveId(bson.ObjectIdHex(id))
	return err
}
