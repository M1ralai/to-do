package db

type Todo struct {
	Id           string `json:"ID"`
	Title        string `json:"title"`
	Desc         string `json:"description"`
	CreationTime string `json:"creation_time"`
	Deadline     string `json:"deadline"`
}

func newTodo(id string, title string, desc string, creationtime string, deadline string) *Todo {
	return &Todo{
		Id:           id,
		Title:        title,
		Desc:         desc,
		CreationTime: creationtime,
		Deadline:     deadline,
	}
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

var getTodo = `SELECT * FROM todo WHERE id = ?;`
