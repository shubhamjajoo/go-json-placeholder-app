package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDBB(dbname string) {
	var err error
	if dbname == "postgres" {
		DB, err = sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable password=admin host=localhost port=5432")
		if err != nil {
			fmt.Println("Could not connect to DB", err.Error())
			return
		}
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTable()
	seedUsers()
	seedTodos()
}

func createTable() {
	createUsersTable := `CREATE TABLE  IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		username TEXT NOT NULL UNIQUE
	)`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to  create user table")
	}

	createTodoTable := `CREATE TABLE  IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		completed BOOLEAN,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`
	_, err = DB.Exec(createTodoTable)
	if err != nil {
		panic("failed to  create todos table")
	}
}
