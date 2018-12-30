package main

import (
	"flag"
	"log"
	"os"

	"github.com/Superbizons/hmm/api"

	"./manager"
)

var (
	URL      string
	Password string
)

func init() {
	flag.StringVar(&URL, "url", "", "URL")
	flag.StringVar(&Password, "p", "", "Password")
	flag.Parse()

	if URL == "" || Password == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	log.Println("Starting HMM client.")

	cmd := api.AuthorizationCommand{&api.Command{"AuthorizationCommand"}, 5, Password}

	err := manager.SendCommand(cmd, URL)

	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	log.Println("Client successfully authorized!")
}
