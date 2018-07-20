package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/mafewo/meliexercise/database/mongo"
	"github.com/mafewo/meliexercise/models"
	"github.com/mafewo/meliexercise/msj"
	mgo "gopkg.in/mgo.v2"
)

// GetWeatherByDay get weather by day
func GetWeatherByDay(w http.ResponseWriter, r *http.Request) {
	queryVals := r.URL.Query()
	day := queryVals.Get("day")
	session, collection, err := _connect()
	if err != nil {
		msj.Set(w, err.Error(), 500).ReturnJSON()
		return
	}
	defer session.Close()
	mw := &models.ModelWeather{
		Conn:       session,
		Collection: collection,
		Data:       nil,
	}
	dayint, _ := strconv.ParseInt(day, 10, 32)
	weathers, err := mw.Getday(int32(dayint))
	if err != nil {
		msj.Set(w, err.Error(), 404).ReturnJSON()
		return
	}
	_response(w, weathers)
}

// GetResumenWheather get and generate resume to weather from 10 years
func GetResumenWheather(w http.ResponseWriter, r *http.Request) {
	session, collection, err := _connect()
	if err != nil {
		msj.Set(w, err.Error(), 500).ReturnJSON()
		return
	}
	defer session.Close()
	mw := &models.ModelWeather{
		Conn:       session,
		Collection: collection,
		Data:       nil,
	}
	weathers, err := mw.GetAll()
	if err != nil {
		msj.Set(w, err.Error(), 404).ReturnJSON()
		return
	}
	resumen := CalculateResumen(weathers)
	_response(w, resumen)
}

// CalculateWeather calulate the weather to Solar System
func CalculateWeather(sliceSS []models.SolarSystem) error {
	session, collection, err := _connect()
	if err != nil {
		return err
	}
	defer session.Close()
	mw := &models.ModelWeather{
		Conn:       session,
		Collection: collection,
		Data:       nil,
	}
	err = mw.DropCollection()
	if err != nil {
		return err
	}
	day := 0
	for _, ss := range sliceSS {
		day++
		weatherData := models.Weather{}
		weatherData.Day = day

		if DroughtDay(ss) {
			weatherData.Estate = "Drougth"
		} else if OptimalDay(ss) {
			weatherData.Estate = "Optimal"
		} else if RainDay(ss) {
			weatherData.Estate = "Rain"
			weatherData.Perimeter = int(ss.Perimiter())
		} else {
			weatherData.Estate = "Unknown"
		}
		_, err = mw.Insert(weatherData)
		if err != nil {
			return err
		}
	}

	return nil
}

// CalculateResumen calulate the weather resumen to Solar System from to 10 years
func CalculateResumen(weathers []models.Weather) map[string]interface{} {
	resumen := make(map[string]interface{})
	drougth := 0
	optimal := 0
	rain := 0
	unknown := 0
	day := 0
	max := weathers[35].Perimeter
	var maxdays []int
	for _, w := range weathers {
		day++
		switch w.Estate {
		case "Drougth":
			drougth++
		case "Optimal":
			optimal++
		case "Rain":
			if max == w.Perimeter {
				maxdays = append(maxdays, day)
			}
			rain++
		default:
			unknown++
		}
	}
	resumen["Drougth"] = drougth
	resumen["Optimal"] = optimal
	resumen["Rain"] = rain
	resumen["Unknown"] = unknown
	resumen["DaysStrom"] = maxdays
	return resumen
}

// OptimalDay return if they are parallels
func OptimalDay(ss models.SolarSystem) bool {
	return ss.TheyAreParallels()
}

// RainDay return if the sun is inside the triangle
func RainDay(ss models.SolarSystem) bool {
	return ss.TriangleContainSun()
}

// DroughtDay return if the planets are aligned with the center
func DroughtDay(ss models.SolarSystem) bool {
	if ss.TheyAreOnAxes() {
		return true
	}
	if ss.TheyAreParallels() {
		if ss.TheyPassThroughTheSun() {
			return true
		}
	}
	return false
}

func _connect() (*mgo.Session, *mgo.Collection, error) {
	//realizo la conexion a la DB
	session, collection, err := mongo.NewMG(
		"Melidb",
		"@melidb",
		"cluster0-shard-00-00-nc9ip.mongodb.net",
		"27017",
		"weather",
		"test",
	).InitializeDatabase()

	if err != nil {
		return nil, nil, errors.New("no se pudo conectar a la base de datos")
	}
	return session, collection, nil
}
