package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
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
	models.Db.Raw(`SELECT
            id,
            name,
            geojson,
            public.ST_Distance(geom, public.geography(public.ST_SetSRID(public.ST_MakePoint(?, ?), 4326))) AS distance
        FROM
            map.irrigations_segment
        WHERE
            public.ST_Distance(geom, public.geography(public.ST_SetSRID(public.ST_MakePoint(?, ?), 4326)))<=100
        ORDER BY
            distance;`, longitude, latitude, longitude, latitude).Scan(&closeSegments)

	err:= json.NewEncoder(w).Encode(closeSegments)
}