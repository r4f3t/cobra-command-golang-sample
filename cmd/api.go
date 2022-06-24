/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/r4f3t/webapi/internal/user/controller"
	"github.com/spf13/cobra"
)

type api struct {
	instance *echo.Echo
	command  *cobra.Command
	Port     string
}

// apiCmd represents the api command
var apiCmd = &api{
	command: &cobra.Command{
		Use:   "api",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("api called")
		},
	},
	Port: "5000",
}

func init() {
	rootCmd.AddCommand(apiCmd.command)
	apiCmd.instance = echo.New()

	// repoları newleconst

	// servisleri newle

	controller.MakeHandler(apiCmd.instance, nil)

}
