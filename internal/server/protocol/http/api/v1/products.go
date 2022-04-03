package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Products(c echo.Context) error {
	return c.JSON(http.StatusForbidden, "dddd")
}
