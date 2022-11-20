package infrastructure

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(ConnectDB),
	fx.Provide(SetupFirebase),
)
