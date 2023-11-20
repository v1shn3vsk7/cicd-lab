package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Healthz ...
func (h *Handlers) Healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
