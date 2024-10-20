package controllers

import(
	"fmt"
	"time"
	"math/rand"
	"go-irrigation-report-backend/models"
)

func makeReport(user_id string, segment_id string, level string, note string, image string) error{
	year, month, day := time.Now().Date()
	strYear := fmt.Sprintf("%v", year)
	shortYear := strYear[2:4]
	rand.Seed(int64(day)+time.Now().UnixNano())
	min := 10000
	max := 99999
	ticket_no := fmt.Sprintf("%v%v%v", shortYear, int(month), (rand.Intn(max - min + 1) + min))
	var report = models.Report{
		UserID: user_id,
		StatusID: "485a2f73-294c-4511-ae87-59e70391a6db",
		TicketNo: ticket_no,
	}
	models.Db.Create(&report)
	report_id := fmt.Sprintf("%v", report.ID)
	var reportSegment = models.ReportSegment{
		ReportID: report_id,
		SegmentID: segment_id,
		Level: level,
		Note: note,
	}
	models.Db.Create(&reportSegment)
	report_segment_id := fmt.Sprintf("%v", reportSegment.ID)
	uploadDumpID, err :=UploadImage(image)
	if err!=nil {
		return err
	}
	var reportPhoto = models.ReportPhoto{
		ReportSegmentID: report_segment_id,
		UploadDumpID: uploadDumpID,
	}
	models.Db.Create(&reportPhoto)
	return nil
}