package controllers

import(
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"math/rand"
	"go-irrigation-report-backend/models"

	"github.com/gorilla/mux"
)

type Reports struct {
	Id              string `json:"id"`
	CreatedAt       string `json:"created_at"`
	DoneAt          string `json:"done_at"`
	Level           string `json:"level"`
	Note            string `json:"note"`
	Status          string `json:"status"`
	CenterPointJson string `json:"center_point_json"`
	IrrigationName  string `json:"irrigation_name"`
	Canal           string `json:"canal"`
	Image           string `json:"image"`
}

func GetReportById(w http.ResponseWriter, r *http.Request){
	report_id := mux.Vars(r)["id"]
	var reports []Reports
	queryReports := models.Db.Table("report.report_list").
		Select("report.report_list.id", "report.report_list.created_at", "report.report_list.done_at", "report.report_segment.level", "report.report_segment.note", "report.status.name as status", "map.irrigations_segment.center_point_json", "map.irrigations.name as irrigation_name", "map.irrigations.type as canal", "report.report_photo.file_url as image").
		Joins("JOIN report.status ON report.status.id = report.report_list.status_id").
		Joins("JOIN report.report_segment ON report.report_segment.report_id = report.report_list.id").
		Joins("JOIN report.report_photo ON report.report_photo.id = report.report_segment.report_photo_id").
		Joins("JOIN map.irrigations_segment ON map.irrigations_segment.id = report.report_segment.segment_id").
		Joins("JOIN map.irrigations ON map.irrigations.id = map.irrigations_segment.irrigation_id").
		Where("report.report_list.id = ?", report_id).Scan(&reports)

	if queryReports.Error != nil {
		fmt.Printf("%v", queryReports.Error)
	}
	err := json.NewEncoder(w).Encode(reports)
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func GetReportByUserId(w http.ResponseWriter, r *http.Request){
	user_id := fmt.Sprintf("%v", r.Context().Value("user_id"))
	filter := r.URL.Query().Get("filter")
	search := r.URL.Query().Get("search")

	var reports []Reports
	var queryReports = models.Db
	if(filter != "" && search != ""){
		queryReports = models.Db.Table("report.report_list").
		Select("DISTINCT ON (report.report_list.id) report.report_list.id", "report.report_list.created_at", "report.report_list.done_at", "report.status.name as status", "map.irrigations.name as irrigation_name", "map.irrigations.type as canal").
		Joins("JOIN report.status ON report.status.id = report.report_list.status_id").
		Joins("JOIN report.report_segment ON report.report_segment.report_id = report.report_list.id").
		Joins("JOIN map.irrigations_segment ON map.irrigations_segment.id = report.report_segment.segment_id").
		Joins("JOIN map.irrigations ON map.irrigations.id = map.irrigations_segment.irrigation_id").
		Where("report.report_list.user_id = ? AND report.status.name = ? AND map.irrigations.name = ?", user_id, filter, search).
		Scan(&reports)
	} else if filter != ""{
		queryReports = models.Db.Table("report.report_list").
		Select("DISTINCT ON (report.report_list.id) report.report_list.id", "report.report_list.created_at", "report.report_list.done_at", "report.status.name as status", "map.irrigations.name as irrigation_name", "map.irrigations.type as canal").
		Joins("JOIN report.status ON report.status.id = report.report_list.status_id").
		Joins("JOIN report.report_segment ON report.report_segment.report_id = report.report_list.id").
		Joins("JOIN map.irrigations_segment ON map.irrigations_segment.id = report.report_segment.segment_id").
		Joins("JOIN map.irrigations ON map.irrigations.id = map.irrigations_segment.irrigation_id").
		Where("report.report_list.user_id = ? AND report.status.name = ?", user_id, filter).
		Scan(&reports)
	} else if search != ""{
		queryReports = models.Db.Table("report.report_list").
		Select("DISTINCT ON (report.report_list.id) report.report_list.id", "report.report_list.created_at", "report.report_list.done_at", "report.status.name as status", "map.irrigations.name as irrigation_name", "map.irrigations.type as canal").
		Joins("JOIN report.status ON report.status.id = report.report_list.status_id").
		Joins("JOIN report.report_segment ON report.report_segment.report_id = report.report_list.id").
		Joins("JOIN map.irrigations_segment ON map.irrigations_segment.id = report.report_segment.segment_id").
		Joins("JOIN map.irrigations ON map.irrigations.id = map.irrigations_segment.irrigation_id").
		Where("report.report_list.user_id = ? AND map.irrigations.name LIKE %?%", user_id, search).
		Scan(&reports)
	} else {
		queryReports = models.Db.Table("report.report_list").
		Select("DISTINCT ON (report.report_list.id) report.report_list.id", "report.report_list.created_at", "report.report_list.done_at", "report.status.name as status", "map.irrigations.name as irrigation_name", "map.irrigations.type as canal").
		Joins("JOIN report.status ON report.status.id = report.report_list.status_id").
		Joins("JOIN report.report_segment ON report.report_segment.report_id = report.report_list.id").
		Joins("JOIN map.irrigations_segment ON map.irrigations_segment.id = report.report_segment.segment_id").
		Joins("JOIN map.irrigations ON map.irrigations.id = map.irrigations_segment.irrigation_id").
		Where("report.report_list.user_id = ?", user_id).Scan(&reports)
	}

	if queryReports.Error != nil {
		fmt.Printf("%v", queryReports.Error)
	}
	err := json.NewEncoder(w).Encode(reports)
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func CreateReport(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	r.ParseForm()
	year, month, day := time.Now().Date()
	strYear := fmt.Sprintf("%v", year)
	shortYear := strYear[2:4]
	rand.Seed(int64(day)+time.Now().UnixNano())
	min := 10000
	max := 99999
	ticket_no := fmt.Sprintf("%v%v%v", shortYear, int(month), (rand.Intn(max - min + 1) + min))
	user_id := fmt.Sprintf("%v", r.Context().Value("user_id"))
	var report = models.Report{
		UserID: user_id,
		StatusID: "485a2f73-294c-4511-ae87-59e70391a6db",
		TicketNo: ticket_no,
	}
	models.Db.Create(&report)
	report_id := fmt.Sprintf("%v", report.ID)
	var reportSegment = models.ReportSegment{
		ReportID: report_id,
		SegmentID: r.Form["segment_id1"][0],
		Level: r.Form["level1"][0],
		Note: r.Form["note1"][0],
	}
	models.Db.Create(&reportSegment)
	report_segment_id := fmt.Sprintf("%v", reportSegment.ID)
	uploadDumpID, err :=UploadImage(r.Form["image1"][0])
	var res Response

	// Create/Insert All Report Segment that is send using loop, to DB
	for i := 0; i < len(reportSegments); i++ {
		reportPhotoID, err := UploadImage(reportSegments[i].Image)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			res.Message = fmt.Sprintf("%s", err)
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				fmt.Printf("%s", err)
			}
			return
		}
		var reportSegmentforDB = models.ReportSegment{
			ReportID:      report_id,
			SegmentID:     reportSegments[i].Segment_id,
			ReportPhotoID: reportPhotoID,
			Level:         reportSegments[i].Level,
			Note:          reportSegments[i].Note,
		}
		models.Db.Create(&reportSegmentforDB)
	}

	w.WriteHeader(http.StatusCreated)
	res.Message = "Create report operation is successful"
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Printf("%s", err)
	}
}
