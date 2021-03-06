package repository

import (
	"fmt"

	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/api/entity"
	"gorm.io/gorm"
)

type Repository interface {
	AddAuthorRegistration(authorReg entity.AuthorRegistration) (entity.AuthorRegistration, error)
	AddReaderRegistration(readerReg entity.ReaderRegistration) (entity.ReaderRegistration, error)
	AddAdminRegistration(adminReg entity.AdminRegistration) (entity.AdminRegistration, error)
	GetAuthorRegistration(email string) (*entity.AuthorRegistration, error)
	GetReaderRegistration(email string) (*entity.ReaderRegistration, error)
	GetAdminRegistration(email string) (*entity.AdminRegistration, error)
	ConvertAuthortoUser(author entity.Author) *entity.User
	ConvertReadertoUser(reader entity.Reader) *entity.User
	ConvertAdmintoUser(admin entity.Admin) *entity.User
	GetCategory(id uint) (*entity.Category, error)

	// CRUD ADMIN
	AddAdmin(admin entity.Admin) (*entity.Admin, error)
	GetAdmin(username string) (*entity.Admin, error)

	// CRUD NEWS
	GetAllNews() ([]entity.News, error)
	GetAllNewsByCategory() ([]entity.News, error)
	GetNews(id uint) (*entity.News, error)
	GetNewsByCategoryID(categoryID uint) ([]entity.News, error)
	GetAllNewsByTrending() ([]entity.News, error)
	AddNews(news entity.News) (*entity.News, error)
	UpdateNews(news entity.News) (*entity.News, error)
	DeleteNews(id uint) (*entity.News, error)
	GetAllTrendingNews() ([]entity.News, error)
	GetAllHighlightNews(authorID uint) ([]entity.News, error)
	GetStatistic() (*entity.Statistic, error)
	AddNewsView(newsID uint, readerID uint) error
	AddNewsShare(newsID uint, readerID uint) error

	// CRUD AUTHOR
	GetAuthor(username string) (*entity.Author, error)
	GetAuthorByID(id uint) (*entity.Author, error)
	AddAuthor(author entity.Author) (*entity.Author, error)
	UpdateAuthor(author entity.Author) (*entity.Author, error)
	DeleteAuthor(id uint) (*entity.Author, error)
	GetNewsStatistic(newsID uint) (*entity.NewsStatistic, error)

	// CRUD READER
	GetReader(username string) (*entity.Reader, error)
	GetReaderByID(id uint) (*entity.Reader, error)
	AddReader(reader entity.Reader) (*entity.Reader, error)
	UpdateReader(reader entity.Reader) (*entity.Reader, error)
	DeleteReader(id uint) (*entity.Reader, error)

	AddComment(comment entity.NewsComment) (*entity.NewsComment, error)
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
func (r *repository) AddAdminRegistration(adminReg entity.AdminRegistration) (entity.AdminRegistration, error) {
	err := r.db.Create(&adminReg).Error
	if err != nil {
		return adminReg, err
	}
	return adminReg, nil
}

func (r *repository) GetAuthorRegistration(email string) (*entity.AuthorRegistration, error) {
	var registration entity.AuthorRegistration
	err := r.db.First(&registration, "email=?", email).Error
	if err != nil {
		return &registration, err
	}
	return &registration, nil
}

func (r *repository) GetReaderRegistration(email string) (*entity.ReaderRegistration, error) {
	var registration entity.ReaderRegistration
	err := r.db.First(&registration, "email=?", email).Error
	if err != nil {
		return &registration, err
	}
	return &registration, nil
}

func (r *repository) GetAdminRegistration(email string) (*entity.AdminRegistration, error) {
	var registration entity.AdminRegistration
	err := r.db.First(&registration, "email=?", email).Error
	if err != nil {
		return &registration, err
	}
	return &registration, nil
}
func (r *repository) AddAuthor(author entity.Author) (*entity.Author, error) {
	err := r.db.Create(&author).Error
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *repository) AddAdmin(admin entity.Admin) (*entity.Admin, error) {
	err := r.db.Create(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *repository) GetAuthor(username string) (*entity.Author, error) {
	var author entity.Author
	err := r.db.First(&author, "username=?", username).Error
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *repository) GetAdmin(username string) (*entity.Admin, error) {
	var admin entity.Admin
	err := r.db.First(&admin, "username=?", username).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
func (r *repository) ConvertAuthortoUser(author entity.Author) *entity.User {
	var user entity.User
	user.ID = author.ID
	user.Model = author.Model
	user.Name = author.Name
	user.Email = author.Email
	user.Password = author.Password
	user.Username = author.Username
	user.ProfPic = author.ProfPic
	user.Role = "author"
	return &user
}
func (r *repository) ConvertReadertoUser(reader entity.Reader) *entity.User {
	var user entity.User

	user.ID = reader.ID
	user.Model = reader.Model
	user.Name = reader.Name
	user.Email = reader.Email
	user.Password = reader.Password
	user.Username = reader.Username
	user.ProfPic = reader.ProfPic
	user.Role = "reader"
	return &user
}

func (r *repository) ConvertAdmintoUser(admin entity.Admin) *entity.User {
	var user entity.User

	user.ID = admin.ID
	user.Model = admin.Model
	user.Name = admin.Name
	user.Email = admin.Email
	user.Password = admin.Password
	user.Username = admin.Username
	user.ProfPic = admin.ProfPic
	user.Role = "admin"
	return &user
}

func (r *repository) AddNews(news entity.News) (*entity.News, error) {
	err := r.db.Create(&news).Error
	if err != nil {
		return nil, err
	}
	return &news, nil
}

func (r *repository) GetCategory(id uint) (*entity.Category, error) {
	var category entity.Category
	err := r.db.First(&category, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *repository) GetAuthorByID(id uint) (*entity.Author, error) {
	var author entity.Author
	err := r.db.First(&author, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *repository) GetAllNews() ([]entity.News, error) {
	var news []entity.News
	result := r.db.Preload("NewsReaders").Preload("NewsComments").Order("created_at desc").Find(&news)
	if result.Error != nil {
		return news, result.Error
	} else if result.RowsAffected < 1 {
		return news, fmt.Errorf("table is empty")
	}
	return news, nil
}
func (r *repository) GetAllNewsByCategory() ([]entity.News, error) {
	var news []entity.News
	result := r.db.Preload("NewsReaders").Preload("NewsComments").Order("category_id asc").Find(&news)
	if result.Error != nil {
		return news, result.Error
	} else if result.RowsAffected < 1 {
		return news, fmt.Errorf("table is empty")
	}
	return news, nil
}
func (r *repository) GetAllNewsByTrending() ([]entity.News, error) {
	statement1 := `
	SELECT news.id, SUM(total_view) FROM news JOIN news_readers
	ON news.id = news_readers.news_id
	GROUP BY news.id
	ORDER BY SUM(total_view) desc
	`

	var sortedTrendingList []entity.Trending
	result := r.db.Raw(statement1).Scan(&sortedTrendingList)
	// fmt.Printf("\n========= REPOSITORY GETALLTRENDINGNEWS: %+v \n\n", trending)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected < 1 {
		return nil, fmt.Errorf("table is empty")
	}

	var news []entity.News
	for _, trending := range sortedTrendingList {
		singleNews, _ := r.GetNews(uint(trending.ID))
		news = append(news, *singleNews)
	}
	return news, nil
}

func (r *repository) GetNews(id uint) (*entity.News, error) {
	var news entity.News
	err := r.db.Preload("NewsReaders").Find(&news, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &news, nil
}

func (r *repository) UpdateNews(news entity.News) (*entity.News, error) {
	err := r.db.Save(&news).Error
	if err != nil {
		return nil, err
	}
	return &news, nil
}

func (r *repository) DeleteNews(id uint) (*entity.News, error) {
	var news entity.News
	err := r.db.Find(&news, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Delete(&news, "id=?", id).Error
	// fmt.Printf("================REPOSITORY NEWS: %+v \n\n", news)
	if err != nil {
		return nil, err
	}
	return &news, nil
}

func (r *repository) GetNewsByCategoryID(categoryID uint) ([]entity.News, error) {
	var news []entity.News
	result := r.db.Preload("NewsComments").Find(&news, "category_id=?", categoryID)
	if result.Error != nil {
		return news, result.Error
	} else if result.RowsAffected < 1 {
		return news, fmt.Errorf("table is empty")
	}
	return news, nil
}

func (r *repository) GetAllTrendingNews() ([]entity.News, error) {
	statement1 := `
	SELECT news.id, SUM(total_view) FROM news JOIN news_readers
	ON news.id = news_readers.news_id
	GROUP BY news.id
	`

	var trending []entity.Trending
	result := r.db.Raw(statement1).Scan(&trending)
	// fmt.Printf("\n========= REPOSITORY GETALLTRENDINGNEWS: %+v \n\n", trending)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected < 1 {
		return nil, fmt.Errorf("table is empty")
	}

	maxView := MaximumView(trending)
	maxTrendingList := GetMaximumTrendingList(trending, maxView)

	var news []entity.News
	for _, trending := range maxTrendingList {
		singleNews, _ := r.GetNews(uint(trending.ID))
		news = append(news, *singleNews)
	}
	return news, nil
}

func (r *repository) GetAllHighlightNews(authorID uint) ([]entity.News, error) {
	statement1 := `
	SELECT news.id, SUM(total_view) FROM news JOIN news_readers
	ON news.id = news_readers.news_id
	WHERE news.author_id= ?
	GROUP BY news.id
	`

	var trending []entity.Trending
	result := r.db.Raw(statement1, authorID).Scan(&trending)
	// fmt.Printf("\n========= REPOSITORY GETALLTRENDINGNEWS: %+v \n\n", trending)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected < 1 {
		return nil, fmt.Errorf("table is empty")
	}

	maxView := MaximumView(trending)
	maxTrendingList := GetMaximumTrendingList(trending, maxView)

	var news []entity.News
	for _, trending := range maxTrendingList {
		singleNews, _ := r.GetNews(uint(trending.ID))
		news = append(news, *singleNews)
	}
	return news, nil
}

func MaximumView(trendings []entity.Trending) int {
	var max int
	for _, trending := range trendings {
		if trending.Sum > max {
			max = trending.Sum
		}
	}
	return max
}

func GetMaximumTrendingList(trendings []entity.Trending, maxView int) []entity.Trending {
	var maxTrending []entity.Trending
	for _, trending := range trendings {
		if trending.Sum == maxView {
			maxTrending = append(maxTrending, trending)
		}
	}
	return maxTrending
}

func (r *repository) UpdateAuthor(author entity.Author) (*entity.Author, error) {
	oldAuthor, _ := r.GetAuthorByID(author.ID)
	err := r.db.First(&oldAuthor).Error
	oldAuthor.Name = author.Name
	oldAuthor.Email = author.Email
	oldAuthor.Password = author.Password
	oldAuthor.Username = author.Username
	oldAuthor.Experienced = author.Experienced
	oldAuthor.ProfPic = author.ProfPic
	oldAuthor.KtpPic = author.KtpPic

	err = r.db.Save(&oldAuthor).Error
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *repository) DeleteAuthor(id uint) (*entity.Author, error) {
	var author entity.Author
	err := r.db.Find(&author, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Delete(&author, "id=?", id).Error
	// fmt.Printf("================REPOSITORY NEWS: %+v \n\n", news)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *repository) GetReader(username string) (*entity.Reader, error) {
	var reader entity.Reader
	err := r.db.Preload("NewsReaders").First(&reader, "username=?", username).Error
	if err != nil {
		return nil, err
	}
	return &reader, nil
}

func (r *repository) GetReaderByID(id uint) (*entity.Reader, error) {
	var reader entity.Reader
	err := r.db.Preload("NewsReaders").First(&reader, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &reader, nil
}

func (r *repository) AddReader(reader entity.Reader) (*entity.Reader, error) {
	err := r.db.Create(&reader).Error
	if err != nil {
		return nil, err
	}
	return &reader, nil
}
func (r *repository) UpdateReader(reader entity.Reader) (*entity.Reader, error) {
	oldReader, _ := r.GetReaderByID(reader.ID)
	err := r.db.First(&oldReader).Error
	oldReader.Name = reader.Name
	oldReader.Email = reader.Email
	oldReader.Password = reader.Password
	oldReader.Username = reader.Username
	oldReader.ProfPic = reader.ProfPic

	err = r.db.Save(&oldReader).Error
	if err != nil {
		return nil, err
	}
	return &reader, nil
}
func (r *repository) DeleteReader(id uint) (*entity.Reader, error) {
	var entity entity.Reader
	err := r.db.Find(&entity, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Delete(&entity, "id=?", id).Error
	// fmt.Printf("================REPOSITORY NEWS: %+v \n\n", news)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *repository) GetStatistic() (*entity.Statistic, error) {
	var totalAuthor int

	statement1 := `
	SELECT SUM(id) FROM authors
	`
	result := r.db.Raw(statement1).Scan(&totalAuthor)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected < 1 {
		return nil, fmt.Errorf("table is empty")
	}

	var totalReader int
	statement2 := `
	SELECT SUM(id) FROM readers
	`
	result = r.db.Raw(statement2).Scan(&totalReader)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected < 1 {
		return nil, fmt.Errorf("table is empty")
	}

	var totalNews int
	statement3 := `
	SELECT SUM(id) FROM news
	`
	result = r.db.Raw(statement3).Scan(&totalNews)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected < 1 {
		return nil, fmt.Errorf("table is empty")
	}

	var stat = entity.Statistic{}
	stat.TotalAuthor = totalAuthor
	stat.TotalReader = totalReader
	stat.TotalNews = totalNews

	return &stat, nil
}

func (r *repository) GetNewsStatistic(newsID uint) (*entity.NewsStatistic, error) {
	var totalShare int

	statement1 := `
	SELECT SUM(total_share) FROM news_readers 
	WHERE news_id=?
	`
	result := r.db.Raw(statement1, newsID).Scan(&totalShare)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected < 1 {
		return nil, fmt.Errorf("table is empty")
	}

	var totalReader int
	statement2 := `
	SELECT SUM(total_view) FROM news_readers 
	WHERE news_id=?
	`
	result = r.db.Raw(statement2, newsID).Scan(&totalReader)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected < 1 {
		return nil, fmt.Errorf("table is empty")
	}

	var totalComment int
	statement3 := `
	SELECT COUNT(id) FROM news_comments 
	WHERE news_id=?
	`
	result = r.db.Raw(statement3, newsID).Scan(&totalComment)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected < 1 {
		return nil, fmt.Errorf("table is empty")
	}

	var stat = entity.NewsStatistic{}
	stat.TotalShare = totalShare
	stat.TotalReader = totalReader
	stat.TotalComment = totalComment

	return &stat, nil
}

func (r *repository) AddNewsView(newsID uint, readerID uint) error {
	var totalLike int

	statement1 := `
	SELECT total_like FROM news_readers 
	WHERE news_id = ? AND reader_id = ?
	`
	result := r.db.Raw(statement1, newsID, readerID).Scan(&totalLike)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("table is empty")
	}

	if totalLike == 0 {
		statement2 := `
		UPDATE news_readers
		SET total_like = 1
		WHERE news_id = ? AND reader_id = ?
		`
		result := r.db.Raw(statement2, newsID, readerID)
		if result.Error != nil {
			return result.Error
		} else if result.RowsAffected < 1 {
			return fmt.Errorf("total_view update failed")
		}
	}
	// } else {
	// 	return fmt.Errorf("only 1 like allowed for 1 reader")
	// }
	return nil

}

func (r *repository) AddComment(comment entity.NewsComment) (*entity.NewsComment, error) {
	err := r.db.Create(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *repository) AddNewsShare(newsID uint, readerID uint) error {
	var totalShare int

	statement1 := `
	SELECT total_share FROM news_readers 
	WHERE news_id = ? AND reader_id = ?
	`
	result := r.db.Raw(statement1, newsID, readerID).Scan(&totalShare)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("table is empty")
	}

	if totalShare == 0 {
		statement2 := `
		UPDATE news_readers
		SET total_share = 1
		WHERE news_id = ? AND reader_id = ?
		`
		result := r.db.Raw(statement2, newsID, readerID)
		if result.Error != nil {
			return result.Error
		} else if result.RowsAffected < 1 {
			return fmt.Errorf("total_view update failed")
		}
	} else {
		return fmt.Errorf("only 1 share allowed for 1 reader")
	}
	return nil
}
