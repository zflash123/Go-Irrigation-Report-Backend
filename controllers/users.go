package controllers

import (
	"encoding/json"
	"fmt"
	"go-irrigation-report-api/models"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

type Response struct{
	Message				string
}
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func Register(w http.ResponseWriter, r *http.Request) {
	var res Response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	hashedPwd, err := hashPassword(r.Form["password"][0])
	if(err!=nil){
		fmt.Println("Error = ", err)
	}
	var user = []models.User{
		{
			UserRoleID: r.Form["urole_id"][0],
			Email: r.Form["email"][0],
			Password: hashedPwd,
			Username: r.Form["username"][0],
			FirstName: r.Form["firstname"][0],
			LastName: r.Form["lastname"][0],
			Avatar: r.Form["avatar"][0],
		},
	}
	models.Db.Create(&user[0])
	fmt.Fprint(w, "Your account successfully registered")
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var users []models.User
	userData := models.Db.Where("email = ?", r.Form["email"][0]).First(&users)
	emailCheckErr := userData.Error
	isPwdCorrect := users[0].Password==r.Form["password"][0]
	type Response struct {
		Message string `json:"message"`
		Auth string `json:"auth_token"`
	}
	if emailCheckErr == nil && isPwdCorrect {
		var res Response
		res.Message = "Your account successfully logged in"
		strJwt := CreateJwt(users[0].Email, users[0].Name)
		res.Auth = strJwt

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		fmt.Fprint(w, "Email or password that you're inputted is wrong")
	}
}
