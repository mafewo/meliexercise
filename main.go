package main

import (
	"log"
	"os"
	"time"

	"github.com/mafewo/meliexercise/config"
)

const version = "0.0.1Alpha"

var flagPath string

//TimeZone time zone location
var TimeZone *time.Location

func main() {

	/// Main
	router := getRouter()

	/// Read Config file
	iniPath := flagPath + "app.ini"
	log.Printf("Leyendo configuraciones desde : %s", iniPath)
	config.Read(iniPath)

	/// Run http/https Servers
	errs := Run(os.Getenv("PORT"), router, map[string]bool{"http": true, "https": true})

	/// This will run forever until channel receives error
	select {
	case err := <-errs:
		log.Printf("No se pudieron iniciar los servidores !!! (error: %s)", err)
	}
}
