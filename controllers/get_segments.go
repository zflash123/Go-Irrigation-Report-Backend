package controllers

import (
	"net/http"
	"go-irrigation-report-api/models"
	"github.com/google/uuid"
)
type CloseSegments struct {
	ID				uuid.UUID
	Name			string
	Geojson		interface{}
}
func GetCloseSegments(w http.ResponseWriter, r *http.Request) {
	latitude:= r.URL.Query().Get("lat")
	longitude:= r.URL.Query().Get("long")

	var closeSegments []CloseSegments
}