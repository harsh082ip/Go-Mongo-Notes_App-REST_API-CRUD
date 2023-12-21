package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notes struct {
	Id    primitive.ObjectID `bson:"_id"`
	Title string             `bson:"title,omitempty"`
	Desc  string             `bson:"desc,omitempty"`
	Time  string             `bson:"time,omitempty"`
}
