package main

import (
	"bougette-backend/cmd/api/handlers"
	"bougette-backend/cmd/api/middlewares"
	"bougette-backend/common"
	"bougette-backend/internal/mailer"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

type Application struct {
	logger        echo.Logger
	server        *echo.Echo
	handler       handlers.Handler
	appMiddleware middlewares.AppMiddlewares
}

func main() {
	e := echo.New()
	err := godotenv.Load()
	db, err := common.NewMysql()
	if err != nil {
		e.Logger.Fatal(err.Error())
	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appMailer := mailer.NewMailer(e.Logger)
	h := handlers.Handler{
		DB:     db,
		Logger: e.Logger,
		Mailer: appMailer,
	}
	appMiddleware := middlewares.AppMiddlewares{
		DB:     db,
		Logger: e.Logger,
	}

	app := Application{
		logger:        e.Logger,
		server:        e,
		handler:       h,
		appMiddleware: appMiddleware,
	}

	//e.Use(middleware.Logger())
	e.Use(middleware.Logger(), middlewares.AnotherMiddleware, middlewares.CustomMiddleware)
	app.routes(h)
	fmt.Println(app)
	port := os.Getenv("APP_PORT")
	appAddress := fmt.Sprintf("localhost:%s", port)
	e.Logger.Fatal(e.Start(appAddress))
}
