package main

import (
	"log"
	"net/http"

	"./handler"
	"./manager"
)

func main() {
	log.Println("Starting HMM server.")
	log.Println("Loading configuration from file.")
	err := manager.CreateFileIfnExist("configuration.json")

	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	err = manager.CreateDirIfnExist("bots")

	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	configuration, err := manager.LoadConfiguration()

	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	if configuration.Port == "" || configuration.Password == "" {
		log.Println("Error: Non set configuration robust.")
		return
	}

	log.Println("Configuration loaded properly.")

	log.Println("Starting HMM http server.")

	mux := http.NewServeMux()

	mux.HandleFunc("/api", handler.HandleAPI)

	err = http.ListenAndServe(":"+configuration.Port, mux)

	if err != nil {
		log.Println("Error: ", err.Error())
	}
}
