package rfsnotify

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
