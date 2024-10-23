package controllers

import(
	"fmt"
	"net/http"
	"encoding/json"
	"go-irrigation-report-backend/models"
)

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	user_id := fmt.Sprintf("%v", r.Context().Value("user_id"))

	var user models.User
	query := models.Db.Where("id = ?", user_id).Take(&user)

	if query.Error != nil {
		var res Response
		res.Message = "There is an error when executing the query."
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			fmt.Printf("%v", err)
		}
	}
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Printf("%v", err)
	}
}