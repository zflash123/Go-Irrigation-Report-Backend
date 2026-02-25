package routes

import (
	"fmt"
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
	r.HandleFunc("/api/auth/register", controllers.Register).Methods("POST")
	r.HandleFunc("/api/auth/login", controllers.Login).Methods("POST")
	//routes for normal users
	user := r.PathPrefix("/api").Subrouter()
	user.Use(middleware.VerifyJwtToken)
	user.HandleFunc("/check-valid-cookie", controllers.CheckValidJwt).Methods("GET")
	user.HandleFunc("/close-segments", controllers.GetCloseSegments).Methods("GET")
	user.HandleFunc("/user/segments", controllers.GetSegmentsByUserId).Methods("GET")
	user.HandleFunc("/user/report/{id}", controllers.GetReportById).Methods("GET")
	user.HandleFunc("/user/reports", controllers.GetReportByUserId).Methods("GET")
	user.HandleFunc("/user/report", controllers.CreateReport).Methods("POST")
	user.HandleFunc("/user/profile", controllers.GetUserProfile).Methods("GET")
	user.HandleFunc("/user/profile", controllers.PutUserProfile).Methods("PUT")

	handler := config.CorsObject.Handler(r)
	fmt.Println("HTTP server run on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
