package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func AnotherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("we are in the another middlewares")
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}
