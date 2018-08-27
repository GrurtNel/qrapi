package mongodb

import "time"

type Model struct {
	ID    string `bson:"_id" json:"id"`
	CTime int64  `bson:"ctime,omitempty" json:"ctime,omitempty"`
	MTime int64  `bson:"mtime,omitempty" json:"mtime,omitempty"`
	DTime int64  `bson:"dtime,omitempty" json:"dtime,omitempty"`
}

func (m *Model) SetID(id string) {
	m.ID = id
}

func (m *Model) BeforeCreate() {
	m.CTime = time.Now().Unix()
}

func (m *Model) GetID() string {
	return m.ID
}

func (m *Model) BeforeUpdate() {
	m.MTime = time.Now().Unix()

}

func (m *Model) BeforeDelete() {
	m.DTime = time.Now().Unix()
}

type IModel interface {
	BeforeCreate()
	BeforeUpdate()
	BeforeDelete()
	SetID(id string)
	GetID() string
}
