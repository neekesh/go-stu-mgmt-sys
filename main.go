package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"learn-go/bootstrap"
)

func main() {
	// infrastructure.InitializeLogger()
	// infrastructure.Logger.Info("started Server")
	godotenv.Load()
	fx.New(bootstrap.Module).Run()

}
