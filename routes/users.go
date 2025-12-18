package routes

import (
	"go-json-placeholder-app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	users, err := repository.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to get users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Users fetched!", "users": users})
}
