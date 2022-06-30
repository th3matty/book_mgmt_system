package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/th3matty/book_mgmt_system/pkg/routes"
	"github.com/th3matty/book_mgmt_system/tree/main/pkg/routes"
)

func main(){
	r:= mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	// create server to this port 
	// comes from the http pkg
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}