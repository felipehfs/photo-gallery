package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Photo represents the pictures on the disk
type Photo struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"id,omitempty"`
	Title    string        `json:"title,"`
	Filename string        `json:"filename,omitempty"`
	Likes    uint          `json:"likes,omitempty"`
	Created  time.Time     `json:"created,omitempty"`
}
