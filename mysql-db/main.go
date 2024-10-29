package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// don't put username and password in the code, this is just an example code.
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	v, err := db.Exec("SELECT 1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(v)
	fmt.Println("accessed")
}

func createUsers(db *sql.DB) {
	query := `
	CREATE TABLE users (
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
	);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
		return
	}
}

func insertUser(db *sql.DB) {
	username := "john"
	password := "secret"
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)

	if err != nil {
		log.Fatal(err)
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(id)
}
