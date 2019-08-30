package terraria

import "os"

// A WorldService is capable of retrieving a world file from disk.
type WorldService interface {
	GetWorldFile(name string) (*os.File, os.FileInfo, error)
}
