package main

import (
	"net/http"
)

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	repondWithJSON(w,200,struct{}{})
}