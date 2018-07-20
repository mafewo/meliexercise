package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mafewo/meliexercise/models"
	"github.com/mafewo/meliexercise/msj"
)

//Planet is an interface
type Planet interface {
	Movement(int) error
}

// GenerateDate calculate and generate data to save in databases
func GenerateDate(w http.ResponseWriter, r *http.Request) {
	var err error
	var vulcan = &models.Vulcan{
		Speed:  5,
		Raduis: 1000,
		X:      0,
		Y:      1000,
	}
	var betazoid = &models.Betazoid{
		Speed:  3,
		Raduis: 2000,
		X:      0,
		Y:      2000,
	}
	var ferengi = &models.Ferengi{
		Speed:  1,
		Raduis: 500,
		X:      0,
		Y:      500,
	}
	var v Planet = vulcan
	var b Planet = betazoid
	var f Planet = ferengi

	sliceSS, err := CalculateMovement(v, b, f, vulcan, betazoid, ferengi)
	if err != nil {
		msj.Set(w, err.Error(), 500).ReturnJSON()
		return
	}

	err = CalculateWeather(sliceSS)
	if err != nil {
		msj.Set(w, err.Error(), 500).ReturnJSON()
		return
	}

	output, err := json.Marshal(sliceSS)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "aplication/json")
	w.Write(output)
}

// CalculateMovement is a function
func CalculateMovement(v, b, f Planet, vulcan *models.Vulcan, betazoid *models.Betazoid, ferengi *models.Ferengi) ([]models.SolarSystem, error) {

	var day = 0
	var err error
	var ss models.SolarSystem
	var sliceSS models.SolarSystems

	for day < 3650 {
		day++
		err = v.Movement(day)
		if err != nil {
			return nil, err
		}
		err = b.Movement(day)
		if err != nil {
			return nil, err
		}
		err = f.Movement(day)
		if err != nil {
			return nil, err
		}

		ss = models.SolarSystem{
			Sun:      models.Sun{X: 0, Y: 0},
			Ferengi:  *ferengi,
			Betazoid: *betazoid,
			Vulcan:   *vulcan,
		}
		sliceSS = append(sliceSS, ss)
	}

	return sliceSS, nil
}
