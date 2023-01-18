/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "job",
	Short: "Example of cron job command",
	Long:  `You can use this file as a job`,
	Run: func(cmd *cobra.Command, args []string) {

		//make cron job works here

		fmt.Println("job called")
	},
}

func init() {
	RootCommand.cobraCommand.AddCommand(serveCmd)
}
