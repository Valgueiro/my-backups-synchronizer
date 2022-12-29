package providers

import (
	"fmt"
	"my-backups-synchronizer/utils"
	"os"
	"os/exec"
)

var GOOGLE_DRIVE_ROOT_PATH string = "google-drive-backup/"

type GoogleDriveProvider struct {
	rcloneConfigName string
}

func (g *GoogleDriveProvider) generateFolderPath() string {
	return ARCHIVE_PATH + GOOGLE_DRIVE_ROOT_PATH + g.rcloneConfigName + "/"
}

func (g *GoogleDriveProvider) doesOutputPathExist() bool {
	return utils.DoesFolderExist(g.generateFolderPath())
}

func (g *GoogleDriveProvider) createOutputPath() error {
	return os.Mkdir(g.generateFolderPath(), 0750)
}

func (g *GoogleDriveProvider) checkOutputPath() error {
	fmt.Println("Checking output path...")
	if g.doesOutputPathExist() {
		fmt.Println("Path does not exist. Creating it...")

		return g.createOutputPath()
	}

	return nil
}

func (g *GoogleDriveProvider) generateRcloneSyncCommand() string {
	return fmt.Sprintf("rclone sync %s:/ %s -P", g.rcloneConfigName, g.generateFolderPath())
}

func (g *GoogleDriveProvider) runRcloneSyncCommand() error {
	return exec.Command(g.generateRcloneSyncCommand()).Run()
}

func (g *GoogleDriveProvider) sync() error {
	fmt.Printf("Syncing Google Drive %s\n", g.rcloneConfigName)
	if err := g.checkOutputPath(); err != nil {
		fmt.Println("Could not access output path.")
		return err
	}

	return g.runRcloneSyncCommand()
}
