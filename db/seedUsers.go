package db

import (
	"fmt"
	"go-json-placeholder-app/models"
	"go-json-placeholder-app/services"
)

func seedUsers() {
	users, err := services.FetchUsers()
	if err != nil {
		fmt.Println("Failed to fetch users")
		return
	}
	dumpUsers(users)
}

func dumpUsers(users *[]models.User) {
	// stmt, err := DB.Prepare("INSERT INTO users (name, username) VALUES (?, ?)")
	stmt, err := DB.Prepare("INSERT INTO users (name, username) VALUES ($1, $2)")
	if err != nil {
		fmt.Print("Failed to prepare", err.Error())
		return
	}

	defer stmt.Close()

	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("Failed to initiate trnx")
		return
	}

	for _, user := range *users {
		tx.Stmt(stmt).Exec(user.Name, user.Username)
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("Failed to commit ", err.Error())
		return
	}
}
