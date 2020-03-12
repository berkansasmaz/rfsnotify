package rfsnotify

import (
	"os"
	"path/filepath"
)

//rfsnotify.Add("directory", recursive=true, file.write, file.create, file.rename)
//blank identifier _
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

// Creates a new Watcher objcet and initializes the internal watch list
// based on the given path.
func NewWatcher(path string, recusive bool, event []Event) *Watcher {
	var watcher = &Watcher{
		Path:      path,
		Recursive: recusive,
		Events:    event,
	}

	initFilePath(watcher)

	return watcher
}

func initFilePath(w *Watcher) {
	givenFileInfo, err := os.Stat(w.Path)
	if err != nil {
		panic(err)
	}

	var allFilePaths []string

	switch mode := givenFileInfo.Mode(); {
	case mode.IsDir():
		allFilePaths = getAllFiles(w.Path)
		w.Include(allFilePaths...)
	case mode.IsRegular():
		w.Include(w.Path)
	}

}

// Finds newly added files in given path.
func (w *Watcher) Refresh() {
	initFilePath(w)
}

//walking directory
func getAllFiles(dirPath string) []string {
	var files []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		//todo check this Logic later.
		if err != nil {
			panic(err)
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return err
	})
	if err != nil {
		panic(err)
	}

	return files
}

// Add new files to the internal watch list to track.
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

// Exludes paths from internal watch list .
func (w *Watcher) Exclude(paths ...string) {
	for _, path := range paths {
		delete(w.filePaths, path)
	}
}
