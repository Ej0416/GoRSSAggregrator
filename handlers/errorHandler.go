package handlers

import (
	"net/http"

	"github.com/Ej0416/GoRSSAggregrator/helpers"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithError(w, 400, "Something went wrong")
}