package mongodb

import (
	"qrapi/g/x/math"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ERR_NOT_FOUND = "not found"

var randMaker = func(prefix string) string {
	return math.RandString(prefix, 20)
}

type Table struct {
	Name   string
	Prefix string
	*mgo.Collection
}

func NewTable(name, prefix string) *Table {
	return &Table{
		Name:       name,
		Prefix:     prefix,
		Collection: database.C(name),
	}
}

func (t *Table) Create(model IModel) error {
	model.SetID(randMaker(t.Prefix))
	model.BeforeCreate()
	return t.Insert(model)
}

func (t *Table) CreateAuth(model IModel) error {
	model.SetID(math.RandString(t.Prefix, 80))
	model.BeforeCreate()
	return t.Insert(model)
}

func (t *Table) DeleteByID(id string) error {
	return t.UpdateId(id, bson.M{"$set": bson.M{"dtime": 0}})
}

func (t *Table) FindAll(result interface{}) error {
	return t.Find(bson.M{"dtime": bson.M{"$ne": 0}}).All(result)
}

func (t *Table) FindWhere(query bson.M, result interface{}) error {
	query["dtime"] = bson.M{"$ne": 0}
	return t.Find(query).All(result)
}

func (t *Table) FindID(id string, result interface{}) error {
	var query = bson.M{}
	query["dtime"] = bson.M{"$ne": 0}
	query["_id"] = id
	return t.Find(query).One(result)
}

func (t *Table) FindOne(query bson.M, result interface{}) error {
	query["dtime"] = bson.M{"$ne": 0}
	return t.Find(query).One(result)
}
