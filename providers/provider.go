package providers

import "os"

func getArchivePath() string {
	val, ok := os.LookupEnv("ARCHIVE_PATH")
	if !ok {
		return "/archives/" // Default value
	}
	return val
}

var ARCHIVE_PATH string = getArchivePath()

type Provider interface {
	sync() error
}
