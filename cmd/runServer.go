package cmd

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunServer() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Status  int
			Message string
		}{Status: http.StatusOK, Message: "Hello caldwellcsc :)"})
	})

	e.GET("/hey", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Status  int
			Message string
		}{Status: http.StatusOK, Message: "Hey Hey Hey Hey"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
