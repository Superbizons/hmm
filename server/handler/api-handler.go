package handler

import "net/http"

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Server connection can be only by post method", 403)
		return
	}

}
