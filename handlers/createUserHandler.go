package handlers

import (
	"net/http"

	"github.com/Ej0416/GoRSSAggregrator/helpers"
	"github.com/Ej0416/GoRSSAggregrator/internal/config"
)

type ExtendedApiConfig struct {
    *config.ApiConfig
}


func (e *ExtendedApiConfig)CreateUserHandler(w http.ResponseWriter, r *http.Request)  {
	helpers.RepondWithJSON(w,200,struct{}{})
}