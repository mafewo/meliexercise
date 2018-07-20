package models

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Weather props
type Weather struct {
	ID        bson.ObjectId `bson:"_id" json:"-"`
	Day       int
	Estate    string
	Perimeter int `bson:"perimeter" json:"-"`
}

// ModelWeather base struct
type ModelWeather struct {
	Conn       *mgo.Session
	Collection *mgo.Collection
	Data       Weathers
}

// Weathers is a collection of Wheather
type Weathers []Weather

// Insert weather
func (mw *ModelWeather) Insert(w Weather) (Weather, error) {
	w.ID = bson.NewObjectId()
	err := mw.Collection.Insert(w)
	if err != nil {
		return w, err
	}
	return w, nil
}

// GetAll weather
func (mw *ModelWeather) GetAll() ([]Weather, error) {
	var weathers []Weather
	err := mw.Collection.Find(nil).All(&weathers)
	if err != nil {
		return weathers, err
	}
	return weathers, nil
}

// Getday weather by day
func (mw *ModelWeather) Getday(c int32) ([]Weather, error) {
	var weathers []Weather
	err := mw.Collection.Find(bson.M{"day": c}).All(&weathers)
	if err != nil {
		return weathers, err
	}
	return weathers, nil
}

// GetMaxAll obteins a slice with
func (mw *ModelWeather) GetMaxAll() ([]Weather, error) {
	var weathers []Weather
	err := mw.Collection.Find(nil).Sort("{Perimeter:-1}").All(&weathers)
	if err != nil {
		return weathers, err
	}
	return weathers, nil
}
