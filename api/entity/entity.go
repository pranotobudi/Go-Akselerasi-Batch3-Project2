package entity

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name        string
	Email       string
	Password    string
	Username    string
	ProfPic     string
	KtpPic      string
	Experienced bool
	News        []News `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type AuthorRegistration struct {
	gorm.Model
	Name              string `json:"name"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Username          string `json:"username"`
	ProfPic           string `json:"prof_pic"`
	KtpPic            string `json:"ktp_pic"`
	Experienced       bool   `json:"experienced"`
	RegistrationToken string
	TimeCreated       time.Time
}
type ReaderRegistration struct {
	gorm.Model
	Name              string `json:"name"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Username          string `json:"username"`
	ProfPic           string `json:"prof_pic"`
	RegistrationToken string
	TimeCreated       time.Time
}

type Reader struct {
	gorm.Model
	Name         string
	Email        string
	Password     string
	Username     string
	Prof_Pic     string
	NewsComments []NewsComment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type RegistrationReader struct {
	gorm.Model
	Name              string
	Email             string
	Password          string
	Username          string
	Prof_Pic          string
	RegistrationToken string
	TimeCreated       time.Time
}

type Admin struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Username string
	Prof_Pic string
}

type Category struct {
	gorm.Model
	Name string
	News []News `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type News struct {
	gorm.Model
	AuthorID     int
	CategoryID   int
	Title        string
	Content      string
	ImageUrl     string
	NewsComments []NewsComment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type NewsComment struct {
	gorm.Model
	ReaderID int
	NewsID   int
	Comment  string
}
