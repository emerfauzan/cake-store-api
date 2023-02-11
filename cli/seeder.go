package cli

import (
	"fmt"

	"github.com/emerfauzan/cake-store-api/config"
	"github.com/emerfauzan/cake-store-api/seeder"
	"github.com/spf13/cobra"
)

var seederCli = &cobra.Command{
	Use:   "seed",
	Short: "Database seeder tool",
	Run: func(cmd *cobra.Command, args []string) {
		db := config.NewDB()
		seeder, _ := seeder.Init(db)
		err := seeder.Seed()

		if err != nil {
			fmt.Println(err)
		}
	},
}
