package rfsnotify

import "fmt"

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
	filePaths []string
}

//Include("path1", "path2", "path3", "...")
func (w *Watcher) Include(path ...string) {
	for _, existingPath := range w.filePaths {
		for _, newPath := range path {
			if existingPath == newPath {
				return
			}
		}
	}
	w.filePaths = append(w.filePaths, path...)
}

//Exlude
func (w *Watcher) Exclude(path ...string) {
	for _, value := range path {
		for i, v := range w.filePaths {
			if value == v {
				w.filePaths = deletePath(w.filePaths, i)
			}
		}
	}
}

func deletePath(paths []string, index int) []string { //this path is not the same paths paste.
	if index > len(paths)-1 {
		panic(fmt.Sprintf("index %v is bigger than the size of the paths slice!", index))
	}
	if index < len(paths)-1 {
		return append(paths[:index], paths[index+1:]...) // or we can move the index to the end and use the deletePath function.
	}

	return paths[:index]
}
