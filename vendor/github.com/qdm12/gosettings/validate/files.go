package validate

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ErrFileDoesNotExist = errors.New("file does not exist")
	ErrFileStat         = errors.New("failed stating file")
	ErrFilePathIsDir    = errors.New("filepath is a directory")
	ErrFilepathIsFile   = errors.New("filepath is a file")
)

// FileExists returns a `nil` error if the given `path` exists
// and is a file. Otherwise, an error is returned, wrapping either
// `ErrFileDoesNotExist`, `ErrFileStat` or `ErrFilePathIsDir`.
func FileExists(path string) (err error) {
	const directory = false
	return fileExists(path, directory)
}

// DirectoryExists returns a `nil` error if the given `path` exists
// and is a directory. Otherwise, an error is returned, wrapping either
// `ErrFileDoesNotExist`, `ErrFileStat` or `ErrFilepathIsFile`.
func DirectoryExists(path string) (err error) {
	const directory = true
	return fileExists(path, directory)
}

func fileExists(path string, directory bool) (err error) {
	path = filepath.Clean(path)

	stat, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("%w: %s", ErrFileDoesNotExist, path)
	} else if err != nil {
		return fmt.Errorf("%w: %w", ErrFileStat, err)
	}

	if directory && !stat.IsDir() {
		return fmt.Errorf("%w: %s", ErrFilepathIsFile, path)
	} else if !directory && stat.IsDir() {
		return fmt.Errorf("%w: %s", ErrFilePathIsDir, path)
	}

	return nil
}
