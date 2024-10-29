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
	db, err := sql.Open("mysql", "akash:password@(localhost:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	createUsers(db)
	insertUser(db)
	printUsers(db)

	if err != nil {
		log.Fatal(err)
	}
}

func createUsers(db *sql.DB) {
	dropTableIfExist(db)
	query := `
	CREATE TABLE users (
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func dropTableIfExist(db *sql.DB) {
	query := `DROP TABLE IF EXISTS users;`
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

	if _, err := result.LastInsertId(); err != nil {
		log.Fatal(err)
		return
	}
}

func printUsers(db *sql.DB) {
	type user struct {
		id         int
		username   string
		password   string
		created_at time.Time
	}

	query := `select id, username, password, created_at from users;`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rows.Close()
	var users []user
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password, &u.created_at)
		if err != nil {
			log.Fatal(err)
			return
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return
	}
	for _, u := range users {
		fmt.Printf("id = %d\nusername = %s\npassword=%s\ntime=%s", u.id, u.username, u.password, u.created_at)
	}

}
