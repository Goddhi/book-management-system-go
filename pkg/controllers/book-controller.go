package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/goddhi/book-management-system/pkg/models"
	"github.com/goddhi/book-management-system/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book // referencing the stuct Book from the models folder

// creating a function that perForm get request of the list of books on the database and respond with the list of books
func GetBook(w http.ResponseWriter, r *http.Request) {  // the GetBook function gets executed when a user points to /book route 
	newBooks := models.GetAllBooks() // passing the GetAllBooks function gottten from the modeels file as a variable into newBooks
	res, _ := json.Marshal(newBooks) // converts the data gotten from the databse into a json format
	w.Header().Set("Content-Type", "pkglication/json") // converting the header to a json format
	w.WriteHeader(http.StatusOK) /// setting the response write header as status code 200
	w.Write(res)  // reponding with the data gotten from the database to postman
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	// basically the below code means we are passisng the book id gotten from the request into the vars varible 
	vars := mux.Vars(r) //This function is provided by the gorilla/mux package. It takes an http.Request object as an argument, which represents the current HTTP request being processed by the handler...... The line var := mux.Vars(r) in the context of a Go web server using the gorilla/mux package is used to extract the route variables from the URL
	bookId := vars["bookId"] // the line bookId := vars["bookId"] is used to retrieve a path parameter from the URL
	ID, err := strconv.ParseInt(bookId,0,0) // connverting the string gotten into an integer
	if err != nil {
		fmt.Println("error while parsing")  /// log out error when there is an error getting the book ID
	}
	bookDetails, _  := models.GetBookById(ID)  // returning the varible getBook and ignoring the db varible returned from the GetBookById function in module folder
	

}
