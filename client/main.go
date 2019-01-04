package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path"
	"strconv"
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

func playgame() {
	log.Println("Starting HMM client.")

	err := manager.CreateDirIfnExist("bots")

	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	cmd := api.AuthorizationCommand{&api.Command{"AuthorizationCommand"}, 5, Port, Password}

	log.Println("Sending command to server.")
	file, err := manager.SendCommand(cmd, "http://"+URL)

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

			rc.Close()
		}
	}

	log.Println("Client successfully authorized!")
	log.Println("All packages of bot are downloaded.")

	lis, err := net.Listen("tcp4", ":"+strconv.Itoa(Port))

	if err != nil {
		log.Println("Error: ", err.Error())
	}

	var conn net.Conn

	for {
		conn, err = lis.Accept()

		if err != nil {
			log.Println("Accept error:", err)
		}

		log.Println("accept:", conn.RemoteAddr())

		if addr, ok := conn.RemoteAddr().(*net.TCPAddr); ok {
			log.Println(addr.IP.String())

			if addr.IP.String() != strings.Split(URL, ":")[0] {
				log.Println("ADDR: ", addr.IP.String())
				log.Println("URL: ", strings.Split(URL, ":")[0])
			}

			fmt.Println("Success!")
			break
		}
	}
	lis.Close()
	gobinary, err := exec.LookPath("go")

	if err != nil {
		log.Println("Error: ", err.Error())
	}

	log.Println("GoBinary: ", gobinary)
	cmdGo := exec.Command(gobinary, "run", path.Join("bots", packagename, "MyBot.go"))
	pwd, err := os.Getwd()

	if err != nil {
		log.Println("Error: ", err.Error())
	}

	log.Println("PWD: ", pwd)

	cmdGo.Env = append(os.Environ(), "GOPATH="+path.Join(pwd, "bots", packagename))
	cmdGo.Stdin = conn
	cmdGo.Stdout = conn

	err = cmdGo.Run()

	if err != nil {
		log.Println("Error: ", err.Error())
	}
}

func main() {
	for {
		playgame()
	}
}
