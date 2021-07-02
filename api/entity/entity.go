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

type Reader struct {
	gorm.Model
	Name         string
	Email        string
	Password     string
	Username     string
	ProfPic      string
	NewsComments []NewsComment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	NewsReaders  []NewsReaders `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Admin struct {
	gorm.Model
	Name       string
	Email      string
	Password   string
	Username   string
	ProfPic    string
	Categories []Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Username string
	ProfPic  string
	Role     string
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
type AdminRegistration struct {
	gorm.Model
	Name              string `json:"name"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Username          string `json:"username"`
	ProfPic           string `json:"prof_pic"`
	RegistrationToken string
	TimeCreated       time.Time
}

type RegistrationReader struct {
	gorm.Model
	Name              string
	Email             string
	Password          string
	Username          string
	ProfPic           string
	RegistrationToken string
	TimeCreated       time.Time
}

type Category struct {
	gorm.Model
	AdminID int
	Name    string
	News    []News `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type News struct {
	gorm.Model
	AuthorID     int
	CategoryID   int
	Title        string
	Content      string
	ImageUrl     string
	NewsComments []NewsComment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	NewsReaders  []NewsReaders `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type NewsReaders struct {
	gorm.Model
	NewsID     int
	ReaderID   int
	TotalLike  int
	TotalShare int
	TotalView  int
}
type NewsComment struct {
	gorm.Model
	ReaderID int
	NewsID   int
	Comment  string
}

type Trending struct {
	ID  int
	Sum int
}

type Statistic struct {
	TotalAuthor int
	TotalReader int
	TotalNews   int
}
