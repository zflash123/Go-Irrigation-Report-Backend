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

type Status	struct {
	ID					uuid.UUID
	Name				string
	CreatedAt 	time.Time			 `gorm:"autoCreateTime"`
	UpdatedAt 	time.Time			 `gorm:"autoUpdateTime"`
}

func (Status) TableName() string {
	return "report.status"
}

type Report	struct {
	ID					uuid.UUID				`gorm:"type:uuid;default:gen_random_uuid()"`
	UserID			string
	StatusID		string
	TicketNo		string
	CreatedAt 	time.Time			 `gorm:"autoCreateTime"`
	UpdatedAt 	time.Time			 `gorm:"autoUpdateTime"`
	User				User
	Status			Status
}

func (Report) TableName() string {
	return "report.report_list"
}

type ReportSegment struct {
	ID					uuid.UUID			 `gorm:"type:uuid;default:gen_random_uuid()"`
	ReportID		string
	SegmentID		string
	Level				string
	Note				string
	CreatedAt 	time.Time			 `gorm:"autoCreateTime"`
	UpdatedAt 	time.Time			 `gorm:"autoUpdateTime"`
}

func (ReportSegment) TableName() string {
	return "report.report_segment"
}

type UploadDump struct {
	ID					uuid.UUID
	Filename		string
	FileType		string
	Size				uint32
	Folder			string
	FileUrl			string
	CreatedAt 	time.Time			 `gorm:"autoCreateTime"`
	UpdatedAt 	time.Time			 `gorm:"autoUpdateTime"`
}

func (UploadDump) TableName() string {
	return "file.upload_dump"
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