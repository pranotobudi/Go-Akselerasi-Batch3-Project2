package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/api/service"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/auth"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/helper"
	"github.com/thanhpk/randstr"
)

type handler struct {
	service     service.Services
	authService auth.AuthService
}

func NewHandler(service service.Services, authService auth.AuthService) *handler {
	return &handler{service, authService}
}

func (h *handler) AuthorRegistrationSendEmail(c echo.Context) error {
	authorReq := new(service.RequestAuthor)
	if err := c.Bind(authorReq); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	//Send Confirmation Email
	regToken := randstr.Hex(16) // generate 128-bit hex string
	newAuthorRegistration, err := h.service.AddAuthorRegistrationSendEmail(*authorReq, regToken)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "ragistration send email failed", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	email := authorReq.Email
	msg := []byte("To: " + email + "\r\n" +
		"Subject: Registration Confirmation Email from News App!\r\n" +
		"\r\n" +
		"This is the email body.\r\n" +
		"http://localhost:8080/api/v1/movie_reviews/register/confirmation?email=" + email + "&token=" + regToken + "&role=author")

	toEmail := []string{email}
	helper.SendEmail(toEmail, msg)

	auth_token, err := h.authService.CreateAccessToken("author")
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}
	authorRegistrationData := service.AuthorRegistrationResponseFormatter(newAuthorRegistration, auth_token)

	response := helper.ResponseFormatter(http.StatusOK, "success", "sending email for author registration successfull", authorRegistrationData)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) ReaderRegistrationSendEmail(c echo.Context) error {
	readerReq := new(service.RequestReader)
	if err := c.Bind(readerReq); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	//Send Confirmation Email
	regToken := randstr.Hex(16) // generate 128-bit hex string
	newReaderRegistration, err := h.service.AddReaderRegistrationSendEmail(*readerReq, regToken)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "ragistration send email failed", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	email := readerReq.Email
	msg := []byte("To: " + email + "\r\n" +
		"Subject: Registration Confirmation Email from News App!\r\n" +
		"\r\n" +
		"This is the email body.\r\n" +
		"http://localhost:8080/api/v1/movie_reviews/register/confirmation?email=" + email + "&token=" + regToken + "&role=author")

	toEmail := []string{email}
	helper.SendEmail(toEmail, msg)

	auth_token, err := h.authService.CreateAccessToken("reader")
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}
	readerRegistrationData := service.ReaderRegistrationResponseFormatter(newReaderRegistration, auth_token)

	response := helper.ResponseFormatter(http.StatusOK, "success", "sending email for reader registration successfull", readerRegistrationData)

	return c.JSON(http.StatusOK, response)
}
