package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "golang.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS golang (id INTEGER PRIMARY KEY, name TEXT, age INTEGER, roll TEXT, address TEXT)")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec("INSERT INTO golang (name, age, roll, address) VALUES (?, ?, ?, ?)", "ayush", 19, "100000", "kathmandu")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data inserted successfully")
}

