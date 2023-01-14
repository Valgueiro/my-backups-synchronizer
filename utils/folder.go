package utils

import (
	"fmt"
	"os"
)

func DoesFolderExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

type Folder struct {
	Path       string
	Permission os.FileMode
}

func (f *Folder) DoesExist() bool {
	_, err := os.Stat(f.Path)
	return !os.IsNotExist(err)
}

func (f *Folder) CreateOnFileSystem() error {
	return os.MkdirAll(f.Path, f.Permission)
}

func (f *Folder) CheckAndCreateOnFileSystem() error {
	fmt.Println("Checking folder...")
	if f.DoesExist() {
		fmt.Println("Folder Already Exists!")
		return nil
	}

	fmt.Printf("Path does not exist %s. Creating it...\n", f.Path)

	err := f.CreateOnFileSystem()
	if err != nil {
		return err
	}
	fmt.Println("Path Created!")
	return err
}
