package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type notes struct {
	Id    primitive.ObjectID `bson:"_id"`
	title string             `bson:"title,omitempty"`
	desc  string             `bson:"desc,omitempty"`
	time  string             `bson:"time,omitempty"`
}

func golang() {

}

// func (n notes) Notes() notes {

// 	return n.Notes()
// }
