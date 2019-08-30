package http

import (
	"fmt"
	"net/http"

	"github.com/cockroachdb/errors"
	echo "github.com/labstack/echo/v4"
	"github.com/tugolo/terraria/terraria"
)

// WorldHandler creates a echo.HandlerFunc that handles requests to download
// world files.
func WorldHandler(
	worlds terraria.WorldService,
	status terraria.StatusService,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		status, err := status.GetStatus()
		if err != nil {
			return errors.Wrap(err, "getting server status")
		}

		// Find world file.
		file, info, err := worlds.GetWorldFile(status.World)
		if err != nil {
			return err
		}
		defer file.Close()

		// Set appropriate headers.
		c.Response().Header().Set(
			"Content-Disposition",
			fmt.Sprintf(`attachment; filename="%s"`, info.Name()),
		)

		// Serve world file.
		http.ServeContent(
			c.Response(), c.Request(),
			info.Name(), info.ModTime(), file,
		)

		err = file.Close()
		return errors.Wrap(err, "closing world file")
	}
}
