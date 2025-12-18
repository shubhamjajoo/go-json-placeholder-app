package services

import (
	"encoding/json"
	"fmt"
	"go-json-placeholder-app/models"
	"io"
	"net/http"
)

func FetchTodos() (*[]models.Todo, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		panic("Failed to fetch todos!")
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var todos []models.Todo
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic("Failed to parse response body")
		}
		err = json.Unmarshal(bytes, &todos)
		if err != nil {
			fmt.Println("Err", err.Error())
			panic("Failed to unmarshel todos json")
		}
		return &todos, nil
	}
	return nil, err
}
