package controllers

import (
	"encoding/json"
	"fmt"
	"go-irrigation-report-backend/models"
	"net/http"

	"github.com/google/uuid"
)

type CloseSegments struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Geojson string    `json:"geojson"`
}

func GetCloseSegments(w http.ResponseWriter, r *http.Request) {
	latitude := r.URL.Query().Get("lat")
	longitude := r.URL.Query().Get("long")

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

	err := json.NewEncoder(w).Encode(closeSegments)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
