package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/api/entity"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/api/repository"
	"golang.org/x/crypto/bcrypt"
)

type Services interface {
	AddAuthorRegistrationSendEmail(req RequestAuthor, regToken string) (entity.AuthorRegistration, error)
	AddReaderRegistrationSendEmail(req RequestReader, regToken string) (entity.ReaderRegistration, error)
	AddAdminRegistrationSendEmail(req RequestAdmin, regToken string) (entity.AdminRegistration, error)
	GetAuthorRegistration(email string) (*entity.AuthorRegistration, error)
	GetReaderRegistration(email string) (*entity.ReaderRegistration, error)
	GetAdminRegistration(email string) (*entity.AdminRegistration, error)
	CreateAuthor(req RequestAuthor) (*entity.Author, error)
	CreateReader(req RequestReader) (*entity.Reader, error)
	CreateAdmin(req RequestAdmin) (*entity.Admin, error)
	AuthUser(req RequestUserLogin) (*entity.User, error)
	AddNews(req RequestNews) (*entity.News, error)
	GetAuthor(username string) (*entity.Author, error)
	GetCategory(id uint) (*entity.Category, error)
	GetAuthorByID(id uint) (*entity.Author, error)
	GetAllNews() ([]entity.News, error)
	GetNews(newsID uint) (*entity.News, error)
	IsNewsExist(id uint) bool
	UpdateNews(req RequestNews) (*entity.News, error)
	DeleteNews(newsID uint) (*entity.News, error)
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

func (s *services) AddAdminRegistrationSendEmail(req RequestAdmin, regToken string) (entity.AdminRegistration, error) {
	registration := entity.AdminRegistration{}
	registration.Email = req.Email
	registration.Name = req.Name
	registration.Password = req.Password
	registration.Username = req.Username
	registration.ProfPic = req.ProfPic
	registration.RegistrationToken = regToken
	registration.TimeCreated = time.Now()

	registration, err := s.repository.AddAdminRegistration(registration)
	return registration, err
}

func (s *services) GetAuthorRegistration(email string) (*entity.AuthorRegistration, error) {
	registration, err := s.repository.GetAuthorRegistration(email)
	if err != nil {
		return registration, err
	}
	return registration, nil
}
func (s *services) GetReaderRegistration(email string) (*entity.ReaderRegistration, error) {
	registration, err := s.repository.GetReaderRegistration(email)
	if err != nil {
		return registration, err
	}
	return registration, nil
}
func (s *services) GetAdminRegistration(email string) (*entity.AdminRegistration, error) {
	registration, err := s.repository.GetAdminRegistration(email)
	if err != nil {
		return registration, err
	}
	return registration, nil
}

func (s *services) CreateAuthor(req RequestAuthor) (*entity.Author, error) {
	author := entity.Author{}
	author.Name = req.Name
	author.Email = req.Email
	author.Username = req.Username
	author.ProfPic = req.ProfPic
	author.KtpPic = req.KtpPic
	author.Experienced = req.Experienced

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	author.Password = string(hashedPassword)

	newAuthor, err := s.repository.AddAuthor(author)
	if err != nil {
		return newAuthor, err
	}
	return newAuthor, nil
}

func (s *services) CreateReader(req RequestReader) (*entity.Reader, error) {
	reader := entity.Reader{}
	reader.Name = req.Name
	reader.Email = req.Email
	reader.Username = req.Username
	reader.ProfPic = req.ProfPic

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	reader.Password = string(hashedPassword)

	newReader, err := s.repository.AddReader(reader)
	if err != nil {
		return newReader, err
	}
	return newReader, nil
}

func (s *services) CreateAdmin(req RequestAdmin) (*entity.Admin, error) {
	admin := entity.Admin{}
	admin.Name = req.Name
	admin.Email = req.Email
	admin.Username = req.Username
	admin.ProfPic = req.ProfPic

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	admin.Password = string(hashedPassword)

	newAdmin, err := s.repository.AddAdmin(admin)
	if err != nil {
		return newAdmin, err
	}
	return newAdmin, nil
}

func (s *services) AuthUser(req RequestUserLogin) (*entity.User, error) {
	username := req.Username
	password := req.Password
	fmt.Println("AUTHUSER CALLED, username: ", username, " password: ", password)

	//check Author Table
	author, err := s.repository.GetAuthor(username)
	// if err != nil {
	// 	return nil, errors.New("username is not registered")
	// }

	//check Reader Table
	reader, err := s.repository.GetReader(username)
	// if err != nil {
	// 	return nil, errors.New("username is not registered")
	// }

	//check Admin Table
	admin, err := s.repository.GetAdmin(username)

	// if err != nil {
	// 	return nil, errors.New("username is not registered - author-reader-admin tables checked")
	// }

	//convert Author or Reader or Admin into User
	var user = &entity.User{}
	if author != nil {
		user = s.repository.ConvertAuthortoUser(*author)
	} else if reader != nil {
		user = s.repository.ConvertReadertoUser(*reader)
	} else if admin != nil {
		user = s.repository.ConvertAdmintoUser(*admin)
	} else {
		// return nil, errors.New("user is not registered")
		return nil, err
	}

	test, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Printf("COMPARES: %s %s \n", user.Password, string(test))
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid email or password")
	}
	return user, nil
}

func (s *services) AddNews(req RequestNews) (*entity.News, error) {
	news := entity.News{}
	news.AuthorID = req.AuthorID
	news.CategoryID = req.CategoryID
	news.Title = req.Title
	news.Content = req.Content
	news.ImageUrl = req.ImageUrl

	newNews, err := s.repository.AddNews(news)
	if err != nil {
		return nil, err
	}
	return newNews, nil

}

func (s *services) GetAuthor(username string) (*entity.Author, error) {
	author, err := s.repository.GetAuthor(username)
	if err != nil {
		return author, err
	}
	return author, nil
}

func (s *services) GetCategory(id uint) (*entity.Category, error) {
	category, err := s.repository.GetCategory(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
func (s *services) GetAuthorByID(id uint) (*entity.Author, error) {
	author, err := s.repository.GetAuthorByID(id)
	if err != nil {
		return nil, err
	}
	return author, nil
}
func (s *services) GetAllNews() ([]entity.News, error) {
	news, err := s.repository.GetAllNews()
	if err != nil {
		return nil, err
	}
	return news, nil
}

func (s *services) GetNews(newsID uint) (*entity.News, error) {
	news, err := s.repository.GetNews(newsID)
	if err != nil {
		return nil, err
	}
	return news, nil
}

func (s *services) IsNewsExist(id uint) bool {
	news, _ := s.repository.GetNews(id)
	if news == nil {
		return false
	}
	return true
}

func (s *services) UpdateNews(req RequestNews) (*entity.News, error) {
	news := entity.News{}
	news.AuthorID = req.AuthorID
	news.CategoryID = req.CategoryID
	news.Title = req.Title
	news.Content = req.Content
	news.ImageUrl = req.ImageUrl

	newNews, err := s.repository.UpdateNews(news)
	if err != nil {
		return nil, err
	}
	return newNews, nil

}

func (s *services) DeleteNews(newsID uint) (*entity.News, error) {
	news, err := s.repository.DeleteNews(newsID)
	if err != nil {
		return nil, err
	}
	return news, nil
}
