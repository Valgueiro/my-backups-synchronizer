package providers

import (
	"fmt"
	"my-backups-synchronizer/utils"
	"os/exec"
)

// rclone sync mateus_personal_drive:/ /media/mvalgueiro/Arquivos/archives/google-drive-backup/mateus_personal_drive -P

type GoogleDriveProvider struct {
	rcloneConfigName string
}

func (g *GoogleDriveProvider) generateFolderPath() string {
	return ARCHIVE_PATH + "google-drive-backup/" + g.rcloneConfigName + "/"
}

func (g *GoogleDriveProvider) doesOutputPathExist() bool {
	out := utils.DoesFolderExist(g.generateFolderPath())
	fmt.Println(out)
	return out
}

func (g *GoogleDriveProvider) createOutputPath() bool {
	return true
}

func (g *GoogleDriveProvider) checkOutputPath() {
	fmt.Println("Checking output path...")
	if g.doesOutputPathExist() {
		fmt.Println("Path does not exist. Creating it...")
		g.createOutputPath()
	}
}

func (g *GoogleDriveProvider) generateRcloneSyncCommand() string {
	return fmt.Sprintf("rclone sync %s:/ %s -P", g.rcloneConfigName, g.generateFolderPath())
}

func (g *GoogleDriveProvider) runRcloneSyncCommand() error {
	return exec.Command(g.generateRcloneSyncCommand()).Run()
}

func (g *GoogleDriveProvider) sync() error {
	fmt.Printf("Syncing Google Drive %s\n", g.rcloneConfigName)
	g.checkOutputPath()

	return g.runRcloneSyncCommand()
}
