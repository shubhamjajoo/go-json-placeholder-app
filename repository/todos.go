package repository

import (
	"fmt"
	"go-json-placeholder-app/db"
	"go-json-placeholder-app/models"
)

func GetAllTodos() (*[]models.Todo, error) {
	query := "SELECT * FROM todos"
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Print("Failed to get todos")
		return nil, err
	}
	defer rows.Close()
	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		rows.Scan(&todo.Id, &todo.Title, &todo.UserId, &todo.Completed)
		todos = append(todos, todo)
	}
	return &todos, nil
}
