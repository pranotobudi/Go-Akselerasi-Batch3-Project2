package service

import "github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/api/entity"

type ResponseAuthorRegistration struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Username    string `json:"username"`
	ProfPic     string `json:"prof_pic"`
	KtpPic      string `json:"ktp_pic"`
	Experienced bool   `json:"experienced"`
	AuthToken   string `json:"auth_token"`
}
type ResponseReaderRegistration struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	ProfPic   string `json:"prof_pic"`
	AuthToken string `json:"auth_token"`
}

func AuthorRegistrationResponseFormatter(authorReg entity.AuthorRegistration, auth_token string) ResponseAuthorRegistration {
	formatter := ResponseAuthorRegistration{
		ID:          authorReg.ID,
		Name:        authorReg.Name,
		Email:       authorReg.Email,
		Password:    authorReg.Password,
		Username:    authorReg.Username,
		ProfPic:     authorReg.ProfPic,
		KtpPic:      authorReg.KtpPic,
		Experienced: authorReg.Experienced,
		AuthToken:   auth_token,
	}
	return formatter
}

func ReaderRegistrationResponseFormatter(readerReg entity.ReaderRegistration, auth_token string) ResponseReaderRegistration {
	formatter := ResponseReaderRegistration{
		ID:        readerReg.ID,
		Name:      readerReg.Name,
		Email:     readerReg.Email,
		Password:  readerReg.Password,
		Username:  readerReg.Username,
		ProfPic:   readerReg.ProfPic,
		AuthToken: auth_token,
	}
	return formatter
}
