package utils

import (
	"os"
	"path/filepath"
)

func FullPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	p, _ := os.Executable()
	return filepath.Join(filepath.Dir(p), path)
}
