package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type resource struct {
	// servis katmanını burda
}

func init() {
	// Echo framework için yeni bir instance oluşturulur
	e := echo.New()

	// Middleware tanımlamaları yapılır
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Endpoint'lerimiz oluşturulur.
	// e.GET("/", services.Hello)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// 3200 portundan API'ı ayağa kaldıralım
	e.Logger.Fatal(e.Start(":3200"))
}
