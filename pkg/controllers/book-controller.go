package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/th3matty/book_mgmt_system/tree/main/pkg/utils"
	"github.com/th3matty/book_mgmt_system/tree/main/pkg/models"
)

// declare new variable newBook after type struct Book in models
var newBook models.Book

// pointer to request receiving from user
func GetBook(w http.ResponseWriter , r *http.Request){
	// get method from models 
	newBooks := models.GetAllBooks()
	// assign marshaled data to response!
	res, _ := json.Marshal(newBooks)
	// set the Header
	w.Header().Set("Content-Type", "pkglication/json")
	// get the server status 200
	w.WriteHeader(http.StatusOK)
	// send response to the frontend e.g.
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	// Paths can have variables.
	// mux.Vars() receives the argument r.
	vars := mux.Vars(r)
	// is there such path with bookId?
	bookId := vars["bookId"]
	// convert string to int and store it in ID
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	// we dont want to use the return db value from the method (->models)GetBookById
	// thats why we use the underscore singn " _ "
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	// set the Header
	w.Header().Set("Content-Type", "pkglication/json")
	// get the status
	w.WriteHeader(http.StatusOK)
	// return to frontend
	w.Write(res)
}

func CreateBook(w http.ResponseWriter , r *http.Request){
	// assign type struct Book from models to CreateBook 
	CreateBook := &models.Book{}
	// Parsebody receives 2 arguments. request , and the interface CreateBook
	utils.ParseBody(r, CreateBook)
	// the database send us the record which has been saved through method CreateBook()
	b := CreateBook.CreateBook()
	//convert it into json and use the record from DB -> b
	res , _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter , r *http.Request){
	// check the path for variables
	vars := mux.Vars(r)
	// is there category "ID" present? if yes save to new variable
	bookId := vars["ID"]
	// convert string to int and store it in ID
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing in deleteBook method")
	}
	// call deleteBook method from models
	book := models.DeleteBook(ID)

	res , _ := json.Marshal(book)
	// set Header, get status and send response to the frontend
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// check strconv.ParseInt arguments!!
// what is "pkglication/json"

func UpdateBook(w http.ResponseWriter , r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	// get our id 
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	// now find the bookId in the DB
	bookDetails , db := models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}