/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Example of cron job command",
	Long:  `You can use this file as a job`,
	Run: func(cmd *cobra.Command, args []string) {

		//make consumer

		fmt.Println("consumer called")
	},
}

func init() {
	RootCommand.cobraCommand.AddCommand(consumerCmd)
}
