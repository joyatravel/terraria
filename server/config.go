package server

import (
	"github.com/sirupsen/logrus"
)

// An Option modifies a Config.
type Option func(*Config)

// WithLogger configures a Server to write logs with log.
func WithLogger(log logrus.FieldLogger) func(*Config) {
	return func(cfg *Config) { cfg.Logger = log }
}
