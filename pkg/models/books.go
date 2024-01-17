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
	
}
