package config

import (
	"github.com/tugolo/terraria/cmd/server/internal/info"
	"github.com/tugolo/terraria/internal/configutil"
)

// Load finds and loads a Config from a set of possible file locations.
func Load(filenames ...string) (*Config, error) {
	cfg := defaultConfig()
	if err := configutil.TryLoadConfig(cfg, info.Name, filenames...); err != nil {
		return nil, err
	}
	return cfg, nil
}
