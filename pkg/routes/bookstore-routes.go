package routes

import (
	"github.com/gorilla/mux"
	"github.com/th3matty/book_mgmt_system/pkg/controllers"
)

// flow of data
// routes give controll to controller, and controllers give controlls to model(book.go)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookid}",controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{book.id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")
}