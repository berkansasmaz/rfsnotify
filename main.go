package rfsnotify

//rfsnotify.Add("directory", recursive=true, file.write, file.create, file.rename)

//GoLang Enum
type Event int

const (
	Unknown Event = 0
	Deleted Event = 1
	Created Event = 2
	Renamed Event = 3
	Write   Event = 4
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
	w.filePaths = append(w.filePaths, path...)
}

//Exlude
func (w *Watcher) Exclude(path ...string) {
	var indices = []int{}
	for _, value := range path {
		for i, v := range w.filePaths {
			if value == v {
				indices = append(indices, i)
			}
		}
	}

	for i := range indices {
		delete(w.filePaths, indices[i])
	}
}

func delete(paths []string, index int) {
	paths = append(paths[:index], paths[index+1:]...) // or we can move the index to the end and use the delete function.
}
