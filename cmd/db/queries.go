package db

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
