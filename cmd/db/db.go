package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() {
	db, err := sql.Open("sqlite3", "./db/mydb.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("successfully connected to a database")
	_, err = db.Exec(clearDb)
	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec(createDb)
	if err != nil {
		log.Println(err)
	}

	defer db.Close()
}

// first variable will be title with max 20 char size, desc will be description with 200 char size and deadline must be written like time.Date(YYYY, MM, DD, HH, MM, SS, NN, time.UTC)
func CreateTodo(title string, desc string, deadline time.Time) {
	db, err := sql.Open("sqlite3", "./db/mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	layout := "2006-01-02 15:04:05"
	formatted := deadline.Format(layout)

	_, err = db.Exec(insertTodo, title, desc, formatted)
	if err != nil {
		log.Println(err)
	}
}

func GetTodo(id int) *Todo {
	var data [5]string
	db, err := sql.Open("sqlite3", "./db/mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r := db.QueryRow(getTodo, id)
	err = r.Scan(&data[0], &data[1], &data[2], &data[3], &data[4])
	if err != nil {
		log.Fatal(err)
	}
	t := newTodo(data[0], data[1], data[2], data[3], data[4])
	return t
}
