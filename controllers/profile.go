package controllers

import (
	"encoding/json"
	"fmt"
	"go-irrigation-report-backend/helperfunctions"
	"go-irrigation-report-backend/models"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type ReportPhoto struct {
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

	//Parse "application/json"
	type ProfileFormData struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Image     string `json:"image"`
	}
	var profileFormData ProfileFormData
	json.NewDecoder(r.Body).Decode(&profileFormData)
	//Check for required request parameters
	isParameterEmpty := helperfunctions.MakeRequiredParameterTypeString(profileFormData.Firstname, "Firstname", w)
	if isParameterEmpty == true {
		return
	}
	isParameterEmpty = helperfunctions.MakeRequiredParameterTypeString(profileFormData.Lastname, "Lastname", w)
	if isParameterEmpty == true {
		return
	}
	//Image Input Data Value Check
	var isImageFieldInFormFilled = false
	if profileFormData.Image != "" {
		isImageFieldInFormFilled = true
	}
	//Creating response Object
	var res Response
	// Executing Query for the given condition
	if isImageFieldInFormFilled == true {
		log.Println("lastname: ", profileFormData.Lastname)
		avatar, imagePath, err := UploadImageForProfile(profileFormData.Image)
		if err != nil {
			log.Println("Error di UploadImageForProfile")
		}
		query := models.Db.Model(&user).Updates(models.User{
			FirstName: profileFormData.Firstname,
			LastName:  profileFormData.Lastname,
			Avatar:    avatar,
		})
		if query.Error != nil {
			res.Message = "There is an error when executing the Update Profile Query."
			w.WriteHeader(http.StatusInternalServerError)
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				fmt.Printf("%v", err)
			}
			//Remove the Uploaded Image in this Working Directory outside of main thread for faster response of API
			return
		}
		var res Response
		res.Message = "Update profile successful"
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(res)
		go func(imagePath string) {
			os.Remove(imagePath)
		}(imagePath)
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
}
