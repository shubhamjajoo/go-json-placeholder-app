package services

import (
	"encoding/json"
	"go-json-placeholder-app/models"
	"io"
	"net/http"
)

func FetchUsers() (*[]models.User, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		panic("Failed to fetch users!")
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var users []models.User
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic("Failed to parse response body")
		}
		err = json.Unmarshal(bytes, &users)
		if err != nil {
			panic("Failed to unmarshel json")
		}
		return &users, nil
	}
	return nil, err
}
