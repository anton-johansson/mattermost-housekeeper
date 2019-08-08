package clean

import (
	"errors"
	"os"
)

func validate(mattermostDataDirectory, databaseName, databaseUser, databaseHost string, retentionDays, fileBatchSize int) error {
	if mattermostDataDirectory == "" {
		return errors.New("Mattermost data directory is required")
	}
	fileInfo, err := os.Stat(mattermostDataDirectory)
	if err != nil {
		return err
	}
	if !fileInfo.Mode().IsDir() {
		return errors.New("Mattermost data directory must be an existing directory")
	}
	if databaseName == "" {
		return errors.New("A database name is required")
	}
	if databaseUser == "" {
		return errors.New("A database user is required")
	}
	if databaseHost == "" {
		return errors.New("A database host is required")
	}
	if fileBatchSize <= 0 {
		return errors.New("File batch size must be a positive integer")
	}
	if retentionDays <= 0 {
		return errors.New("Retention days must be a positive integer")
	}
	return nil
}