package main

import (
	"github.com/anton-johansson/mattermost-housekeeper/pkg/clean"
	"github.com/anton-johansson/mattermost-housekeeper/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	mattermostDataDirectory string
	databaseName            string
	databaseUser            string
	databasePassword        string
	databaseHost            string
	databasePort            int
	retentionDays           int
	fileBatchSize           int
	logFormat               string
	logLevel                string
)

func init() {
	command := &cobra.Command{
		Use:   "clean",
		Short: "Cleans a Mattermost Team Edition",
		Run: func(command *cobra.Command, arguments []string) {
			log.Initialize(logFormat, logLevel)
			if err := clean.Clean(mattermostDataDirectory, databaseName, databaseUser, databasePassword, databaseHost, databasePort, retentionDays, fileBatchSize); err != nil {
				logrus.Error("Error occurred when running the clean command: ", err)
			}
		},
	}
	command.Flags().StringVar(&mattermostDataDirectory, "data-dir", "", "The directory of Mattermost data")
	command.Flags().StringVar(&databaseName, "database-name", "", "The name of the PostgreSQL database")
	command.Flags().StringVar(&databaseUser, "database-user", "", "The user to connect to the PostgreSQL database with")
	command.Flags().StringVar(&databasePassword, "database-password", "", "The password for the user to connect to the PostgreSQL database with")
	command.Flags().StringVar(&databaseHost, "database-host", "", "The hostname of the PostgreSQL database")
	command.Flags().IntVar(&databasePort, "database-port", 5432, "The port of the PostgreSQL database")
	command.Flags().IntVar(&retentionDays, "retention-days", 365, "The age in days of the posts to delete")
	command.Flags().IntVar(&fileBatchSize, "file-batch-size", 50, "The number of files to fetch and remove in each batch")
	command.Flags().StringVar(&logFormat, "log-format", "text", "The format of logs")
	command.Flags().StringVar(&logLevel, "log-level", "info", "The minimum level of logs")
	rootCommand.AddCommand(command)
}
