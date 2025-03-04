package main

import (
	"bougette-backend/cmd/api/handlers"
)

func (app *Application) routes(handler handlers.Handler) {
	app.server.GET("/", handler.HealthCheck)
	app.server.POST("/register", handler.RegisterHandler)
	app.server.POST("/login", handler.LoginHandler)
}
