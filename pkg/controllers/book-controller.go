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
	res, _ := json.Marshal(bookDetails) // convertig the bookDetails(string) variable into json
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/// a function that creates a model/table in the database
func CreateBook(w http.ResponseWriter, r* http.Request){
	CreateBook := &models.Book{}  // passing the Book(struct) into a variable called CreateBook
	utils.Parsebody(r, CreateBook) // passing the CreateBook(struct) as parameter so it that it can be converted into bytes of code easily understood from the databse
	b := CreateBook.CreateBook()  /// returning the value of what was gotten from the datbase
	res, _ := json.Marshal(b) /// converting the data gotten from the database into json 
	w.WriteHeader(http.StatusOK)
	w.Write(res)  // printing the result of the data to postman 
}

func DeleteBookId(w http.ResponseWriter, r* http.Request){
	vars := mux.Vars(r)  // using vars to access the request from the data
	bookId := vars["bookId"] // placing the value into a variable
	ID, err := strconv.ParseInt(bookId, 0,0) //converting the value of bookid into an int
	if err != nil {
		fmt.Print("error whilw passimg")
	}
	book := models.DeleteBook(ID) // passing the return variable as data    
	res, _ := json.Marshal(book) // converting the value into json
	w.Header().Set("Content-Type",  "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


/// creating a function that update the book data
func UpdateBook(w http.ResponseWriter, r *http.Request) { //
	var UpdateBook = &models.Book{} // passing the book(struct) as a variable so we can create a new database table 
	utils.Parsebody(r, UpdateBook)  // converting the  UpdateBook from a struct to byte so the database can utilize it
	vars := mux.Vars(r)  /// passing the request value into a variable called vars
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)// converting the string into an integer
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if UpdateBook.Name != "" {
		bookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		bookDetails.Publication = UpdateBook.Publication
	}
}