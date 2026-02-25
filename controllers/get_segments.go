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

type Segment struct {
	SegmentID 					uuid.UUID `json:"id"`
	IrrigationName			string    `json:"irrigation"`
	Geojson 						string    `json:"geojson"`
	Level								string		`json:"level"`
	Status							string		`json:"status"`
	Canal								string		`json:"canal"`
	Image								string		`json:"image"`
}

func GetCloseSegments(w http.ResponseWriter, r *http.Request) {
	latitude := r.URL.Query().Get("lat")
	longitude := r.URL.Query().Get("long")

	if latitude == "" || longitude== ""{
		var res Response
		res.Message = "latitude or longitude query param is required"
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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

	if closeSegments == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(closeSegments)
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func GetSegmentsByUserId(w http.ResponseWriter, r *http.Request) {
	user_id := fmt.Sprintf("%v", r.Context().Value("user_id"))

	var segments []Segment
	query := models.Db.Table("report.report_list").Select("report.status.name as status", "report.report_segment.segment_id", "report.report_segment.level", "map.irrigations_segment.geojson", "map.irrigations.name as irrigation_name", "map.irrigations.type as canal", "file.upload_dump.file_url as image").
	Joins("JOIN report.status ON report.status.id = report.report_list.status_id").
	Joins("JOIN report.report_segment ON report.report_segment.report_id = report.report_list.id").
	Joins("JOIN report.report_photo ON report.report_photo.report_segment_id = report.report_segment.id").
	Joins("JOIN file.upload_dump ON file.upload_dump.id = report.report_photo.upload_dump_id").
	Joins("JOIN map.irrigations_segment ON map.irrigations_segment.id = report.report_segment.segment_id").
	Joins("JOIN map.irrigations ON map.irrigations.id = map.irrigations_segment.irrigation_id").
	Where("report.report_list.user_id = ?", user_id).Scan(&segments)

	if query.Error != nil {
		var res Response
		res.Message = "There is an error when executing the query."
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			fmt.Printf("%v", err)
		}
	}
	err := json.NewEncoder(w).Encode(segments)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
