package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MakeHandler(instance *echo.Echo, s *resource) {
	// Echo framework için yeni bir instance oluşturulur
	e := instance

	// Middleware tanımlamaları yapılır
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Endpoint'lerimiz oluşturulur.

	e.GET("/api/v1/user", s.getUser)

}
