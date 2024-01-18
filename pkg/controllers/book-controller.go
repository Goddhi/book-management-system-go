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
	vars := mux.Vars(r) // assesing the book id t
}
