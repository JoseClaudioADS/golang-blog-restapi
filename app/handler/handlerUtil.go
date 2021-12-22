package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot encode json. err=%v\n", err)
	}
}
