package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/sshanzel/update-right/domain/entities"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":2020"))
}

func hello(c echo.Context) error {
	var p = entities.Publisher{Title: "Test"}

	return c.String(http.StatusOK, p.Title)
}
