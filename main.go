package main

import (
	"net/http"

	"github.com/Advik-Gupta/collacode-be/pkg/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := router.NewRouter()

	// Simple route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})



	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}