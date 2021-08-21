package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Collection1 struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty" map:"_id,omitempty"`
	Status     int                `json:"status" bson:"status" map:"status"`
	CreateTime time.Time          `json:"create_time" bson:"create_time" map:"create_time"`
	UpdateTime time.Time          `json:"update_time" bson:"update_time" map:"update_time"`
}
