package models

import (
	"gorm.io/gorm"
	"time"
	"github.com/google/uuid"
)

type UserRole struct {
	ID					uuid.UUID			 `gorm:"type:uuid;default:gen_random_uuid()"`
	Name				string
}

func (UserRole) TableName() string {
	return "user.user_roles"
}

type User struct {
	ID					uuid.UUID			 `gorm:"type:uuid;default:gen_random_uuid()"`
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

func (User) TableName() string {
	return "user.users"
}

type Report	struct {
	ID					uuid.UUID				`gorm:"type:uuid;default:gen_random_uuid()"`
	UserID			string
	StatusID		string
	NoTicket		string
	Note				string
	CreatedAt 	time.Time			 `gorm:"autoCreateTime"`
	UpdatedAt 	time.Time			 `gorm:"autoUpdateTime"`
	User				User
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