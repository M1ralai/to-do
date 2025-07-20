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

func GetTodo() ([]*Todo, error) {
	db, err := sql.Open("sqlite3", "./db/mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r, err := db.Query(getTodo)
	if err != nil {
		log.Println(err)
	}
	defer r.Close()
	var todos []*Todo
	for r.Next() {
		var todo Todo
		err := r.Scan(&todo.Id, &todo.Title, &todo.Desc, &todo.CreationTime, &todo.Deadline)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}

func DeleteTodo(id int) {
	db, err := sql.Open("sqlite3", "./db/mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(deleteTodo, id)
	if err != nil {
		log.Println("paniced")
		log.Println(err)
	}
}
