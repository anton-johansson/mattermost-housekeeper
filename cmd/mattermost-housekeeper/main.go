package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCommand = &cobra.Command{
	Use:   "mattermost-housekeeper",
	Short: "A tool for removing old posts in Mattermost Team Edition",
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
