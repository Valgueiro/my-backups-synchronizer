package providers

import (
	"fmt"
	"my-backups-synchronizer/utils"
	"os"
	"os/exec"
	"strings"
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
		return nil
	}

	fmt.Printf("Path does not exist %s. Creating it...\n", g.generateFolderPath())

	return g.createOutputPath()

}

func (g *GoogleDriveProvider) generateRcloneSyncCommand() (string, []string) {
	cmd := fmt.Sprintf("sync %s:/ %s -P", g.rcloneConfigName, g.generateFolderPath())

	return "rclone", strings.Split(cmd, " ")
}

func (g *GoogleDriveProvider) runRcloneSyncCommand() (output string, err error) {
	executable, args := g.generateRcloneSyncCommand()
	out, err := exec.Command(executable, args...).CombinedOutput()
	return string(out), err
}

func (g *GoogleDriveProvider) sync() error {
	fmt.Printf("Syncing Google Drive %s into %s\n", g.rcloneConfigName, g.generateFolderPath())
	if err := g.checkOutputPath(); err != nil {
		fmt.Println("Could not access output path.")
		return err
	}

	out, err := g.runRcloneSyncCommand()
	fmt.Println(out)
	return err
}
