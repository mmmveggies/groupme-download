package util

import (
	"os"
	"path/filepath"
)

func GetApplicationDir() (string, error) {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(configPath, "groupme-files")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return "", err
		}
	}

	return path, err
}
