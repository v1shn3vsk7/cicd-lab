package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handlers) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
