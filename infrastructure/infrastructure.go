package infrastructure

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewConnectDB),
	fx.Provide(NewFirebaseApp),
	fx.Provide(NewFirebaseAuth),
	fx.Provide(NewRouter),
)
