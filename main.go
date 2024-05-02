package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func generateRandomString(length int) string {
	const charset = "abcde"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

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

	var n int

	_, err = fmt.Scan(&n)

	if err != nil {
		return
	}
	var ages []int = make([]int, n)
	var names []string = make([]string, n)
	var rolls []string = make([]string, n)
	var addresses []string = make([]string, n)

	for i := 0; i < n; i++ {
		ages[i] = i + 1
	}

	for i := 0; i < n; i++ {
		names[i] = generateRandomString(8)
		addresses[i] = generateRandomString(8)
		rolls[i] = generateRandomString(3)
	}
	for i := 0; i < n; i++ {
		_, err = db.Exec("INSERT INTO golang (name, age, roll, address) VALUES (?, ?, ?, ?)", names[i], ages[i], rolls[i], addresses[i])
		if err != nil {
			log.Fatal("error")
			return
		}
	}
	fmt.Println("Data inserted successfully")
}

