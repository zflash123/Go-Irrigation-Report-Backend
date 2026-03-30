package controllers

import (
	"encoding/json"
	"fmt"
	"go-irrigation-report-backend/models"
	"net/http"

	"github.com/google/uuid"
)

type UploadDump struct {
	FileUrl string
}

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

func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	user_id := fmt.Sprintf("%v", r.Context().Value("user_id"))
	var user models.User
	user.ID, _ = uuid.Parse(user_id)
	r.ParseForm()
	if r.Form["image"][0] == "" {
		query := models.Db.Model(&user).Updates(models.User{
			FirstName: r.Form["firstname"][0],
			LastName:  r.Form["lastname"][0],
		})
		if query.Error != nil {
			var res Response
			res.Message = "There is an error when executing the query."
			w.WriteHeader(http.StatusInternalServerError)
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				fmt.Printf("%v", err)
			}
		}
	} else {
		query := models.Db.Model(&user).Updates(models.User{
			FirstName: profileFormData.Firstname,
			LastName:  profileFormData.Lastname,
		})
		if query.Error != nil {
			var res Response
			res.Message = "There is an error when executing the Update Profile Query."
			w.WriteHeader(http.StatusInternalServerError)
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				fmt.Printf("%v", err)
				return
			}
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
