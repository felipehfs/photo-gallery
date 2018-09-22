package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PhotoDao represents the database transactions
type PhotoDao struct {
	DB *mgo.Session
}

func (pd *PhotoDao) getCollection() *mgo.Collection {
	return pd.DB.DB("photosgallery").C("pictures")
}

// Insert push the new data into database
func (pd *PhotoDao) Insert(p Photo) error {
	return pd.getCollection().Insert(p)
}

// Read retrieves all photos
func (pd *PhotoDao) Read() ([]Photo, error) {
	var photos []Photo
	err := pd.getCollection().Find(bson.M{}).All(&photos)
	if err != nil {
		return nil, err
	}
	return photos, nil
}
