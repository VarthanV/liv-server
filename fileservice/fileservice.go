package fileservice

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
)

func (f *Service) InitWatcher(ctx context.Context) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logrus.Error("error in creating watcher ", err)
		return
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case <-ctx.Done():
				logrus.Info("context done")
				return
			case event, ok := <-watcher.Events:
				if !ok {
					logrus.Info("watcher events channel closed")
					return
				}
				logrus.Info("event: ", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					logrus.Info("modified file: ", event.Name)
				}
			}

		}
	}()
}

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
