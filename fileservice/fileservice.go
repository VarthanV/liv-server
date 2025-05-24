package fileservice

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func (f *Service) InitWatcher(ctx context.Context, conn *websocket.Conn) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logrus.Error("error in creating watcher ", err)
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		logrus.Error("error in getting pwd ", err)
		return
	}
	f.addRecursive(watcher, pwd)

	go func() {
		logrus.Info("Initialized watcher")
		for {
			select {
			case <-ctx.Done():
				logrus.Info("context done")
				watcher.Close()
				return
			case event, ok := <-watcher.Events:
				if !ok {
					logrus.Info("watcher closed")
					return
				}
				logrus.Info("event: ", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					logrus.Info("modified file: ", event.Name)
					conn.WriteMessage(websocket.TextMessage, []byte("reload"))
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

func (f *Service) GetHTML(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		logrus.Error("error in reading file ", err)
		return "", err
	}

	html := string(data)
	customScript := `<script>const socket=new WebSocket("ws://127.0.0.1:8070/ws");socket.onopen=function(){console.log("WebSocket connected"),socket.send("connected")},socket.onmessage=function(o){"reload"===o.data&&window.location.reload()};</script>`
	if strings.Contains(html, "</body>") {
		html = strings.Replace(html, "</body>", customScript+"</body>", 1)
	} else {
		html += customScript
	}
	return html, nil
}

func (f *Service) addRecursive(watcher *fsnotify.Watcher, root string) error {
	return filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			logrus.Info("adding directory: ", path)
			return watcher.Add(path)
		}
		return nil
	})
}
