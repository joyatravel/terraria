package httputil

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/cockroachdb/errors/exthttp"
	echo "github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// ErrorHandler creates an echo.HTTPErrorHandler that handles errors by
// sending JSON responses containing error details.
//
// It will attempt to extract additional error context using cockroachdb/errors.
func ErrorHandler(log logrus.FieldLogger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		var data struct {
			Error   string   `json:"error"`
			Cause   string   `json:"cause,omitempty"`
			Details []string `json:"details,omitempty"`
		}

		// Check if error contains a status code.
		status := exthttp.GetHTTPCode(err, http.StatusInternalServerError)

		// Check if error comes from Echo.
		{
			var httpError *echo.HTTPError
			if errors.As(err, &httpError) {
				status = httpError.Code
				data.Error = strings.ToLower(fmt.Sprint(httpError.Message))
				goto Send
			}
		}

		// Build error response.
		data.Error = err.Error()
		if cause := errors.UnwrapAll(err); cause != err {
			data.Cause = cause.Error()
		}
		if details := errors.GetAllDetails(err); len(details) > 0 {
			data.Details = details
		}

	Send:
		// Send error, handle JSON marshalling failures.
		if err = c.JSON(status, &data); err != nil {
			const msg = "Failed to write JSON error."
			c.Response().WriteHeader(http.StatusInternalServerError)
			io.WriteString(c.Response(), msg)
			log.WithError(err).Error(msg)
		}
	}
}
