package database

import (
	"github.com/haiyiyun/mongodb"
	"github.com/haiyiyun/mongodb/database"
	"go.mongodb.org/mongo-driver/bson"
)

type Database struct {
	*database.Database `json:"-" bson:"-" map:"-"`
}

func NewDatabase(mgo mongodb.Mongoer, col bson.M) *Database {
	mdl := &Database{
		Database: database.NewDatabase(mgo, col),
	}

	return mdl
}
