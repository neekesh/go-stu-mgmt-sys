package routes

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewStudentRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	studentRoutes StudentRoutes,
) Routes {
	return Routes{
		studentRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
