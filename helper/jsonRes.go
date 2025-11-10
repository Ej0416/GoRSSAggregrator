package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func repondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
		dat, err := json.Marshal(payload)
		if err != nil {
			log.Printf("Failed to marshal JSON response: %v", payload)
			return
		}

		w.Header().Add("Content-type","application/json")
		w.WriteHeader(statusCode)
		w.Write(dat)
}