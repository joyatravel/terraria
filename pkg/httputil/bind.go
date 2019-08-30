package httputil

import (
	"net/http"

	"github.com/cockroachdb/errors/exthttp"
	echo "github.com/labstack/echo/v4"
	"go.stevenxie.me/gopkg/zero"
)

// Bind calls c.Bind, but associates any errors from it with the HTTP status
// code 422 (unprocessable entity).
func Bind(c echo.Context, v zero.Interface) error {
	if err := c.Bind(v); err != nil {
		return exthttp.WrapWithHTTPCode(err, http.StatusUnprocessableEntity)
	}
	return nil
}
