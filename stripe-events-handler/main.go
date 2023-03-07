package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	initLogger()

	config := buildConfig()

	// db := initDB(config)
	e := initServer()

	e.GET("/helloworld", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	e.Logger.Fatal(e.Start(config.HTTP.Port))
}
