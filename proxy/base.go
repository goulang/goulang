package proxy

import (
	"reflect"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type proxy interface {
	Get(id string) (interface{}, error)
	GetOne(query interface{}) (interface{}, error)
	GetMany(query interface{}, page int, limit int) (interface{}, error)
	Create(body interface{}) error
	Update(id string, body interface{}) error
	Delete(id string) error
	Count(query interface{}) (int, error)
}

var _ proxy = &baseProxy{}

type baseProxy struct {
	coll  *mgo.Collection
	model reflect.Type
}

func (p *baseProxy) Get(id string) (interface{}, error) {
	var data = reflect.New(p.model).Interface()
	err := p.coll.FindId(bson.ObjectIdHex(id)).One(data)
	return reflect.ValueOf(data).Elem().Interface(), err
}

func (p *baseProxy) GetOne(query interface{}) (interface{}, error) {
	var data = reflect.New(p.model).Interface()
	err := p.coll.Find(query).One(data)
	return reflect.ValueOf(data).Elem().Interface(), err
}

func (p *baseProxy) GetMany(query interface{}, page int, limit int) (interface{}, error) {
	// fmt.Println(strconv.Itoa(page))

	var data = reflect.MakeSlice(reflect.SliceOf(p.model), 0, 10).Interface()
	err := p.coll.Find(query).Sort("-_id").Limit(limit).Skip((page - 1) * limit).All(&data)
	return data, err
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

func (p *baseProxy) Count(query interface{}) (int, error) {
	n, err := p.coll.Find(query).Count()
	return n, err
}
