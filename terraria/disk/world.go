package disk

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"go.stevenxie.me/gopkg/zero"

	"github.com/cockroachdb/errors"
	"github.com/sirupsen/logrus"
	"github.com/tugolo/terraria/terraria"
)

// NewWorldFinder creates a new WorldFinder that finds worlds located at dir.
// It implements a terraria.WorldService.
func NewWorldFinder(dir string, opts ...func(*WorldFinderConfig)) WorldFinder {
	cfg := WorldFinderConfig{
		Logger: zero.Logger(),
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return WorldFinder{
		dir: dir,
		log: cfg.Logger,
	}
}

type (
	// A WorldFinder implements a terraria.WorldService that looks for world files
	// in a particular directory on a disk.
	WorldFinder struct {
		dir string
		log logrus.FieldLogger
	}

	// A WorldFinderConfig configures a WorldFinder.
	WorldFinderConfig struct {
		Logger logrus.FieldLogger
	}
)

var _ terraria.WorldService = (*WorldFinder)(nil)

// GetWorldFile gets the Terraria world file with the corresponding name from
// the worlds directory on disk.
func (wf WorldFinder) GetWorldFile(name string) (*os.File, os.FileInfo, error) {
	var (
		log        = wf.log.WithField("world", name)
		files, err = ioutil.ReadDir(wf.dir)
	)
	if err != nil {
		log.WithError(err).Error("Error while reading world directory.")
		return nil, nil, errors.Wrap(err, "disk: reading world directory")
	}

	for _, f := range files {
		n := f.Name()
		world := strings.TrimSuffix(n, path.Ext(n))
		if world == name {
			file, err := os.Open(filepath.Join(wf.dir, n))
			if err != nil {
				log.WithField("filename", n).Error("Failed to open file.")
				return nil, nil, err
			}
			return file, f, nil
		}
	}

	return nil, nil, ErrWorldNotFound
}

// ErrWorldNotFound is returned by a WorldFinder when it fails to find the file
// corresponding to a particular world.
var ErrWorldNotFound = errors.New("disk: no such world file was found")
