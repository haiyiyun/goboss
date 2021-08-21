package collection1

import (
	"goboss/internal/app/app1/database"
	"goboss/internal/app/app1/database/schema"

	"github.com/haiyiyun/mongodb"
)

type Model struct {
	*database.Database `json:"-" bson:"-" map:"-"`
}

func NewModel(mgo mongodb.Mongoer) *Model {
	obj := &Model{
		Database: database.NewDatabase(mgo, schema.Collection1),
	}

	return obj
}
