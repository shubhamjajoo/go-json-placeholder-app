package repository

import (
	"fmt"
	"go-json-placeholder-app/db"
	"go-json-placeholder-app/models"
)

func GetAllUsers() (*[]models.User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Print("Failed to get users")
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Name, &user.Username)
		users = append(users, user)
	}
	return &users, nil
}
