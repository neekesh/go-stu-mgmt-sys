package bootstrap

import (
	"learn-go/api"
	infrastructure "learn-go/infrastructure"
	"learn-go/middleware"
	"os"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"

	"go.uber.org/fx"
)

var Module = fx.Options(
	infrastructure.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	fbauth *auth.Client,
	db infrastructure.Database,
) {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db.DB)
		c.Set("firebaseAuth", fbauth)
	})

	router.Use(middleware.FirebaseAuth)

	router.GET("/", api.GetAllStudent)

	router.POST("/create", api.PostStudent)

	router.DELETE("/delete/:id", api.DeleteStudent)

	router.PUT("/update/:id", api.UpdateStudent)

	router.Run(":" + os.Getenv("ServerPort"))
}
