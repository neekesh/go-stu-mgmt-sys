package bootstrap

import (
	"learn-go/api/controllers"
	"learn-go/api/routes"
	infrastructure "learn-go/infrastructure"
	"os"

	"go.uber.org/fx"
)

var Module = fx.Options(
	controllers.Module,
	infrastructure.Module,
	routes.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	handler infrastructure.Router,
	routes routes.Routes,
) {
	routes.Setup()
	handler.Gin.Run(":" + os.Getenv("ServerPort"))

}
