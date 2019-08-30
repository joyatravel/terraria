package http

import (
	echo "github.com/labstack/echo/v4"
	"github.com/tugolo/terraria/pkg/httputil"
	"github.com/tugolo/terraria/terraria"
)

// StatusHandler creates an echo.HandlerFunc that handles requests for Terraria
// game status information.
func StatusHandler(svc terraria.StatusService) echo.HandlerFunc {
	return func(c echo.Context) error {
		status, err := svc.GetStatus()
		if err != nil {
			return err
		}
		return httputil.JSONPretty(c, status)
	}
}
