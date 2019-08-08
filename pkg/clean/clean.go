package clean

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"time"
)

// Clean performs the actual cleaning duty
func Clean(mattermostDataDirectory, databaseName, databaseUser, databasePassword, databaseHost string, databasePort, retentionDays, fileBatchSize int) error {
	if err := validate(mattermostDataDirectory, databaseName, databaseUser, databaseHost, retentionDays, fileBatchSize); err != nil {
		return err
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", databaseHost, databasePort, databaseUser, databasePassword, databaseName)
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	defer database.Close()

	if err = database.Ping(); err != nil {
		return err
	}

	millisecondEpoch := time.Now().AddDate(0, 0, -retentionDays).UnixNano() / 1000000
	if err := cleanFiles(database, millisecondEpoch, mattermostDataDirectory, fileBatchSize); err != nil {
		return err
	}

	if err := deleteFileInfoRows(database, millisecondEpoch); err != nil {
		return err
	}

	return deletePostRows(database, millisecondEpoch)
}

func cleanFiles(database *sql.DB, millisecondEpoch int64, mattermostDataDirectory string, fileBatchSize int) error {
	batch := 0

	moreResults := true
	for moreResults {
		haveMoreResults, err := cleanFilesBatch(database, millisecondEpoch, mattermostDataDirectory, fileBatchSize, batch)
		if err != nil {
			return nil
		}
		moreResults = haveMoreResults
		batch = batch + 1
	}

	return nil
}

func cleanFilesBatch(database *sql.DB, millisecondEpoch int64, mattermostDataDirectory string, fileBatchSize, batch int) (bool, error) {
	rows, err := database.Query(`
		SELECT  info.path
		,       info.thumbnailpath
		,       info.previewpath
		FROM    fileinfo info
		WHERE   info.createat < $1
		OFFSET  $2
		LIMIT   $3;`,
		millisecondEpoch,
		fileBatchSize*batch,
		fileBatchSize)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	var (
		path          string
		thumbnailPath string
		previewPath   string
	)

	moreResults := false
	for rows.Next() {
		moreResults = true
		if err := rows.Scan(&path, &thumbnailPath, &previewPath); err != nil {
			return false, err
		}

		if err := removeFiles(mattermostDataDirectory, path, thumbnailPath, previewPath); err != nil {
			return false, err
		}
	}

	return moreResults, nil
}

func deleteFileInfoRows(database *sql.DB, millisecondEpoch int64) error {
	result, err := database.Exec(`
		DELETE
		FROM    fileinfo info
		WHERE   info.createat < $1`,
		millisecondEpoch)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Info("Removed", rows, "file information rows")
	return nil
}

func deletePostRows(database *sql.DB, millisecondEpoch int64) error {
	result, err := database.Exec(`
		DELETE
		FROM    posts post
		WHERE   post.createat < $1`,
		millisecondEpoch)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Info("Removed", rows, "post rows")
	return nil
}
