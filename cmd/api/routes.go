package main

import (
	"bougette-backend/cmd/api/handlers"
)

func (app *Application) routes(handler handlers.Handler) {
	apiGroup := app.server.Group("/api")
	publicAuthRoutes := apiGroup.Group("/auth")
	{
		app.server.POST("/register", handler.RegisterHandler)
		publicAuthRoutes.POST("/login", handler.LoginHandler)

	}
	profileRoutes := apiGroup.Group("/profile", app.appMiddleware.AuthenticationMiddleware)
	{
		profileRoutes.GET("/authenticated/user", handler.GetAuthenticatedUser)
		profileRoutes.PATCH("/update/password", handler.UpdateUserPassword)
	}

	publicAuthRoutes.GET("/", handler.HealthCheck)

}
