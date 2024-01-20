package models

import (
	"github.com/jinzhu/gorm"
	"github.com/goddhi/book-management-system/pkg/config" // importing the config folder so that this file can connect to the database
)

var db *gorm.DB


// this struct are created as module or table to store things in the database
type Book struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init()  {  // using the init function to initialize database with the connect function in the config folder 
	config.Connect()  // this function is gotten from config folder which connect to the database
	db = config.GetDb()  // re-storing db gotten from the config file
	db.AutoMigrate(&Book{})  /// auto migrate with the empty Book struct
}

/// All functions below will be use to interact with the database

func (b *Book) CreateBook() *Book{ // passing the Book(struct) as a parameter and returning the Book(struct) as b. 
	db.NewRecord(b) // db which is used to connect to the database and NewRecord which is gotten from the gorm(ORM) is used to create a record based on the Book struct 
	db.Create(&b) // db which is used to connect to the database and NewRecord which is gotten from the gorm(ORM) is used to create a table based on the Book struct 
	return b  /// returning the Book struct as b
}

func GetAllBooks() []Book{  //  function that get all books as slice(list) from the Book(stuct)
	var Books []Book // creating a varaible Books with a data type of slice
	db.Find(&Books)  // funding the lsit of Books in the database
	return Books  // return a slice of Books
}


func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book  // passing getBook variable as Book(struct)
	db := db.Where("ID=?", Id).Find(&getBook) // using the db which is used to connect to the database to find ID = Id and then find the particular book 
	return &getBook, db // returning getBook variable pointed to * Book(struct) and the value of the ID(db)
}


func DeleteBook(ID int64) Book {  // function that delete a book ID and returns the Book(struct)   
	var book Book // assingning book variable to the Book struct
	db.Where("ID=?", ID).Delete(book) // using  db to delete a book ID
	return book
}

