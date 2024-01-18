package main

import (
	"log"
	"net/http"
	"github.com/goddhi/book-management-system/pkg/routes" // importing the route folder
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()  // initializing the router
	routes.RegisterBookStoreRoutes(r)  //// passing the mux.Router as a parameter also as an instance for mux.NewRouter
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r)) /// the server listens on port 9010 and takes in the router(r) as a parameter. also if there is an error in logs out the error
	
}