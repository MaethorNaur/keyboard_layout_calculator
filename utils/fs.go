package utils

import (
	"errors"
	"io/fs"
	"os"
)

func EnsureDirExists(dir string) error {
	if err := os.MkdirAll(dir, os.ModeDir); err != nil && !errors.Is(err, fs.ErrExist) {
		return (err)
	}
	return nil
}
