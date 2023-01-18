package cmd

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/r4f3t/webapi/internal/user"
	usercontroller "github.com/r4f3t/webapi/internal/user/controller"
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
		Short: "",
		Long:  "",
	},
	Port: "5000",
}

func init() {
	RootCommand.AddCommand(apiCmd.command)
	//defines flags for api command specific
	apiCmd.command.Flags().StringVarP(&apiCmd.Port, "port", "p", "5000", "Service Port")
	apiCmd.instance = echo.New()

	apiCmd.command.RunE = func(cmd *cobra.Command, args []string) error {

		// construct repository
		userRepository := user.NewRepository(RootCommand.AppConfig.SqlDbSettings.Uri, RootCommand.AppConfig.SqlDbSettings.DatabaseName)

		// construct service
		userService := user.NewService(userRepository)

		usercontroller.MakeHandler(apiCmd.instance, usercontroller.NewController(userService))

		apiCmd.instance.Logger.Fatal(apiCmd.instance.Start(fmt.Sprintf(":%s", apiCmd.Port)))

		return nil
	}

}
