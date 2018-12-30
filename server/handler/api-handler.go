package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"../basic"
	"../manager"
	"github.com/Superbizons/hmm/api"
	"github.com/tomasen/realip"
)

var Clients = make(map[string]*basic.Client)

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Server connection can be only by post method", 403)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var cmd api.AuthorizationCommand
	err := decoder.Decode(&cmd)

	if err != nil {
		log.Println("Error: ", err.Error())
		http.Error(w, "Error with decoding message.", 422)
		return
	}

	switch cmd.Command.Cmd {
	case "AuthorizationCommand":
		if cmd.Password != manager.Configuration.Password {
			http.Error(w, "Unauthorized.", 403)
			return
		}

		ip := realip.FromRequest(r)

		if Clients[ip] == nil {
			client := basic.NewClient(ip)
			Clients[client.IP] = client
			log.Printf("New Client ID: %v AND IP: %s - is registered!", client.ID, client.IP)
		}

		return
	default:
		http.Error(w, "Bad request.", 400)
		return
	}
}
