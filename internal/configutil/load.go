package configutil

import (
	"go.stevenxie.me/gopkg/configutil"
	"go.stevenxie.me/gopkg/zero"
)

// TryLoadConfig finds and loads a YAML config file from a set of possible
// filenames.
//
// If no filenames are provided, it will derive a set of default locations from
// the component name.
func TryLoadConfig(
	cfg zero.Interface,
	component string,
	filenames ...string,
) error {
	if len(filenames) == 0 {
		filenames = ConfigPaths(component)
	}
	return configutil.TryLoadConfig(cfg, filenames...)
}
