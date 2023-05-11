package models

import (
	"github.com/Sanjay-21pw/go-bookstore-new/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) createBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func getAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func getBookById(Id int64) (*Book, *gorm.DB) {
	var getbook Book
	db := db.Where("ID=?", Id).Find(&getbook)
	return &getbook, db
}

func deleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
