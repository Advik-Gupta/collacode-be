package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Simple route
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "Pong")
	})

	return e
}