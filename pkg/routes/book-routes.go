package routes

import (
	"github.com/gorilla/mux"
	"github.com/goddhi/book-management-system/pkg/controllers"

) 


var RegisterBookStoreRoutes = func(router *mux.Router) {  // function for the route  the mux.Router is taken as a paramter
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookBYID).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("UPDATE")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")


	
} 




