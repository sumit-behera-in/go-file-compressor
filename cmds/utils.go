package cmds

import (
	"os"
	"path/filepath"
)

func GetDefaultDownloadPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, "Downloads"), nil
}
