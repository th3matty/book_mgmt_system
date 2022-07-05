package models

import (
	"github.com/jinzhu/gorm"
	"github.com/th3matty/book_mgmt_system/tree/main/pkg/config"
)

var db * gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string	`json:"publication"` 
}

func init(){
	config.Connect()
	db = config.GetDB()
	// Migrate the schema
	db.AutoMigrate(&Book{})
}

// method CreateBook() is a func with special receiver argument
// CreateBook method has receiver of type Book namend b
func (b *Book) CreateBook() *Book{
	//return is of type book 
	// NewRecord and Create comes from gorm
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// returns a slice from type Book
func GetAllBooks()[]Book{
	var Books []Book
	db.Find(&Books)
	return Books
}


func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}