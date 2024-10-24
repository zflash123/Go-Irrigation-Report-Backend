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
	Id									string
	CreatedAt						string
	DoneAt							string
	Level								string
	Note								string
	Status							string
	CenterPointJson			string
	IrrigationName			string
	Canal								string
	Image								string
}

func GetReportById(w http.ResponseWriter, r *http.Request){
	report_id := mux.Vars(r)["id"]
	var reports []Reports
	queryReports:= models.Db.Table("report.report_list").
	Select("report.report_list.id", "report.report_list.created_at", "report.report_list.done_at", "report.report_segment.level", "report.report_segment.note", "report.status.name as status", "map.irrigations_segment.center_point_json", "map.irrigations.name as irrigation_name", "map.irrigations.type as canal", "file.upload_dump.file_url as image").
	Joins("JOIN report.status ON report.status.id = report.report_list.status_id").
	Joins("JOIN report.report_segment ON report.report_segment.report_id = report.report_list.id").
	Joins("JOIN report.report_photo ON report.report_photo.report_segment_id = report.report_segment.id").
	Joins("JOIN file.upload_dump ON file.upload_dump.id = report.report_photo.upload_dump_id").
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

	var reports []Reports
	queryReports:= models.Db.Table("report.report_list").
	Select("DISTINCT ON (report.report_list.id) report.report_list.id", "report.report_list.created_at", "report.report_list.done_at", "report.status.name as status", "map.irrigations.name as irrigation_name", "map.irrigations.type as canal").
	Joins("JOIN report.status ON report.status.id = report.report_list.status_id").
	Joins("JOIN report.report_segment ON report.report_segment.report_id = report.report_list.id").
	Joins("JOIN map.irrigations_segment ON map.irrigations_segment.id = report.report_segment.segment_id").
	Joins("JOIN map.irrigations ON map.irrigations.id = map.irrigations_segment.irrigation_id").
	Where("report.report_list.user_id = ?", user_id).Scan(&reports)

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
	if err!=nil {
		res.Message = fmt.Sprintf("%s", err)
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			fmt.Printf("%s", err)
		}
		return
	}
	var reportPhoto = models.ReportPhoto{
		ReportSegmentID: report_segment_id,
		UploadDumpID: uploadDumpID,
	}
	models.Db.Create(&reportPhoto)
	if r.Form["segment_id2"][0] != "" {
		var reportSegment = models.ReportSegment{
			ReportID: report_id,
			SegmentID: r.Form["segment_id2"][0],
			Level: r.Form["level2"][0],
			Note: r.Form["note2"][0],
		}
		models.Db.Create(&reportSegment)
		report_segment_id := fmt.Sprintf("%v", reportSegment.ID)
		uploadDumpID, err :=UploadImage(r.Form["image2"][0])
		if err!=nil {
			res.Message = fmt.Sprintf("%s", err)
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				fmt.Printf("%s", err)
			}
			return
		}
		var reportPhoto = models.ReportPhoto{
			ReportSegmentID: report_segment_id,
			UploadDumpID: uploadDumpID,
		}
		models.Db.Create(&reportPhoto)
	}
	if r.Form["segment_id3"][0] != "" {
		var reportSegment = models.ReportSegment{
			ReportID: report_id,
			SegmentID: r.Form["segment_id3"][0],
			Level: r.Form["level3"][0],
			Note: r.Form["note3"][0],
		}
		models.Db.Create(&reportSegment)
		report_segment_id := fmt.Sprintf("%v", reportSegment.ID)
		uploadDumpID, err :=UploadImage(r.Form["image3"][0])
		if err!=nil {
			res.Message = fmt.Sprintf("%s", err)
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				fmt.Printf("%s", err)
			}
			return
		}
		var reportPhoto = models.ReportPhoto{
			ReportSegmentID: report_segment_id,
			UploadDumpID: uploadDumpID,
		}
		models.Db.Create(&reportPhoto)
	}
	w.WriteHeader(http.StatusCreated)
	res.Message = "Create report operation is successful"
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Printf("%s", err)
	}
}