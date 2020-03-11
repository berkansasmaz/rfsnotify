package rfsnotify

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

//TestMethodName_TestInputs_TestOutputs

func TestInclude_AddingPaths_AddedNewPaths(t *testing.T) {
	var watcher = &Watcher{} // equals new(Watcher)
	watcher.Include("test1", "test2")
	if len(watcher.filePaths) != 2 {
		t.Error("len(watcher. filePaths) must be 2.")
	}
	if !watcher.filePaths["test1"] || !watcher.filePaths["test2"] {
		t.Error("watcher.filePaths does not have correct items.")

	}
}

func TestInclude_AddingNothing_ReturnsNil(t *testing.T) {
	var watcher = new(Watcher)
	if watcher.filePaths != nil {
		t.Error("wathcer.filePaths must be nil.")
	}
}

func TestExclude_RemovingExistingItems_ItemsRemoved(t *testing.T) {
	var watcher = new(Watcher)
	watcher.Include("file1.txt", "file2.txt", "file3.txt")
	watcher.Exclude("file2.txt")
	if len(watcher.filePaths) != 2 {
		t.Error("len(watcher.filePaths) must be 2.")
	}
	if !watcher.filePaths["file1.txt"] || !watcher.filePaths["file3.txt"] {
		t.Error("watcher.filePath[0] must be 'file1.txt' and watcher.filePath[1] must be 'file3.txt'.")
	}
}

func TestInclude_RemovingNonExistingItem_SliceRemainedTheSame(t *testing.T) {
	var watcher = new(Watcher)
	watcher.Include("file1.txt", "file2.txt")
	watcher.Exclude("file3.txt")
	if len(watcher.filePaths) != 2 {
		t.Error("len(watcher.filePaths must be 2)")
	}
	if !watcher.filePaths["file1.txt"] && watcher.filePaths["file2.txt"] {
		t.Error("watcher.filePath[0] must be 'file1.txt' and watcher.filePath[1] must be 'file3.txt.'")
	}
}

func TestInclude_AddingDuplicateItem_DuplicateItemsNotAdded(t *testing.T) {
	var watcher = new(Watcher)
	watcher.Include("file1.txt")
	watcher.Include("file1.txt")
	watcher.Include("file2.txt")
	if len(watcher.filePaths) != 2 {
		t.Error("len(watcher.filePath) must be 2.")
	}
}

func TestInclude_AddingDuplicateItemAtTheSameTime_DuplicateItemsNotAdded(t *testing.T) {
	var watcher = new(Watcher)
	watcher.Include("file1.txt", "file1.txt", "file2.txt")
	if len(watcher.filePaths) != 2 {
		t.Error("len(watcher.filePath) must be 2.")
	}
}

func TestNewWatcher_GivenDirectory_ReturnsAllFiles(t *testing.T) {
	//Setup
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		t.Fatal("Cannot create a temp directory")
	}
	defer os.RemoveAll(dir) // clean up

	err = os.MkdirAll(path.Join(dir, "dir1", "dir2", "dir3"), os.ModePerm)

	if err != nil {
		t.Fatal("Cannot create a temp directory")
	}

	var tempFiles = []string{path.Join(dir, "dir1", "file1"),
		path.Join(dir, "dir1", "file2"),
		path.Join(dir, "dir1", "dir2", "file3"),
		path.Join(dir, "dir1", "dir2", "file4"),
		path.Join(dir, "dir1", "dir2", "dir3", "file5"),
		path.Join(dir, "dir1", "dir2", "dir3", "file6"),
	}

	for _, fileName := range tempFiles {
		err := ioutil.WriteFile(fileName, []byte("hello world"), os.ModePerm)
		if err != nil {
			t.Error("Cannot create file" + fileName)
		}
	}

	//Test
	watcher, err := NewWatcher(dir, true, nil)
	if err != nil {
		t.Error("An unidentified error", err)
	}
	if len(watcher.filePaths) != 6 {
		t.Error("watcher didn't find all the files")
	}

}
