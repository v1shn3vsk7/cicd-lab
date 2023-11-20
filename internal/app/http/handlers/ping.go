package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Ping ...
func (h *Handlers) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
