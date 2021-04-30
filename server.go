package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type responseGetHelloWorld struct {
	Message string `json:"message"`
}

type responseGetHealth struct {
	Status string `json:"status"`
}

func getHealth(c echo.Context) error {
	r := &responseGetHealth{}
	r.Status = "healthy"
	return c.JSON(http.StatusOK, r)
}

func getHelloWorld(c echo.Context) error {
	r := &responseGetHelloWorld{}
	r.Message = "hello world"
	return c.JSON(http.StatusOK, r)
}

func main() {
	e := echo.New()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		c.JSON(report.Code, report)
	}
	e.GET("/health", getHealth)
	e.GET("/hello-world", getHelloWorld)
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":80"))
}
