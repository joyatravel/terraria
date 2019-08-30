package httputil

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"go.stevenxie.me/gopkg/zero"
)

// JSONPretty sends v as JSON, but with nice indentations.
func JSONPretty(c echo.Context, v zero.Interface) error {
	return c.JSONPretty(http.StatusOK, v, "  ")
}
