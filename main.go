package main

import (
	"go-json-placeholder-app/db"
	"go-json-placeholder-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDBB()
	server := gin.Default()
	routes.InitRoutes(server)
	server.Run()
}
