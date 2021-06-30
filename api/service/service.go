package service

import (
	"time"

	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/api/entity"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/api/repository"
)

type Services interface {
	AddAuthorRegistrationSendEmail(req RequestAuthor, regToken string) (entity.AuthorRegistration, error)
	AddReaderRegistrationSendEmail(req RequestReader, regToken string) (entity.ReaderRegistration, error)
}

type services struct {
	repository repository.Repository
}

func NewServices(repository repository.Repository) *services {
	return &services{repository}
}

func (s *services) AddAuthorRegistrationSendEmail(req RequestAuthor, regToken string) (entity.AuthorRegistration, error) {
	registration := entity.AuthorRegistration{}
	registration.Email = req.Email
	registration.Name = req.Name
	registration.Password = req.Password
	registration.Username = req.Username
	registration.ProfPic = req.ProfPic
	registration.KtpPic = req.KtpPic
	registration.Experienced = req.Experienced
	registration.RegistrationToken = regToken
	registration.TimeCreated = time.Now()

	registration, err := s.repository.AddAuthorRegistration(registration)
	return registration, err
}

func (s *services) AddReaderRegistrationSendEmail(req RequestReader, regToken string) (entity.ReaderRegistration, error) {
	registration := entity.ReaderRegistration{}
	registration.Email = req.Email
	registration.Name = req.Name
	registration.Password = req.Password
	registration.Username = req.Username
	registration.ProfPic = req.ProfPic
	registration.RegistrationToken = regToken
	registration.TimeCreated = time.Now()

	registration, err := s.repository.AddReaderRegistration(registration)
	return registration, err
}
