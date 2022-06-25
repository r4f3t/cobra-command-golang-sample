package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/r4f3t/webapi/internal/user"
)

type resource struct {
	service user.Service
}

func NewController(service user.Service) *resource {
	return &resource{
		service: service,
	}
}

func (receiver *resource) getUser(c echo.Context) error {
	userId := receiver.service.GetUser()
	return c.JSON(http.StatusOK, userId)
}
