package providers

import (
	"fmt"
	utils "my-backups-synchronizer/utils"
	"os/exec"
	"strings"
)

var GOOGLE_DRIVE_ROOT_PATH string = "google-drive-backup/"

type GoogleDriveProvider struct {
	rcloneConfigName string
}

func (g *GoogleDriveProvider) generateOutputPath() string {
	return ARCHIVE_PATH + GOOGLE_DRIVE_ROOT_PATH + g.rcloneConfigName + "/"
}

func (g *GoogleDriveProvider) generateRcloneSyncCommand() (string, []string) {
	cmd := fmt.Sprintf("sync %s:/ %s -P", g.rcloneConfigName, g.generateOutputPath())

	return "rclone", strings.Split(cmd, " ")
}

func (g *GoogleDriveProvider) runRcloneSyncCommand() (output string, err error) {
	executable, args := g.generateRcloneSyncCommand()
	out, err := exec.Command(executable, args...).CombinedOutput()
	return string(out), err
}

func (g *GoogleDriveProvider) sync() error {
	fmt.Printf("Syncing Google Drive %s into %s\n", g.rcloneConfigName, g.generateOutputPath())

	outputPath := &utils.Folder{
		Path:       g.generateOutputPath(),
		Permission: 0750,
	}

	if err := outputPath.CheckAndCreateOnFileSystem(); err != nil {
		fmt.Println("Could not access or create output path.")
		return err
	}

	out, err := g.runRcloneSyncCommand()
	fmt.Println(out)
	return err
}
