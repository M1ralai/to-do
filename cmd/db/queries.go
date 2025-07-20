package db

type Todo struct {
	Id           string `json:"id" db:"id"`
	Title        string `json:"title" db:"title"`
	Desc         string `json:"description" db:"desc"`
	CreationTime string `json:"creation_time" db:"creation_date"`
	Deadline     string `json:"deadline" db:"deadline"`
}

var clearDb = `DROP TABLE IF EXISTS todo`
var createDb = `CREATE TABLE IF NOT EXISTS todo (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title VARCHAR(20),
	desc VARCHAR(200),
	creation_date DEFAULT CURRENT_TIMESTAMP,
	deadline TIMESTAMP
);`

var insertTodo = `INSERT INTO todo (title, desc, deadline) VALUES (?, ?, ?);`

var getTodo = `SELECT * FROM todo ;`

var deleteTodo = `DELETE FROM todo WHERE id = ?;`
