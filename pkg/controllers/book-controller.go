package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/th3matty/book_mgmt_system/pkg/utils"
	"github.com/th3matty/book_mgmt_system/pkg/models"
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
}