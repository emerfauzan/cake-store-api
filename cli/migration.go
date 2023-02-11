package cli

import (
	"fmt"

	"github.com/emerfauzan/cake-store-api/config"
	"github.com/emerfauzan/cake-store-api/migration"
	"github.com/spf13/cobra"
)

var migrationCli = &cobra.Command{
	Use:   "migrate",
	Short: "Database migrations tool",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var createMigrationCli = &cobra.Command{
	Use:   "create",
	Short: "Create a new empty migrations file",
	Run: func(cli *cobra.Command, args []string) {
		name, err := cli.Flags().GetString("name")
		if err != nil {
			fmt.Println("Unable to read flag `name`", err.Error())
			return
		}

		if err := migration.Create(name); err != nil {
			fmt.Println("Unable to create migration", err.Error())
			return
		}
	},
}

var upMigrationCli = &cobra.Command{
	Use:   "up",
	Short: "Run up migrations",
	Run: func(cli *cobra.Command, args []string) {
		step, err := cli.Flags().GetInt("step")
		if err != nil {
			fmt.Println("Unable to read flag `step`")
			return
		}

		db := config.NewDB()

		migrator, err := migration.Init(db)
		if err != nil {
			fmt.Println("Unable to fetch migrator")
			return
		}

		err = migrator.Up(step)
		if err != nil {
			fmt.Println("Unable to run `up` migrations")
			return
		}
	},
}

var downMigrationCli = &cobra.Command{
	Use:   "down",
	Short: "Run down migrations",
	Run: func(cli *cobra.Command, args []string) {
		step, err := cli.Flags().GetInt("step")
		if err != nil {
			fmt.Println("Unable to read flag `step`")
			return
		}

		db := config.NewDB()

		migrator, err := migration.Init(db)
		if err != nil {
			fmt.Println("Unable to fetch migrator")
			return
		}

		err = migrator.Down(step)
		if err != nil {
			fmt.Println("Unable to run `down` migrations")
			return
		}
	},
}

var statusMigrationCli = &cobra.Command{
	Use:   "status",
	Short: "display status of each migrations",
	Run: func(cmd *cobra.Command, args []string) {
		db := config.NewDB()

		migrator, err := migration.Init(db)
		if err != nil {
			fmt.Println("Unable to fetch migrator")
			return
		}

		if err := migrator.MigrationStatus(); err != nil {
			fmt.Println("Unable to fetch migration status")
			return
		}

		return
	},
}

func init() {
	createMigrationCli.Flags().StringP("name", "n", "", "Name for the migration")
	upMigrationCli.Flags().IntP("step", "s", 0, "Number of migrations to execute")
	downMigrationCli.Flags().IntP("step", "s", 0, "Number of migrations to execute")
	migrationCli.AddCommand(createMigrationCli, upMigrationCli, downMigrationCli, statusMigrationCli)
}
