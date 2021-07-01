package repository

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/api/entity"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/helper"
	"gorm.io/gorm"
)

func DBSeed(db *gorm.DB) error {
	AuthorDataSeed(db)
	AdminDataSeed(db)
	ReaderDataSeed(db)
	CategoryDataSeed(db)
	NewsDataSeed(db)
	NewsReaderDataSeed(db)
	// NewsCommentDataSeed(db)
	return nil
}

func AuthorDataSeed(db *gorm.DB) {
	statement := "INSERT INTO authors (name, email, password, username, prof_pic, ktp_pic, experienced, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	db.Exec(statement, "writer1", "writer1@gmail.com", helper.GeneratePassword("writer1"), "writer1", "http://prof_pic_url_writer1.jpg", "http://ktp_pic_url1_writer1.jpg", true, faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "writer2", "writer2@gmail.com", helper.GeneratePassword("writer2"), "writer2", "http://prof_pic_url_writer2.jpg", "http://ktp_pic_url1_writer2.jpg", true, faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "writer3", "writer3@gmail.com", helper.GeneratePassword("writer3"), "writer3", "http://prof_pic_url_writer3.jpg", "http://ktp_pic_url1_writer3.jpg", false, faker.Timestamp(), faker.Timestamp())

}
func AdminDataSeed(db *gorm.DB) {
	statement := "INSERT INTO admins (name, email, password, username, prof_pic, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	db.Exec(statement, "admin", "admin@gmail.com", helper.GeneratePassword("admin"), "admin", "http://prof_pic_url_admin.jpg", faker.Timestamp(), faker.Timestamp())
}

func ReaderDataSeed(db *gorm.DB) {
	statement := "INSERT INTO readers (name, email, password, username, prof_pic, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	db.Exec(statement, "reader1", "reader1@gmail.com", helper.GeneratePassword("reader1"), "reader1", "http://prof_pic_url_reader1.jpg", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "reader2", "reader2@gmail.com", helper.GeneratePassword("reader2"), "reader2", "http://prof_pic_url_reader2.jpg", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "reader3", "reader3@gmail.com", helper.GeneratePassword("reader3"), "reader3", "http://prof_pic_url_reader3.jpg", faker.Timestamp(), faker.Timestamp())

}

func CategoryDataSeed(db *gorm.DB) {
	statement := "INSERT INTO categories (admin_id, name, created_at, updated_at) VALUES (?, ?, ?, ?)"

	db.Exec(statement, 1, "category1", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 1, "category2", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 1, "category3", faker.Timestamp(), faker.Timestamp())
}

func NewsDataSeed(db *gorm.DB) {
	statement := "INSERT INTO news (author_id, category_id, title, content, image_url, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	db.Exec(statement, 1, 1, "title1", "content1", "http://image_url_title1", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 1, 2, "title2", "content2", "http://image_url_title2", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 1, 3, "title3", "content3", "http://image_url_title3", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 1, "title4", "content4", "http://image_url_title4", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 2, "title5", "content5", "http://image_url_title5", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 3, "title6", "content6", "http://image_url_title6", faker.Timestamp(), faker.Timestamp())
}

func NewsCommentDataSeed(db *gorm.DB) {
	statement := "INSERT INTO news_comments (reader_id, news_id, comment, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"

	db.Exec(statement, 1, 1, "comment1", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 1, "comment2", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 3, 1, "comment3", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 1, 2, "comment4", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 2, "comment5", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 3, 2, "comment6", faker.Timestamp(), faker.Timestamp())
}

func NewsReaderDataSeed(db *gorm.DB) {
	// statement := "INSERT INTO news_readers (news_id, reader_id, like, share, view) VALUES (?, ?, ?, ?, ?)"
	statement := "INSERT INTO news_readers (reader_id, news_id, total_like, total_share, total_view, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	db.Exec(statement, 1, 1, 1, 1, 1, faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 1, 1, 1, 1, faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 3, 1, 1, 1, 1, faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 1, 2, 1, 1, 1, faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 2, 1, 1, 1, faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 3, 2, 1, 1, 1, faker.Timestamp(), faker.Timestamp())
}

func InitDBTable(db *gorm.DB) {
	// db.AutoMigrate(&User{}, &Event{}, &Transaction{}, &Registration{})
	db.AutoMigrate(entity.AuthorRegistration{}, entity.ReaderRegistration{}, entity.AdminRegistration{}, entity.Author{}, entity.Admin{}, entity.Reader{}, entity.Category{}, entity.News{}, entity.NewsReaders{})

	// Create Fresh AuthorRegistration Table
	if (db.Migrator().HasTable(&entity.AuthorRegistration{})) {
		fmt.Println("AuthorRegistration table exist")
		db.Migrator().DropTable(&entity.AuthorRegistration{})
	}
	db.Migrator().CreateTable(&entity.AuthorRegistration{})

	// Create Fresh ReaderRegistration Table
	if (db.Migrator().HasTable(&entity.ReaderRegistration{})) {
		fmt.Println("ReaderRegistration table exist")
		db.Migrator().DropTable(&entity.ReaderRegistration{})
	}
	db.Migrator().CreateTable(&entity.ReaderRegistration{})

	// Create Fresh AdminRegistration Table
	if (db.Migrator().HasTable(&entity.AdminRegistration{})) {
		fmt.Println("AdminRegistration table exist")
		db.Migrator().DropTable(&entity.AdminRegistration{})
	}
	db.Migrator().CreateTable(&entity.AdminRegistration{})

	// Create Fresh Author Table
	if (db.Migrator().HasTable(&entity.Author{})) {
		fmt.Println("Author table exist")
		db.Migrator().DropTable(&entity.Author{})
	}
	db.Migrator().CreateTable(&entity.Author{})

	// Create Fresh Admin Table
	if (db.Migrator().HasTable(&entity.Admin{})) {
		fmt.Println("Admin table exist")
		db.Migrator().DropTable(&entity.Admin{})
	}
	db.Migrator().CreateTable(&entity.Admin{})

	// Create Fresh Reader Table
	if (db.Migrator().HasTable(&entity.Reader{})) {
		fmt.Println("Reader table exist")
		db.Migrator().DropTable(&entity.Reader{})
	}
	db.Migrator().CreateTable(&entity.Reader{})

	// Create Fresh Category Table
	if (db.Migrator().HasTable(&entity.Category{})) {
		fmt.Println("Category table exist")
		db.Migrator().DropTable(&entity.Category{})
	}
	db.Migrator().CreateTable(&entity.Category{})

	// Create Fresh News Table
	if (db.Migrator().HasTable(&entity.News{})) {
		fmt.Println("News table exist")
		db.Migrator().DropTable(&entity.News{})
	}
	db.Migrator().CreateTable(&entity.News{})

	// Create Fresh NewsComment Table
	// if (db.Migrator().HasTable(&entity.NewsComment{})) {
	// 	fmt.Println("NewsComment table exist")
	// 	db.Migrator().DropTable(&entity.NewsComment{})
	// }
	// db.Migrator().CreateTable(&entity.NewsComment{})

	// Create Fresh NewsReaders Table
	if (db.Migrator().HasTable(&entity.NewsReaders{})) {
		fmt.Println("NewsReaders table exist")
		db.Migrator().DropTable(&entity.NewsReaders{})
	}
	db.Migrator().CreateTable(&entity.NewsReaders{})

}
