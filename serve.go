package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mafewo/meliexercise/handler"
	"github.com/mafewo/meliexercise/middleware"
)

func getRouter() *mux.Router {
	// Middleware chain

	jj := middleware.Chain(middleware.JSONHeader)

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", handler.Index).Methods("OPTIONS", "GET")
	router.HandleFunc("/generate", jj(handler.GenerateDate)).Methods("OPTIONS", "GET")
	router.HandleFunc("/resumen", jj(handler.GetResumenWheather)).Methods("OPTIONS", "GET")
	router.HandleFunc("/weather", jj(handler.GetWeatherByDay)).Methods("OPTIONS", "GET")

	router.NotFoundHandler = http.HandlerFunc(handler.NotFound)
	router.Use(middleware.CORS)

	return router
}

//Run Servers
func Run(addr string, router *mux.Router, prt map[string]bool) chan error {

	errs := make(chan error)
	if prt["http"] {
		// Starting HTTP server
		go func() {
			log.Printf("Iniciando servicio HTTP en %s ...", addr)
			if err := http.ListenAndServe(addr, router); err != nil {
				errs <- err
			}

		}()
	}

	return errs
}
