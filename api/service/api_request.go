package service

type RequestAuthor struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Username    string `json:"username"`
	ProfPic     string `json:"prof_pic"`
	KtpPic      string `json:"ktp_pic"`
	Experienced bool   `json:"experienced"`
}
type RequestReader struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	ProfPic  string `json:"prof_pic"`
}

type RequestAdmin struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	ProfPic  string `json:"prof_pic"`
}

type RequestUserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RequestNews struct {
	AuthorID   int    `json:"author_id"`
	CategoryID int    `json:"category_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	ImageUrl   string `json:"image_url"`
}
