package rfsnotify

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
)

//rfsnotify.Add("directory", recursive=true, file.write, file.create, file.rename)
//blank identifier _

//type Watcher
type Watcher struct {
	Path   string
	Events chan fsnotify.Event // Enum array
	Errors chan error
	// todo think about the usefulness of this backing slice. Because we can use fsnotify.Watcher internal backing slice.
	filePaths       map[string]bool
	internalWatcher *fsnotify.Watcher
}

// Creates a new Watcher objcet and initializes the internal watch list
// based on the given path.
func NewWatcher(path string) *Watcher {
	var watcher = &Watcher{
		Path: path,
	}

	initFilePath(watcher)

	var fsWatcher, err = fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
	}

	setInternalWatcher(watcher)

	for path := range watcher.filePaths {
		// todo handle error.
		fsWatcher.Add(path)
	}

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

	if w.internalWatcher == nil {
		setInternalWatcher(w)
	}

	for _, path := range paths {
		if !w.filePaths[path] {
			w.filePaths[path] = true
			// todo handle error
			w.internalWatcher.Add(path)
		}
	}
}

// Exludes paths from internal watch list .
func (w *Watcher) Exclude(paths ...string) {
	for _, path := range paths {
		delete(w.filePaths, path)
		// todo handle error here.
		w.internalWatcher.Remove(path)
	}
}

func setInternalWatcher(w *Watcher) {
	w.internalWatcher, _ = fsnotify.NewWatcher()
	w.Events = w.internalWatcher.Events
	w.Errors = w.internalWatcher.Errors
}
