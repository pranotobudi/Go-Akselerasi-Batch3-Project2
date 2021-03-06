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
type ResponseAdminRegistration struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	ProfPic   string `json:"prof_pic"`
	AuthToken string `json:"auth_token"`
}
type ResponseAuthor struct {
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
type ResponseReader struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	ProfPic   string `json:"prof_pic"`
	AuthToken string `json:"auth_token"`
}
type ResponseAdmin struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	ProfPic   string `json:"prof_pic"`
	AuthToken string `json:"auth_token"`
}

type ResponseUser struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	ProfPic   string `json:"prof_pic"`
	Role      string `json:"role"`
	AuthToken string `json:"auth_token"`
}

type ResponseNews struct {
	ID           uint                 `json:"id"`
	AuthorID     int                  `json:"author_id"`
	Author       entity.Author        `json:"author"`
	CategoryID   int                  `json:"category_id"`
	Category     entity.Category      `json:"category"`
	Title        string               `json:"title"`
	Content      string               `json:"content"`
	ImageUrl     string               `json:"image_url"`
	NewsComments []entity.NewsComment `json:"news_comments"`
	NewsReaders  []entity.NewsReaders `json:"news_readers"`
}
type ResponseStatistic struct {
	TotalAuthor int `json:"total_author`
	TotalReader int `json:"total_reader`
	TotalNews   int `json:total_news`
}
type ResponseNewsStatistic struct {
	TotalShare   int `json:"total_share`
	TotalReader  int `json:"total_reader`
	TotalComment int `json:total_comment`
}

type ResponseComment struct {
	ReaderID int    `json:"reader_id`
	NewsID   int    `json:"news_id`
	Comment  string `json:comment`
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

func AdminRegistrationResponseFormatter(adminReg entity.AdminRegistration, auth_token string) ResponseAdminRegistration {
	formatter := ResponseAdminRegistration{
		ID:        adminReg.ID,
		Name:      adminReg.Name,
		Email:     adminReg.Email,
		Password:  adminReg.Password,
		Username:  adminReg.Username,
		ProfPic:   adminReg.ProfPic,
		AuthToken: auth_token,
	}
	return formatter
}

func AuthorResponseFormatter(authorReg entity.Author, auth_token string) ResponseAuthor {
	formatter := ResponseAuthor{
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

func ReaderResponseFormatter(readerReg entity.Reader, auth_token string) ResponseReader {
	formatter := ResponseReader{
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

func AdminResponseFormatter(adminReg entity.Admin, auth_token string) ResponseAdmin {
	formatter := ResponseAdmin{
		ID:        adminReg.ID,
		Name:      adminReg.Name,
		Email:     adminReg.Email,
		Password:  adminReg.Password,
		Username:  adminReg.Username,
		ProfPic:   adminReg.ProfPic,
		AuthToken: auth_token,
	}
	return formatter
}

func UserResponseFormatter(user entity.User, auth_token string) ResponseUser {
	formatter := ResponseUser{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Username:  user.Username,
		ProfPic:   user.ProfPic,
		Role:      user.Role,
		AuthToken: auth_token,
	}
	return formatter
}

func NewsResponseFormatter(news entity.News, author entity.Author, category entity.Category) ResponseNews {
	formatter := ResponseNews{
		ID:           news.ID,
		AuthorID:     news.AuthorID,
		Author:       author,
		CategoryID:   news.CategoryID,
		Category:     category,
		Title:        news.Title,
		Content:      news.Content,
		ImageUrl:     news.ImageUrl,
		NewsComments: news.NewsComments,
		NewsReaders:  news.NewsReaders,
	}
	return formatter
}

func StatisticResponseFormatter(stat entity.Statistic) ResponseStatistic {
	formatter := ResponseStatistic{
		TotalAuthor: stat.TotalAuthor,
		TotalReader: stat.TotalReader,
		TotalNews:   stat.TotalNews,
	}
	return formatter
}
func NewsStatisticResponseFormatter(stat entity.NewsStatistic) ResponseNewsStatistic {
	formatter := ResponseNewsStatistic{
		TotalShare:   stat.TotalShare,
		TotalReader:  stat.TotalReader,
		TotalComment: stat.TotalComment,
	}
	return formatter
}

func CommentResponseFormatter(comment entity.NewsComment) ResponseComment {
	formatter := ResponseComment{
		ReaderID: comment.ReaderID,
		NewsID:   comment.NewsID,
		Comment:  comment.Comment,
	}
	return formatter
}
