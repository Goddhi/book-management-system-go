package routes

import (
	"github.com/gorilla/mux"
	// "github.com/goddhi/book-maagement-system/pkg/controllers"
	"github.com/goddhi/mnt/c/Users/Owner/go/book-maagement-system/pkg/controllers"
	
)

var RegisterBookStoreRoutes = func(router *mux.Router) {  // function for the route
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")  // create a book
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookBYID).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("UPDATE")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
} 

