package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HeartbeatRoute helps confirm whether or not the app is running
func HeartbeatRoute(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
