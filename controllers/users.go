package controllers

import (
	"encoding/json"
	"fmt"
	"go-irrigation-report-backend/models"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Message string
}

type CredentialsData struct {
	Email string
	Password string
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func isPasswordMatched(hashedPwd string, inputtedPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(inputtedPwd))
	if err != nil {
		return false
	}
	return true
}

func Register(w http.ResponseWriter, r *http.Request) {
	var res Response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	r.ParseForm()
	hashedPwd, err := hashPassword(r.Form["password"][0])
	if err != nil {
		fmt.Println("Error = ", err)
	}
	var user = []models.User{
		{
			UserRoleID: r.Form["urole_id"][0],
			Email:      r.Form["email"][0],
			Password:   hashedPwd,
			Username:   r.Form["username"][0],
			FirstName:  r.Form["firstname"][0],
			LastName:   r.Form["lastname"][0],
			Avatar:     r.Form["avatar"][0],
		},
	}
	createUser := models.Db.Create(&user[0])
	if createUser.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = "There is an error when registering your account"
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		res.Message = "Your account successfully registered"
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var email = ""
	var password = ""
	var credentialsData CredentialsData
	//If the req body use JSON data, then decode the JSON to an object
	if(r.Header.Get("Content-Type") == "application/json"){
		json.NewDecoder(r.Body).Decode(&credentialsData)
		email = credentialsData.Email
		password = credentialsData.Password
	//Else If the req body use x-www-form-urlencoded data, so it will parsed as form data
	} else if(r.Header.Get("Content-Type") == "application/x-www-form-urlencoded") {
		r.ParseForm()
		email = r.Form["email"][0]
		password = r.Form["password"][0]
	} else {
		type Response struct {
			Message string `json:"message"`
		}
		var res Response
		
		res.Message = "Content-Type not supported"
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Println(err)
		}
		return
	}


	var users []models.User
	userData := models.Db.Where("email = ?", email).First(&users)
	emailCheckErr := userData.Error
	isPwdCorrect := isPasswordMatched(users[0].Password, password)
	type Data struct {
		Message string `json:"message"`
		Auth    string `json:"jwtToken"`
	}
	type Response struct {
		Data		Data		`json:"data"`
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var res Response
	if emailCheckErr == nil && isPwdCorrect {
		res.Data.Message = "Your account successfully logged in"
		strUserID := fmt.Sprintf("%v", users[0].ID)
		strJwt := CreateJwt(strUserID, users[0].Email, users[0].FirstName)
		res.Data.Auth = strJwt

		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		res.Data.Message = "The email or password that you inputted is wrong"
		res.Data.Auth = "Not generated"

		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Println(err)
		}
	}
}

func CheckValidJwt(w http.ResponseWriter, r *http.Request) {
	return
}
