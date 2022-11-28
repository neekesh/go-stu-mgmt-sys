package middleware

import "go.uber.org/fx"

var Module = fx.Options(
	// fx.Provide(NewMiddewares),
	fx.Provide(NewFirebaseAuth),
)

// type Middleware interface {
// 	Handle()
// }

// type Middlewares []Middleware

// func NewMiddewares() Middlewares {
// 	return Middlewares{}
// }

// func (m Middlewares) Handle() {
// 	for _, middleware := range m {
// 		middleware.Handle()
// 	}
// }
