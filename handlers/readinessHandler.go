package handlers

import (
	"net/http"

	"github.com/Ej0416/GoRSSAggregrator/helpers"
)

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RepondWithJSON(w,200,struct{}{})
}