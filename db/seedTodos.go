package db

import (
	"fmt"
	"go-json-placeholder-app/models"
	"go-json-placeholder-app/services"
)

func seedTodos() {
	todos, err := services.FetchTodos()
	if err != nil {
		fmt.Println("Failed to fetch todos")
		return
	}
	dumpTodos(todos)
}

func dumpTodos(todos *[]models.Todo) {
	stmt, err := DB.Prepare("INSERT INTO todos (user_id, title, completed) VALUES ($1, $2, $3)")
	if err != nil {
		fmt.Print("Failed to prepare")
		return
	}

	defer stmt.Close()

	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("Failed to initiate trnx")
		return
	}

	for _, todo := range *todos {
		tx.Stmt(stmt).Exec(todo.UserId, todo.Title, todo.Completed)
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("Failed to commit tod", err.Error())
		return
	}
}
