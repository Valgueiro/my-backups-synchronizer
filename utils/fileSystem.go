package utils

import (
	"os"
)

func DoesFolderExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
