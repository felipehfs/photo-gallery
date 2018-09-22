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

// FindByID retrieve the photo by ID
func (pd *PhotoDao) FindByID(id string) (*Photo, error) {
	var photo Photo
	err := pd.getCollection().Find(bson.M{"id": bson.ObjectIdHex(id)}).One(&photo)
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

// Remove drop the photo by ID and returns a error
func (pd *PhotoDao) Remove(id string) error {
	return pd.getCollection().Remove(bson.M{"id": bson.ObjectIdHex(id)})
}
