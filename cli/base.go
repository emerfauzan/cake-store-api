package cli

import (
	"log"

	"github.com/emerfauzan/cake-store-api/app"
	"github.com/spf13/cobra"
)

var appCli = &cobra.Command{
	Use:   "run",
	Short: "Run app",
	Run: func(cmd *cobra.Command, args []string) {
		app.NewApp()
	},
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Application Description",
}

func init() {
	rootCmd.AddCommand(appCli, migrationCli, seederCli)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
