package main

import (
	"net/http"
	"github.com/GoRSSAggregrator/helper"
)

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	helper.RepondWithJSON(w,200,struct{}{})
}