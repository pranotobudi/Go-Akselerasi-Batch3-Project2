package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/api/service"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/auth"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/helper"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/middleware"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/task"
	"github.com/thanhpk/randstr"
)

type handler struct {
	service     service.Services
	authService auth.AuthService
	taskService task.BackgroundTask
}

func NewHandler(service service.Services, authService auth.AuthService, taskService task.BackgroundTask) *handler {
	return &handler{service, authService, taskService}
}

func (h *handler) AuthorRegistrationSendEmail(c echo.Context) error {
	// Input Binding
	authorReq := new(service.RequestAuthor)
	if err := c.Bind(authorReq); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	// Add to database
	regToken := randstr.Hex(16) // generate 128-bit hex string
	newAuthorRegistration, err := h.service.AddAuthorRegistrationSendEmail(*authorReq, regToken)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "ragistration send email failed", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	//Send Confirmation Email

	// email := authorReq.Email
	// msg := []byte("To: " + email + "\r\n" +
	// 	"Subject: Registration Confirmation Email from News App!\r\n" +
	// 	"\r\n" +
	// 	"This is the email body.\r\n" +
	// 	"http://localhost:8080/api/register/confirmation?email=" + email + "&token=" + regToken + "&role=author")

	// toEmail := []string{email}
	// helper.SendEmail(toEmail, msg)

	//Schedule Email Sending
	email := authorReq.Email
	toEmail := []string{email}
	emailStruct := task.Email{toEmail, "author"}
	h.taskService.AddEmailQueue(emailStruct)

	// Create JWT token
	auth_token, err := h.authService.CreateAccessToken("author", newAuthorRegistration.ID)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}
	authorRegistrationData := service.AuthorRegistrationResponseFormatter(newAuthorRegistration, auth_token)

	response := helper.ResponseFormatter(http.StatusOK, "success", "sending email for author registration successfull", authorRegistrationData)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) ReaderRegistrationSendEmail(c echo.Context) error {
	// Input Binding
	readerReq := new(service.RequestReader)
	if err := c.Bind(readerReq); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	// Add to database
	regToken := randstr.Hex(16) // generate 128-bit hex string
	newReaderRegistration, err := h.service.AddReaderRegistrationSendEmail(*readerReq, regToken)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "ragistration send email failed", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	//Send Confirmation Email
	// email := readerReq.Email
	// msg := []byte("To: " + email + "\r\n" +
	// 	"Subject: Registration Confirmation Email from News App!\r\n" +
	// 	"\r\n" +
	// 	"This is the email body.\r\n" +
	// 	"http://localhost:8080/api/register/confirmation?email=" + email + "&token=" + regToken + "&role=reader")

	// toEmail := []string{email}
	// helper.SendEmail(toEmail, msg)

	//Schedule Email Sending
	email := readerReq.Email
	toEmail := []string{email}
	emailStruct := task.Email{toEmail, "reader"}
	h.taskService.AddEmailQueue(emailStruct)

	// create JWT token
	auth_token, err := h.authService.CreateAccessToken("reader", newReaderRegistration.ID)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}
	readerRegistrationData := service.ReaderRegistrationResponseFormatter(newReaderRegistration, auth_token)

	response := helper.ResponseFormatter(http.StatusOK, "success", "sending email for reader registration successfull", readerRegistrationData)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) AdminRegistrationSendEmail(c echo.Context) error {
	// Input Binding
	adminReq := new(service.RequestAdmin)
	if err := c.Bind(adminReq); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	// Add to database
	regToken := randstr.Hex(16) // generate 128-bit hex string
	newAdminRegistration, err := h.service.AddAdminRegistrationSendEmail(*adminReq, regToken)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "ragistration send email failed", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	//Send Confirmation Email
	// email := adminReq.Email
	// msg := []byte("To: " + email + "\r\n" +
	// 	"Subject: Registration Confirmation Email from News App!\r\n" +
	// 	"\r\n" +
	// 	"This is the email body.\r\n" +
	// 	"http://localhost:8080/api/register/confirmation?email=" + email + "&token=" + regToken + "&role=admin")

	// toEmail := []string{email}
	// helper.SendEmail(toEmail, msg)

	//Schedule Email Sending
	email := adminReq.Email
	toEmail := []string{email}
	emailStruct := task.Email{toEmail, "admin"}
	h.taskService.AddEmailQueue(emailStruct)

	//Create JWT Token
	auth_token, err := h.authService.CreateAccessToken("admin", newAdminRegistration.ID)
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
	auth_token, err := h.authService.CreateAccessToken("author", newAuthor.ID)
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
	auth_token, err := h.authService.CreateAccessToken("reader", newReader.ID)
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
	auth_token, err := h.authService.CreateAccessToken("admin", newAdmin.ID)
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
	fmt.Println("We're IN HERE: USERLOGIN INSIDE: authUser: ", authUser)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusUnauthorized, "error", err.Error(), nil)
		return c.JSON(http.StatusUnauthorized, response)
	}
	// role, _ := h.service.GetRole(authUser.ID)

	auth_token, err := h.authService.CreateAccessToken(authUser.Role, authUser.ID)
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

func (h *handler) GetAllNewsByCategory(c echo.Context) error {
	news, err := h.service.GetAllNewsByCategory()
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

	response := helper.ResponseFormatter(http.StatusOK, "success", "get all news sort by category succeeded", finalNewsData)

	return c.JSON(http.StatusOK, response)
}
func (h *handler) GetAllNewsByTrending(c echo.Context) error {
	news, err := h.service.GetAllNewsByTrending()
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

	response := helper.ResponseFormatter(http.StatusOK, "success", "get all news sort by trending succeeded", finalNewsData)

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

	// Add Total View
	readerID := middleware.GetJwtID(c)
	err = h.service.AddNewsView(uint(newsID), readerID)
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
	if !isExist {
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
func (h *handler) GetAllHighlightNews(c echo.Context) error {
	authorID := middleware.GetJwtID(c)
	fmt.Printf("\n =======Handler GetAllHighLightNews authorID: %+v \n", authorID)
	news, err := h.service.GetAllHighlightNews(authorID)
	fmt.Printf("\n ========Handler GetAllHighLightNews: %+v \n", news)

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

	response := helper.ResponseFormatter(http.StatusOK, "success", "get all highlight news succeeded", finalNewsData)

	return c.JSON(http.StatusOK, response)
}

// CRUD AUTHOR

func (h *handler) AddAuthor(c echo.Context) error {
	author := new(service.RequestAuthor)
	if err := c.Bind(author); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	newAuthor, err := h.service.AddAuthor(*author)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	auth_token, _ := h.authService.CreateAccessToken("author", newAuthor.ID)
	authorData := service.AuthorResponseFormatter(*newAuthor, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "author successfully added", authorData)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetAuthor(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	newAuthor, err := h.service.GetAuthorByID(uint(id))
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	auth_token, _ := h.authService.CreateAccessToken("author", newAuthor.ID)
	authorData := service.AuthorResponseFormatter(*newAuthor, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "author successfully added", authorData)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) UpdateAuthor(c echo.Context) error {
	author := new(service.RequestAuthor)
	if err := c.Bind(author); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	newAuthor, err := h.service.UpdateAuthor(*author)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	auth_token, _ := h.authService.CreateAccessToken("author", newAuthor.ID)
	authorData := service.AuthorResponseFormatter(*newAuthor, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "author successfully updated", authorData)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) DeleteAuthor(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("id"))

	author, err := h.service.DeleteAuthor(uint(ID))
	fmt.Printf("================HANDLER NEWS: %+v \n\n", author)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	auth_token, _ := h.authService.CreateAccessToken("author", author.ID)
	newsData := service.AuthorResponseFormatter(*author, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "delete news successfull", newsData)

	return c.JSON(http.StatusOK, response)
}

// CRUD READER

func (h *handler) AddReader(c echo.Context) error {
	reader := new(service.RequestReader)
	if err := c.Bind(reader); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	newreader, err := h.service.AddReader(*reader)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	auth_token, err := h.authService.CreateAccessToken("reader", newreader.ID)
	data := service.ReaderResponseFormatter(*newreader, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "reader successfully added", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetReader(c echo.Context) error {
	newsID, _ := strconv.Atoi(c.Param("id"))
	reader, err := h.service.GetReader(uint(newsID))
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	auth_token, _ := h.authService.CreateAccessToken("author", reader.ID)
	data := service.ReaderResponseFormatter(*reader, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "get reader successful", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) UpdateReader(c echo.Context) error {
	reader := new(service.RequestReader)
	if err := c.Bind(reader); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	newReader, err := h.service.UpdateReader(*reader)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	auth_token, _ := h.authService.CreateAccessToken("reader", newReader.ID)
	data := service.ReaderResponseFormatter(*newReader, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "reader successfully updated", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) DeleteReader(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	reader, err := h.service.DeleteReader(uint(id))
	fmt.Printf("================HANDLER NEWS: %+v \n\n", reader)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	auth_token, _ := h.authService.CreateAccessToken("reader", reader.ID)
	data := service.ReaderResponseFormatter(*reader, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "delete reader successfull", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetStatistic(c echo.Context) error {

	statistic, err := h.service.GetStatistic()
	fmt.Printf("================HANDLER NEWS: %+v \n\n", statistic)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	data := service.StatisticResponseFormatter(*statistic)
	response := helper.ResponseFormatter(http.StatusOK, "success", "delete reader successfull", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetAuthorNewsStatistic(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	statistic, err := h.service.GetNewsStatistic(uint(id))
	fmt.Printf("================HANDLER NEWS: %+v \n\n", statistic)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	data := service.NewsStatisticResponseFormatter(*statistic)
	response := helper.ResponseFormatter(http.StatusOK, "success", "delete reader successfull", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) AddComment(c echo.Context) error {
	comment := new(service.RequestComment)
	if err := c.Bind(comment); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	newComment, err := h.service.AddComment(*comment)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	data := service.CommentResponseFormatter(*newComment)
	response := helper.ResponseFormatter(http.StatusOK, "success", "comment successfully added", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) AddNewsShare(c echo.Context) error {
	newsID, _ := strconv.Atoi(c.Param("id"))
	readerID := middleware.GetJwtID(c)
	err := h.service.AddNewsShare(uint(newsID), readerID)

	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	data := ""
	response := helper.ResponseFormatter(http.StatusOK, "success", "add share successfull", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) UpdateAuthorProfile(c echo.Context) error {
	authorID, _ := strconv.Atoi(c.Param("id"))
	author, _ := h.service.GetAuthorByID(uint(authorID))

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	username := c.FormValue("username")
	experienced, _ := strconv.ParseBool(c.FormValue("experienced"))
	profPicHeader, _ := c.FormFile("prof_pic")
	ktpPicHeader, _ := c.FormFile("ktp_pic")
	// if err != nil {
	// 	return err
	// }
	profPicfile, _ := profPicHeader.Open()
	ktpPicfile, _ := ktpPicHeader.Open()
	// if err != nil {
	// 	return err
	// }
	defer profPicfile.Close()
	profPicPath, _ := helper.UploadFile(profPicfile, profPicHeader.Filename)
	ktpPicPath, _ := helper.UploadFile(ktpPicfile, profPicHeader.Filename)

	author.Name = name
	author.Email = email
	author.Password = password
	author.Username = username
	author.Experienced = experienced
	author.ProfPic = profPicPath
	author.KtpPic = ktpPicPath

	newAuthor, err := h.service.UpdateAuthorProfile(*author)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	auth_token, _ := h.authService.CreateAccessToken("author", newAuthor.ID)
	authorData := service.AuthorResponseFormatter(*newAuthor, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "author successfully updated", authorData)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) UpdateAuthorPassword(c echo.Context) error {
	authorID, _ := strconv.Atoi(c.Param("id"))
	author, _ := h.service.GetAuthorByID(uint(authorID))
	type Password struct {
		Password string `json:"password"`
	}
	password := new(Password)
	// password := c.FormValue("password")
	if err := c.Bind(password); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	author.Password = helper.GeneratePassword(password.Password)

	newAuthor, err := h.service.UpdateAuthorProfile(*author)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	auth_token, _ := h.authService.CreateAccessToken("author", newAuthor.ID)
	authorData := service.AuthorResponseFormatter(*newAuthor, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "author successfully updated", authorData)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) UpdateReaderProfile(c echo.Context) error {
	readerID, _ := strconv.Atoi(c.Param("id"))
	reader, _ := h.service.GetReader(uint(readerID))

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	username := c.FormValue("username")
	profPicHeader, _ := c.FormFile("prof_pic")
	// if err != nil {
	// 	return err
	// }
	profPicfile, _ := profPicHeader.Open()
	// if err != nil {
	// 	return err
	// }
	defer profPicfile.Close()
	profPicPath, _ := helper.UploadFile(profPicfile, profPicHeader.Filename)

	reader.Name = name
	reader.Email = email
	reader.Password = password
	reader.Username = username
	reader.ProfPic = profPicPath

	newReader, err := h.service.UpdateReaderProfile(*reader)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	auth_token, _ := h.authService.CreateAccessToken("reader", newReader.ID)
	data := service.ReaderResponseFormatter(*newReader, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "author successfully updated", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) UpdateReaderPassword(c echo.Context) error {
	readerID, _ := strconv.Atoi(c.Param("id"))
	reader, _ := h.service.GetReader(uint(readerID))
	type Password struct {
		Password string `json:"password"`
	}
	password := new(Password)
	// password := c.FormValue("password")
	if err := c.Bind(password); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	reader.Password = helper.GeneratePassword(password.Password)

	newReader, err := h.service.UpdateReaderProfile(*reader)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	auth_token, _ := h.authService.CreateAccessToken("author", newReader.ID)
	data := service.ReaderResponseFormatter(*newReader, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "author successfully updated", data)

	return c.JSON(http.StatusOK, response)
}
