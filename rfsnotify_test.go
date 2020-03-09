package rfsnotify

import "testing"

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
	var watcher = &Watcher{}
	watcher.Include("test1", "test2")
	if len(watcher.filePaths) == 2 {
		t.Error("len(watcher.filePaths) must be 2.")
	}
	if watcher.filePaths[0] == "test1" && watcher.filePaths[1] == "test2" {
		t.Error("watcher.filePaths does not have correct items.")

	}
}
