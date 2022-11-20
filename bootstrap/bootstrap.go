package bootstrap

import (
	"learn-go/api"
	infrastructure "learn-go/infrastructure"
	"learn-go/middleware"
	"os"

	"github.com/gin-gonic/gin"

	"go.uber.org/fx"
)

var Module = fx.Options(
	infrastructure.Module,
	fx.Invoke(bootstrap),
)

func bootstrap() {
	router := gin.Default()
	db := infrastructure.ConnectDB()

	firebaseAuth := infrastructure.SetupFirebase()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("firebaseAuth", firebaseAuth)
	})

	router.Use(middleware.FirebaseAuth)

	router.GET("/", api.GetAllStudent)

	router.POST("/create", api.PostStudent)

	router.DELETE("/delete/:id", api.DeleteStudent)

	router.PUT("/update/:id", api.UpdateStudent)

	router.Run(":" + os.Getenv("ServerPort"))
}
