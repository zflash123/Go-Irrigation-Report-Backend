package controllers

import(
	"encoding/json"
	"fmt"
	"net/http"
	"go-irrigation-report-backend/models"

	"github.com/gorilla/mux"
)

type Reports struct {
	Id									string
	CreatedAt						string
	DoneAt							string
	Level								string
	Note								string
	Status							string
	CenterPointJson			string
	IrrigationName			string
}

func GetReportById(w http.ResponseWriter, r *http.Request){
	report_id := mux.Vars(r)["id"]
	var reports []Reports
	queryReports:= models.Db.Table("report.report_list").
	Select("report.report_list.id", "report.report_list.created_at", "report.report_list.done_at", "report.report_segment.level", "report.report_segment.note", "report.status.name as status", "map.irrigations_segment.center_point_json").
	Joins("JOIN report.status ON report.status.id = report.report_list.status_id").
	Joins("JOIN report.report_segment ON report.report_segment.report_id = report.report_list.id").
	Joins("JOIN map.irrigations_segment ON map.irrigations_segment.id = report.report_segment.segment_id").
	Where("report.report_list.id = ?", report_id).Scan(&reports)

	if queryReports.Error != nil {
		fmt.Printf("%v", queryReports.Error)
	}
	err := json.NewEncoder(w).Encode(reports)
	if err != nil {
		fmt.Printf("%v", err)
	}
}