package handler

import (
	"fmt"
	"net/http"
	"strconv"

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
		"http://localhost:8080/api/register/confirmation?email=" + email + "&token=" + regToken + "&role=author")

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
		"http://localhost:8080/api/register/confirmation?email=" + email + "&token=" + regToken + "&role=reader")

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

func (h *handler) AdminRegistrationSendEmail(c echo.Context) error {
	adminReq := new(service.RequestAdmin)
	if err := c.Bind(adminReq); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	//Send Confirmation Email
	regToken := randstr.Hex(16) // generate 128-bit hex string
	newAdminRegistration, err := h.service.AddAdminRegistrationSendEmail(*adminReq, regToken)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "ragistration send email failed", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	email := adminReq.Email
	msg := []byte("To: " + email + "\r\n" +
		"Subject: Registration Confirmation Email from News App!\r\n" +
		"\r\n" +
		"This is the email body.\r\n" +
		"http://localhost:8080/api/register/confirmation?email=" + email + "&token=" + regToken + "&role=admin")

	toEmail := []string{email}
	helper.SendEmail(toEmail, msg)

	auth_token, err := h.authService.CreateAccessToken("admin")
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}
	adminRegistrationData := service.AdminRegistrationResponseFormatter(newAdminRegistration, auth_token)

	response := helper.ResponseFormatter(http.StatusOK, "success", "sending email for admin registration successfull", adminRegistrationData)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) RegisterConfirmation(c echo.Context) error {
	role := c.QueryParam("role")
	fmt.Println("INSIDE Handler:RegisterConfirmation role:", role)

	var err error
	if role == "author" {
		err = h.AuthorRegisterConfirmation(c)
	} else if role == "reader" {
		err = h.ReaderRegisterConfirmation(c)
	} else if role == "admin" {
		err = h.AdminRegisterConfirmation(c)
	}
	return err
}

func (h *handler) AuthorRegisterConfirmation(c echo.Context) error {
	email := c.QueryParam("email")
	token := c.QueryParam("token")
	fmt.Println("INSIDE userHandler:AuthorRegisterConfirmation email:", email, " token: ", token)
	registration, err := h.service.GetAuthorRegistration(email)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if registration.RegistrationToken != token {
		errorFormatter := helper.ErrorFormatter(fmt.Errorf("token is not valid"))
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)

	}

	// Add to Author Table
	author := new(service.RequestAuthor)
	author.Name = registration.Name
	author.Email = registration.Email
	author.Password = registration.Password
	author.Username = registration.Username
	author.ProfPic = registration.ProfPic
	author.KtpPic = registration.KtpPic
	author.Experienced = registration.Experienced

	newAuthor, _ := h.service.CreateAuthor(*author)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	// Send response

	// role, _ := h.service.GetRole(newUser.ID)
	auth_token, err := h.authService.CreateAccessToken("author")
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}
	authorData := service.AuthorResponseFormatter(*newAuthor, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "author registration confirmation successfull, author created", authorData)

	return c.JSON(http.StatusOK, response)
}
func (h *handler) ReaderRegisterConfirmation(c echo.Context) error {
	email := c.QueryParam("email")
	token := c.QueryParam("token")
	fmt.Println("INSIDE userHandler:ReaderRegisterConfirmation email:", email, " token: ", token)
	registration, err := h.service.GetReaderRegistration(email)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if registration.RegistrationToken != token {
		errorFormatter := helper.ErrorFormatter(fmt.Errorf("token is not valid"))
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)

	}

	// Add to Reader Table
	reader := new(service.RequestReader)
	reader.Name = registration.Name
	reader.Email = registration.Email
	reader.Password = registration.Password
	reader.Username = registration.Username
	reader.ProfPic = registration.ProfPic

	newReader, _ := h.service.CreateReader(*reader)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	// Send response

	// role, _ := h.service.GetRole(newUser.ID)
	auth_token, err := h.authService.CreateAccessToken("reader")
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}
	readerData := service.ReaderResponseFormatter(*newReader, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "reader registration confirmation successfull, reader created", readerData)

	return c.JSON(http.StatusOK, response)

}
func (h *handler) AdminRegisterConfirmation(c echo.Context) error {
	email := c.QueryParam("email")
	token := c.QueryParam("token")
	fmt.Println("INSIDE userHandler:ReaderRegisterConfirmation email:", email, " token: ", token)
	registration, err := h.service.GetAdminRegistration(email)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if registration.RegistrationToken != token {
		errorFormatter := helper.ErrorFormatter(fmt.Errorf("token is not valid"))
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)

	}

	// Add to Admin Table
	admin := new(service.RequestAdmin)
	admin.Name = registration.Name
	admin.Email = registration.Email
	admin.Password = registration.Password
	admin.Username = registration.Username
	admin.ProfPic = registration.ProfPic

	newAdmin, _ := h.service.CreateAdmin(*admin)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	// Send response

	// role, _ := h.service.GetRole(newUser.ID)
	auth_token, err := h.authService.CreateAccessToken("admin")
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}
	adminData := service.AdminResponseFormatter(*newAdmin, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "admin registration confirmation successfull, admin created", adminData)

	return c.JSON(http.StatusOK, response)

}

func (h *handler) UserLogin(c echo.Context) error {
	userLogin := new(service.RequestUserLogin)
	if err := c.Bind(userLogin); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	authUser, err := h.service.AuthUser(*userLogin)
	if err != nil {
		fmt.Println("We're IN HERE: USERLOGIN INSIDE")
		response := helper.ResponseFormatter(http.StatusUnauthorized, "error", err.Error(), nil)
		return c.JSON(http.StatusUnauthorized, response)
	}
	// role, _ := h.service.GetRole(authUser.ID)

	auth_token, err := h.authService.CreateAccessToken(authUser.Role)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}

	userData := service.UserResponseFormatter(*authUser, auth_token)

	response := helper.ResponseFormatter(http.StatusOK, "success", "user authenticated", userData)
	return c.JSON(http.StatusOK, response)
}

func (h *handler) AddNews(c echo.Context) error {
	news := new(service.RequestNews)
	if err := c.Bind(news); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	newNews, err := h.service.AddNews(*news)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}
	author, _ := h.service.GetAuthorByID(uint(newNews.AuthorID))
	category, _ := h.service.GetCategory(uint(newNews.CategoryID))

	newsData := service.NewsResponseFormatter(*newNews, *author, *category)
	response := helper.ResponseFormatter(http.StatusOK, "success", "news successfully added", newsData)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetAllNews(c echo.Context) error {
	news, err := h.service.GetAllNews()
	fmt.Printf("\n Handler GetAllNews: %+v \n", news)

	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}
	var finalNewsData []service.ResponseNews
	for _, singleNews := range news {
		author, _ := h.service.GetAuthorByID(uint(singleNews.AuthorID))
		category, _ := h.service.GetCategory(uint(singleNews.CategoryID))

		newsData := service.NewsResponseFormatter(singleNews, *author, *category)
		finalNewsData = append(finalNewsData, newsData)
	}

	response := helper.ResponseFormatter(http.StatusOK, "success", "get all news succeeded", finalNewsData)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetNews(c echo.Context) error {
	newsID, _ := strconv.Atoi(c.Param("id"))

	news, err := h.service.GetNews(uint(newsID))
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}
	author, _ := h.service.GetAuthorByID(uint(news.AuthorID))
	category, _ := h.service.GetCategory(uint(news.CategoryID))

	newsData := service.NewsResponseFormatter(*news, *author, *category)
	response := helper.ResponseFormatter(http.StatusOK, "success", "get news successfull", newsData)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) UpdateNews(c echo.Context) error {
	newsID, _ := strconv.Atoi(c.Param("id"))

	news := new(service.RequestNews)
	if err := c.Bind(news); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	isExist := h.service.IsNewsExist(uint(newsID))
	if isExist == false {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "record not found", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	newNews, err := h.service.UpdateNews(*news)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}
	author, _ := h.service.GetAuthorByID(uint(newNews.AuthorID))
	category, _ := h.service.GetCategory(uint(newNews.CategoryID))

	newsData := service.NewsResponseFormatter(*newNews, *author, *category)
	response := helper.ResponseFormatter(http.StatusOK, "success", "news successfully updated", newsData)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) DeleteNews(c echo.Context) error {
	newsID, _ := strconv.Atoi(c.Param("id"))

	news, err := h.service.DeleteNews(uint(newsID))
	fmt.Printf("================HANDLER NEWS: %+v \n\n", news)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}
	author, _ := h.service.GetAuthorByID(uint(news.AuthorID))
	category, _ := h.service.GetCategory(uint(news.CategoryID))

	newsData := service.NewsResponseFormatter(*news, *author, *category)
	response := helper.ResponseFormatter(http.StatusOK, "success", "delete news successfull", newsData)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetNewsByCategoryID(c echo.Context) error {
	categoryID, _ := strconv.Atoi(c.Param("id"))
	news, err := h.service.GetNewsByCategoryID(uint(categoryID))
	fmt.Printf("\n Handler GetNewsByCategoryID: %+v \n", news)

	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}
	var finalNewsData []service.ResponseNews
	for _, singleNews := range news {
		author, _ := h.service.GetAuthorByID(uint(singleNews.AuthorID))
		category, _ := h.service.GetCategory(uint(singleNews.CategoryID))

		newsData := service.NewsResponseFormatter(singleNews, *author, *category)
		finalNewsData = append(finalNewsData, newsData)
	}

	response := helper.ResponseFormatter(http.StatusOK, "success", "get news by category succeeded", finalNewsData)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetAllTrendingNews(c echo.Context) error {
	news, err := h.service.GetAllTrendingNews()
	fmt.Printf("\n Handler GetAllTrendingNews: %+v \n", news)

	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}
	var finalNewsData []service.ResponseNews
	for _, singleNews := range news {
		author, _ := h.service.GetAuthorByID(uint(singleNews.AuthorID))
		category, _ := h.service.GetCategory(uint(singleNews.CategoryID))

		newsData := service.NewsResponseFormatter(singleNews, *author, *category)
		finalNewsData = append(finalNewsData, newsData)
	}

	response := helper.ResponseFormatter(http.StatusOK, "success", "get all news succeeded", finalNewsData)

	return c.JSON(http.StatusOK, response)
}
