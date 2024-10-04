package models

import (
	"gorm.io/gorm"
	"time"
	"github.com/google/uuid"
)

type UserRole struct {
	ID					uuid.UUID			 `gorm:"primarykey"`
	Name				string
}

type User struct {
	ID					uint					 `gorm:"primarykey"`
	UserRoleID 	string
	Email				string
	Password		string
	Username		string
	FirstName		string
	LastName		string
	Avatar			string
	CreatedAt 	time.Time			 `gorm:"autoCreateTime"`
	UpdatedAt 	time.Time			 `gorm:"autoUpdateTime"`
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