package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handlers) Healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
