package rfsnotify

import (
	"testing"
)

//TestMethodName_TestInputs_TestOutputs

func getSamplesPaths() []string {
	return []string{"hello", "world", "mars"}
}

func TestDeletePath_IndexSmallerThanLen_DeletePathAtIndex(t *testing.T) {
	var index = 1
	var paths = getSamplesPaths()
	var result = deletePath(paths, index)

	if len(result) != 2 {
		t.Fatal("len(result) should not be bigger than len(paths).")
	}
}
func TestDeletePath_IndexGreaterThanLen_Panics(t *testing.T) {
	var index = 3
	var paths = getSamplesPaths()
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Test must have paniced!")
		}
	}()
	_ = deletePath(paths, index)
}

func TestDeletePath_IndexIsLenMinusOne_DeletesLastElement(t *testing.T) {
	var index = 2
	var paths = getSamplesPaths()
	var result = deletePath(paths, index)

	if len(result) != 2 {
		t.Log("Len(result) should have been two.")
		t.Fail()
	}
	if result[1] != "world" {
		t.Log("result[1] must have been 'world'.")
		t.Fail()
	}
}

func TestInclude_AddingPaths_AddedNewPaths(t *testing.T) {
	var watcher = &Watcher{} // equals new(Watcher)
	watcher.Include("test1", "test2")
	if len(watcher.filePaths) != 2 {
		t.Error("len(watcher.filePaths) must be 2.")
	}
	if watcher.filePaths[0] != "test1" && watcher.filePaths[1] != "test2" {
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
		t.Error("len(watcher.filePaths must be 2)")
	}
	if watcher.filePaths[0] != "file1.txt" && watcher.filePaths[1] != "file3.txt" {
		t.Error("watcher.filePath[0] must be 'file1.txt' and watcher.filePath[1] must be 'file3.txt'")
	}
}

func TestInclude_RemovingNonExistingItem_SliceRemainedTheSame(t *testing.T) {
	var watcher = new(Watcher)
	watcher.Include("file1.txt", "file2.txt")
	watcher.Exclude("file3.txt")
	if len(watcher.filePaths) != 2 {
		t.Error("len(watcher.filePaths must be 2)")
	}
	if watcher.filePaths[0] != "file1.txt" && watcher.filePaths[1] != "file2.txt" {
		t.Error("watcher.filePath[0] must be 'file1.txt' and watcher.filePath[1] must be 'file3.txt'")
	}
}
