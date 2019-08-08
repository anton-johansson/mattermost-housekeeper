package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var rootCommand = &cobra.Command{
	Use:   "mattermost-housekeeper",
	Short: "A tool for removing old posts in Mattermost Team Edition",
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		log.Error("Error occurred when running command:", err)
		os.Exit(1)
	}
}
