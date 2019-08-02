package main

import (
	"github.com/anton-johansson/mattermost-maid/pkg/clean"
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
)

func init() {
	command := &cobra.Command{
		Use:   "clean",
		Short: "Cleans a Mattermost Team Edition",
		RunE: func(command *cobra.Command, arguments []string) error {
			return clean.Clean(mattermostDataDirectory, databaseName, databaseUser, databasePassword, databaseHost, databasePort, retentionDays, fileBatchSize)
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
	rootCommand.AddCommand(command)
}
