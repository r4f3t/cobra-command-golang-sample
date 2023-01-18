/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"os"

	"github.com/r4f3t/webapi/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type RootCmd struct {
	cfgFile      string
	AppConfig    *config.Configuration
	cobraCommand *cobra.Command
}

// rootCmd represents the base command when called without any subcommands
var RootCommand = RootCmd{
	cobraCommand: &cobra.Command{
		Use:   "webapi",
		Short: "Servis Solution",
		Long:  "Service Solution Smple",
	},
	AppConfig: &config.Configuration{},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCommand.cobraCommand.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCommand.cobraCommand.PersistentFlags().StringVarP(&RootCommand.cfgFile, "config", "c", "config.qa.json", "")

}

func initConfig() {
	if RootCommand.cfgFile != "" {
		viper.SetConfigFile(RootCommand.cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("webapi.json")
	}

	viper.Set("Verbose", true)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err == nil {

		err = viper.Unmarshal(RootCommand.AppConfig)
		if err != nil {
			vipers := viper.GetViper()
			errors.New(vipers.ConfigFileUsed())
		}
	} else {
		RootCommand.cobraCommand.Help()
		// hata okuma yaparken
	}

}

func (c *RootCmd) AddCommand(cmd *cobra.Command) {
	c.cobraCommand.AddCommand(cmd)
}
