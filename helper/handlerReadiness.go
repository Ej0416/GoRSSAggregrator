package main

import (
	"net/http"
)

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	repondWithJSON(w,200,struct{}{})
}