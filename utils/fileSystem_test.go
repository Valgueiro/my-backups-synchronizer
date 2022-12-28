package utils

import (
	"testing"
)

func TestDoesFolderExist(t *testing.T) {
	existingFolder := "/home"
	if DoesFolderExist(existingFolder) {
		return
	}
	t.Fatalf(`/%s Folder should exist!`, existingFolder)
}

func TestFolderDoesNotExist(t *testing.T) {
	existingFolder := "/path/to/unexistent/folder"
	if !DoesFolderExist(existingFolder) {
		return
	}
	t.Fatalf(`/%s Folder should not exist!`, existingFolder)
}
