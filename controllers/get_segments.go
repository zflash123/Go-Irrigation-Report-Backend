package controllers

import (
	"net/http"
)

func GetCloseSegments(w http.ResponseWriter, r *http.Request) {
	latitude:= r.URL.Query().Get("lat")

}