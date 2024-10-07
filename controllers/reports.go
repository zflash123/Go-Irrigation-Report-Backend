package controllers

import(
	"net/http"
	"go-irrigation-report-backend/models"

	"github.com/gorilla/mux"
)

type Reports struct {
	Id									string
}

func GetReportById(w http.ResponseWriter, r *http.Request){
	
}