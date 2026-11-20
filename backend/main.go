package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Health struct {
	Res string `json:"res"`
}

func healthHandler(c echo.Context) error {
	testHealth := Health{"ok"}
	return c.JSON(http.StatusOK, testHealth)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	api.GET("/health", healthHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
