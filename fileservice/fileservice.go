package fileservice

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

// List lists all the files and subdirectories in the given path
func (f *Service) List(rootPath string, requestPath string) ([]FileItem, error) {

	files := []FileItem{}
	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			logrus.Error("error in walking directory ", err)
			return err
		}

		if rootPath == path {
			logrus.Info("skipping root")
			return nil
		}

		// Get the relative depth
		relPath, err := filepath.Rel(rootPath, path)
		if err != nil {
			return err
		}
		if strings.Count(relPath, string(os.PathSeparator)) >= 1 {
			if d.IsDir() {
				return fs.SkipDir // prevent entering this directory
			}
			return nil // skip file too (not top level)
		}

		files = append(files, FileItem{
			Path:    filepath.Base(path),
			RootDir: requestPath,
			IsDir:   d.IsDir(),
		})
		return nil

	})

	return files, err
}
