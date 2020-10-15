package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	} else {
		fmt.Println("conn success")
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO user (name, age) VALUES(?, ?)") // ? = placeholder
	if err != nil {
		fmt.Println(err.Error()) // proper error handling instead of panic in your app
	} else {
		fmt.Println("prepare success")
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT name,age FROM user WHERE age = ?")
	if err != nil {
		fmt.Println(err.Error()) // proper error handling instead of panic in your app
	} else {
		fmt.Println("prepare success")
	}
	defer stmtOut.Close()

	// Insert square numbers for 0-24 in the database
	/*for i := 0; i < 25; i++ {
		name := fmt.Sprintf("test%d", i)
		_, err = stmtIns.Exec(name, i) // Insert tuples (i, i^2)
		if err != nil {
			fmt.Println(err.Error()) // proper error handling instead of panic in your app
		} else {
			fmt.Println("insert success")
		}
	}*/

	var name string // we "scan" the result in here
	var age int

	// Query the age of 20
	err = stmtOut.QueryRow(20).Scan(&name, &age) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The age of 20 is: %s:%d", name, age)

	// Query the age of 15
	err = stmtOut.QueryRow(15).Scan(&name, &age) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The age of 15 is: %s:%d", name, age)
}