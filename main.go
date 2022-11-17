package main

import (
	"learn-go/api"
	database "learn-go/infrastructure"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	server := gin.Default()

	godotenv.Load()

	database.ConnectDB()

	server.GET("/", api.GetStudent)

	server.POST("/create", api.CreateStudent)

	server.DELETE("/delete", api.DeleteStudent)

	server.PUT("/update", api.UpdateStudent)
	server.Run(":2000")
}
