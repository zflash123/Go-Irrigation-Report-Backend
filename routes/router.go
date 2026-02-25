package routes

import (
	"go-irrigation-report-backend/controllers"
	"go-irrigation-report-backend/models"
	"go-irrigation-report-backend/config"
	"log"
	"net/http"

	"go-irrigation-report-backend/middleware"

	"github.com/gorilla/mux"
)

func Routes() {
	models.Db_connection()
	r := mux.NewRouter()
	//basic
	r.HandleFunc("/api/register", controllers.Register).Methods("POST")
	r.HandleFunc("/api/login", controllers.Login).Methods("POST")
	//routes for normal users
	user := r.PathPrefix("/api/user").Subrouter()
	user.Use(middleware.VerifyJwtToken)
	user.HandleFunc("/close-segments", controllers.GetCloseSegments).Methods("GET")
	user.HandleFunc("/segments", controllers.GetSegmentsByUserId).Methods("GET")
	user.HandleFunc("/report/{id}", controllers.GetReportById).Methods("GET")
	user.HandleFunc("/reports", controllers.GetReportByUserId).Methods("GET")
	user.HandleFunc("/report", controllers.CreateReport).Methods("POST")
	user.HandleFunc("/profile", controllers.GetUserProfile).Methods("GET")
	user.HandleFunc("/profile", controllers.PutUserProfile).Methods("PUT")

	handler := config.CorsObject.Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
