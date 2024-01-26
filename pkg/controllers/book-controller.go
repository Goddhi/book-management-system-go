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
	newBooks := models.GetAllBooks() //  calling the GetAllBooks function gottten from the models file and returning the Book slice as newBook variable
	res, _ := json.Marshal(newBooks) // converts the data gotten from the database which is struct into a json format for the client or postman
	w.Header().Set("Content-Type", "pkglication/json") // converting the header to a json format
	w.WriteHeader(http.StatusOK) /// setting the response write header as status code 200
	w.Write(res)  // reponding with the data gotten from the database to postman or client
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
	res, _ := json.Marshal(bookDetails) // convertig the bookDetails(struct) variable into json
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)  // ouput is written to the client or postman
}

/// a function that creates a model/table in the database
func CreateBook(w http.ResponseWriter, r* http.Request){  /// This line defines the CreateBook function, which is an HTTP handler in a Go web application. It takes two parameters: w, an http.ResponseWriter that allows you to write an HTTP response to the client, and r, an *http.Request that contains all the details about the incoming HTTP request.
	CreateBook := &models.Book{}  // Here, a variable named CreateBook is created, which is a pointer to a new instance of the Book type from the models package. This struct will be populated with data from the HTTP request body.
	utils.Parsebody(r, CreateBook) // The Parsebody function from the utils package is called with the current HTTP request and the CreateBook pointer. This function is expected to read the request body, unmarshal the JSON data into the CreateBook struct, and prepare it for database insertion. Note that the actual conversion from struct to bytes for database operations would typically happen inside the CreateBook method of the Book struct, not in the Parsebody function.
	b := CreateBook.CreateBook()  /// This line calls the CreateBook method on the CreateBook struct (which is a bit confusing due to the naming). This method is presumably responsible for inserting the new book data into the database and returning the result. The result is stored in a variable b.
	res, _ := json.Marshal(b) /// The json.Marshal function is used to convert the b variable, which holds the newly created book record, into JSON format.
	w.WriteHeader(http.StatusOK)  // The WriteHeader method is called on the response writer w to send an HTTP 200 OK status code to the client. This informs the client that the request has been successfully processed.
	w.Write(res)  // the Write method is called on the response writer w to write the JSON-encoded book record (stored in res) to the HTTP response body. This sends the data back to the client, which could be Postman or any other HTTP client.
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
// the UpdateBook function is a combination of create function and getBookID function
func UpdateBook(w http.ResponseWriter, r *http.Request) { //
	var UpdateBook = &models.Book{} // Here, a new variable UpdateBook is declared and initialized as a pointer to a new instance of the Book struct from the models package. This struct will be used to unmarshal the JSON payload from the request body and may not necessarily create a new database table, but rather is used to hold the data you want to update in an existing record.
	utils.Parsebody(r, UpdateBook)  // the unmarshalling converts the JSON data from the request body into the struct
	vars := mux.Vars(r)  /// The mux.Vars function extracts the route variables from the request. The result is a map where the keys are the variable names defined in the URL pattern, and the values are the corresponding segments from the URL.
	bookId := vars["bookId"]  // This line retrieves the value of bookId from the route variables map. It corresponds to the {bookId} parameter in the route's URL.
	ID, err := strconv.ParseInt(bookId, 0,0)// The strconv.ParseInt function attempts to parse the bookId string as an integer. The second argument 0 tells the parser to infer the base from the string (allowing for decimal, hexadecimal, etc.), and the third argument 0 specifies that the result should fit into the default integer size.
	if err != nil {   // This checks if there was an error during the parsing of bookId. If there is an error, it prints an error message to the console
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)  // This line calls a function GetBookById from the models package, presumably fetching the book details from the database by the given ID, and returns the book details and a database handle or instance.
	if UpdateBook.Name != "" {
		bookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		bookDetails.Publication = UpdateBook.Publication
	}  // note the update logic is happening in Golang
	db.Save(&bookDetails)  //  This line presumably saves the updated bookDetails back to the database using the database handle db.
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res) //// writing the output of the code to postman
}