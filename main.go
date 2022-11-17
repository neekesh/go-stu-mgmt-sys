package main

import (
	"learn-go/api"
	infrastructure "learn-go/infrastructure"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	server := gin.Default()
	// infrastructure.InitializeLogger()
	// infrastructure.Logger.Info("started Server")
	godotenv.Load()

	infrastructure.ConnectDB()

	server.GET("/", api.GetAllStudent)

	server.POST("/create", api.PostStudent)

	server.DELETE("/delete", api.DeleteStudent)

	server.PUT("/update", api.UpdateStudent)
	server.Run(":2000")
}
