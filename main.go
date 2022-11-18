package main

import (
	"learn-go/api"
	infrastructure "learn-go/infrastructure"
	"learn-go/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	server := gin.Default()
	// infrastructure.InitializeLogger()
	// infrastructure.Logger.Info("started Server")
	godotenv.Load()

	infrastructure.ConnectDB()

	infrastructure.SetupFirebase()

	server.Use(middleware.FirebaseAuth)

	server.GET("/", api.GetAllStudent)

	server.POST("/create", api.PostStudent)

	server.DELETE("/delete/:id", api.DeleteStudent)

	server.PUT("/update/:id", api.UpdateStudent)

	server.Run(":2000")
}
