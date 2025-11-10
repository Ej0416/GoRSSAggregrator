package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	if statusCode > 499 {
		log.Println("Responding with 5XX error: ", message)
	}

	RepondWithJSON(w,statusCode,errResponse{
		Error: message,
	})
}

func RepondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
		dat, err := json.Marshal(payload)
		if err != nil {
			log.Printf("Failed to marshal JSON response: %v", payload)
			return
		}

		w.Header().Add("Content-type","application/json")
		w.WriteHeader(statusCode)
		w.Write(dat)
}