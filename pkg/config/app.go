package config

// these packages are used to connect to mysql database

import (
	// "github.com/jinzhu/gorm" for gorm v1
	//"github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/mysql" // for gorm V1
	"gorm.io/driver/mysql" // for gorm V2
	"gorm.io/gorm" // for gorm V2

)

// the main reason for this file is to return a variable called db so other files can interact with DB
//
var (
	db * gorm.DB
)

func Connect(){  // this functions helps to open connection with the mysql database 
	// connecting to mysql database using gorm V1

	// d, err := gorm.Open("mysql", "goddhi:goddhipassword%4012@/simplerest?charset=utf8&parseTime=True&loc=Local") // specifying the database username, password, table name
	// if err != nil { // if an error occurs 
	// 	// In this example, if gorm.Open fails to  clonnect to the database, the function will panic, printing the error message and stopping the execution of the program.
	// 	panic(err)  // panic is called, it immediately stops the execution of the function in which it was called and begins to unwind the stack, running any deferred functions along the way. If the panic is not recovered, it eventually stops the entire program. The statement panic(err) is commonly used to handle unexpected errors where the program cannot (or should not) continue its execution. Here's a basic breakdown:   err: This represents an error value. In Go, error handling is usually done by returning an error type as the last value in a function's return list. If an error is not nil, it indicates that something went wrong.  panic(err): This line triggers a panic with the error err. It's often used in situations where the error is critical and there is no way for the program to recover or proceed further. For example, if your program absolutely needs to read a configuration file to start, and it can't find that file, you might use panic to halt the program. 
	// }   

	// connecting to mysql database using gorm V2
	dsn := "goddhi:goddhipassword@tcp(localhost:3306)/simplerest?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err) // // panic is called, it immediately stops the execution of the function in which it was called and begins to unwind the stack, running any deferred functions along the way. If the panic is not recovered, it eventually stops the entire program. The statement panic(err) is commonly used to handle unexpected errors where the program cannot (or should not) continue its execution. Here's a basic breakdown:   err: This represents an error value. In Go, error handling is usually done by returning an error type as the last value in a function's return list. If an error is not nil, it indicates that something went wrong.  panic(err): This line triggers a panic with the error err. It's often used in situations where the error is critical and there is no way for the program to recover or proceed further. For example, if your program absolutely needs to read a configuration file to start, and it can't find that file, you might use panic to halt the program.
	}
	db = d 
} 
/// 
func GetDb() *gorm.DB{  // this function returns db which connect to the mysql database running locally  to other folders
	return db
}

/// testing connectivity 