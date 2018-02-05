package Models

import "gopkg.in/mgo.v2/bson"

type Team struct {
	Id      bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Name string `json:"Name"`
}
