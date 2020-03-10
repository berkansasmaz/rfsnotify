package rfsnotify

import (
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
	if len(watcher.filePaths) != 1 {
		t.Error("len(watcher.filePath) must be 1.")
	}
}
