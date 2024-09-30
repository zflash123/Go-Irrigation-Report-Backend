package models

import (
	"gorm.io/gorm"
	"time"
)

type UserRole struct {
	ID					string				 `gorm:"primarykey"`
	Name				string
}

type User struct {
	ID					uint					 `gorm:"primarykey"`
	UserRoleID 	string
	Name				string
	Email				string
	Password		string
	Username		string
	FirstName		string
	LastName		string
	Avatar			string
	CreatedAt 	time.Time			 `gorm:"autoCreateTime"`
	UpdatedAt 	time.Time			 `gorm:"autoUpdateTime"`
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
	UserRole		UserRole
}

type Book struct {
	gorm.Model
	Name		string
	Year		int
	Author		string
	Summary		string
	Publisher	string
	PageCount	int
	ReadPage	int
	Finished	bool
	Reading		bool
}

var Db *gorm.DB
var Err error