package cmdutil

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger creates an application-level logger, which logs at debug level
// when GOENV == "development".
func Logger() *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)

	// Set logger level.
	if os.Getenv("GOENV") == "development" {
		log.SetLevel(logrus.DebugLevel)
	}
	return log
}
