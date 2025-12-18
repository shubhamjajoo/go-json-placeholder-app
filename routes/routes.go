package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(server *gin.Engine) {
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Sever up and running"})
	})
	server.GET("/users", getUsers)
	server.GET("/todos", getTodos)
}
