package config

import (
	"time"

	"github.com/cockroachdb/errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	// Config is used to configure command 'server'.
	Config struct {
		Server struct {
			ShutdownTimeout *time.Duration `yaml:"shutdownTimeout"`
		} `yaml:"server"`

		Terraria struct {
			Address  string `yaml:"serverAddress"`
			WorldDir string `yaml:"worldDir"`
		} `yaml:"terraria"`
	}
)

// Validate returns an error if the Config is not valid.
func (cfg *Config) Validate() error {
	{
		server := &cfg.Server
		if err := validation.ValidateStruct(
			server,
			validation.Field(&server.ShutdownTimeout, validation.Min(0)),
		); err != nil {
			return errors.Wrap(err, "config: validating Server")
		}
	}
	return nil
}

func defaultConfig() *Config {
	cfg := new(Config)
	return cfg
}
