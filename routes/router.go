package routes

import (
	"go-irrigation-report-backend/controllers"
	"go-irrigation-report-backend/models"
	"log"
	"net/http"

	"go-irrigation-report-backend/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Routes() {
	models.Db_connection()
	r := mux.NewRouter()
	//basic
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	//routes for normal users
	user := r.PathPrefix("/api/user").Subrouter()
	user.Use(middleware.VerifyJwtToken)
	user.HandleFunc("/close-segments", controllers.GetCloseSegments).Methods("GET")
	user.HandleFunc("/report/{id}", controllers.GetReportById).Methods("GET")
	user.HandleFunc("/report", controllers.CreateReport).Methods("POST")

	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
