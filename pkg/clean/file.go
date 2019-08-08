package clean

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func removeFiles(mattermostDataDirectory string, filePaths ...string) error {
	for _, filePath := range filePaths {
		if err := removeFile(mattermostDataDirectory, filePath); err != nil {
			return err
		}
	}
	return nil
}

func removeFile(mattermostDataDirectory, filePath string) error {
	if len(filePath) > 0 {
		log.Info("Removing file from disk:", filePath)
		if true {
			// TODO: remove this if
			return nil
		}
		fullPath := filepath.Join(mattermostDataDirectory, filePath)
		return os.Remove(fullPath)
	}
	return nil
}