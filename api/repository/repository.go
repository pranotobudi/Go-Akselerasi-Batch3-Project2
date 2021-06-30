package repository

import (
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/api/entity"
	"gorm.io/gorm"
)

type Repository interface {
	AddAuthorRegistration(authorReg entity.AuthorRegistration) (entity.AuthorRegistration, error)
	AddReaderRegistration(readerReg entity.ReaderRegistration) (entity.ReaderRegistration, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddAuthorRegistration(authorReg entity.AuthorRegistration) (entity.AuthorRegistration, error) {
	err := r.db.Create(&authorReg).Error
	if err != nil {
		return authorReg, err
	}
	return authorReg, nil
}
func (r *repository) AddReaderRegistration(readerReg entity.ReaderRegistration) (entity.ReaderRegistration, error) {
	err := r.db.Create(&readerReg).Error
	if err != nil {
		return readerReg, err
	}
	return readerReg, nil
}
