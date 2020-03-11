package rfsnotify

import (
	"fmt"
	"os"
	"path/filepath"
)

//rfsnotify.Add("directory", recursive=true, file.write, file.create, file.rename)
//blank idendtifier _
//GoLang Enum
type Event int

const (
	Unknown Event = iota // number of the current const specification in a (usually parenthesized)
	Deleted
	Created
	Renamed
	Write
)

//type Watcher
type Watcher struct {
	Path      string
	Recursive bool
	Events    []Event //Enum array
	filePaths map[string]bool
}

//Wathcer Constructor
func NewWatcher(path string, recusive bool, event []Event) (*Watcher, error) {
	var watcher = &Watcher{
		Path:      path,
		Recursive: recusive,
		Events:    event,
	}

	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	var allFilePaths []string

	switch mode := fi.Mode(); {
	case mode.IsDir():
		allFilePaths = walkDir(path)
		watcher.Include(allFilePaths...)
	case mode.IsRegular():
		watcher.Include(path)
	}

	return watcher, nil
}

//walking directory
func walkDir(dirPath string) []string {
	var files []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		//todo check this Logic later.
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return err
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dirPath, err)
		return nil
	}

	return files
}

//Include("path1", "path2", "path3", "...")
func (w *Watcher) Include(paths ...string) {
	if w.filePaths == nil {
		w.filePaths = make(map[string]bool)
	}
	for _, path := range paths {
		if !w.filePaths[path] {
			w.filePaths[path] = true
		}
	}
}

//Exlude
func (w *Watcher) Exclude(paths ...string) {
	for _, path := range paths {
		delete(w.filePaths, path)
	}
}
