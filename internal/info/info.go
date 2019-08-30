package info

// Version is the current module version, and is set at compile-time
// with the following linker flag:
//   -X github.com/tugolo/terraria/internal/info.Version=$(VERSION)
var Version = "unset"

// Namespace is the module name, used for things like envvar prefixes.
const Namespace = "terraria"
