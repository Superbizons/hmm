package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/Superbizons/hmm/api"

	"./manager"
)

var (
	URL      string
	Password string
	Port     = 0
)

func init() {
	flag.StringVar(&URL, "url", "", "URL")
	flag.StringVar(&Password, "pass", "", "Password")
	flag.IntVar(&Port, "p", 0, "Port")
	flag.Parse()

	if URL == "" || Password == "" || Port == 0 {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	log.Println("Starting HMM client.")

	err := manager.CreateDirIfnExist("bots")

	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	cmd := api.AuthorizationCommand{&api.Command{"AuthorizationCommand"}, 5, Port, Password}

	file, err := manager.SendCommand(cmd, URL)

	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	r, err := zip.OpenReader(file.Name())
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}
	defer r.Close()

	packagename := strings.TrimSuffix(file.Name(), ".zip")

	err = manager.CreateDirIfnExist("bots/" + packagename)

	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	for _, f := range r.File {

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll("bots/"+packagename+"/"+f.Name, os.ModePerm)
		} else {
			fmt.Printf("Contents of %s:\n", f.Name)
			rc, err := f.Open()
			if err != nil {
				log.Println("Error: ", err.Error())
			}
			file, err := os.Create("bots/" + packagename + "/" + f.Name)

			if err != nil {
				log.Println("Error: ", err.Error())
			}

			_, err = io.Copy(file, rc)

			if err != nil {
				log.Println("Error: ", err.Error())
			}
			/*_, err = io.CopyN(os.Stdout, rc, 68)
			if err != nil {
				log.Println("Error: ", err.Error())
			}*/
			rc.Close()
			//fmt.Println()
		}
	}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", Port))
	if err != nil {
		log.Println("Error: ", err.Error())
	}

	conn, err := ln.Accept()

	if err != nil {
		log.Println("Error: ", err.Error())
	}
	var buff = make([]byte, 64000)

	value := 5

	for value != 0 {
		value, err := conn.Read(buff)

		if err != nil {
			log.Println("Error: ", err.Error())
			log.Println("Value: ", value)
			log.Println("Buffer: ", string(buff[:value]))
			break
		}
		if value != 0 {
			log.Println("Value: ", value)
			log.Println("Buffer: ", string(buff[:value]))
		}
	}

	log.Println("Value: ", value)
	log.Println("Buffer: ", string(buff[:value]))

	log.Println("Client successfully authorized!")
}
