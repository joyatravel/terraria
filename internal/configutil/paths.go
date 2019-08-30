package configutil

import (
	"fmt"

	"github.com/tugolo/terraria/internal/info"
)

// ConfigDir is the configuration directory corresponding to this module.
const ConfigDir = "/etc/" + info.Namespace

// ConfigPaths returns a slice of config file locations for a given component
// name.
func ConfigPaths(component string) []string {
	return []string{
		component + ".yaml",
		fmt.Sprintf("%s/%s.yaml", ConfigDir, component),
	}
}
