package providers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateInstance(rcloneConfigName string) *GoogleDriveProvider {
	return &GoogleDriveProvider{
		rcloneConfigName: rcloneConfigName,
	}
}

func TestGenerateFolderPath(t *testing.T) {
	rcloneName := "test"
	expected := fmt.Sprintf("%s%s%s/", ARCHIVE_PATH, GOOGLE_DRIVE_ROOT_PATH, rcloneName)
	actual := generateInstance(rcloneName).generateFolderPath()

	assert.Equal(t, expected, actual)
}

func TestDoesOutputPathExist(t *testing.T) {
	return
}

func TestGenerateRcloneSyncCommand(t *testing.T) {
	rcloneName := "test"
	gInstance := generateInstance(rcloneName)
	expected := fmt.Sprintf("rclone sync %s:/ %s -P", rcloneName, gInstance.generateFolderPath())
	actual := gInstance.generateRcloneSyncCommand()

	assert.Equal(t, expected, actual)
}
