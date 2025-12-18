package routes

import (
	"go-json-placeholder-app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTodos(c *gin.Context) {
	todos, err := repository.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to get todos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todos fetched!", "todos": todos})
}
