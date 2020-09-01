package cmd

import (
	"github.com/spf13/cobra"
	"github.com/weblair/lair/internal/config"
	"github.com/weblair/lair/internal/database"
	"github.com/weblair/lair/internal/migrations"
	"github.com/weblair/lair/internal/seeding"
)

var forceCreate bool
var createAndSeed bool

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new database",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnvConfig(environment)
		database.CreateDatabaseFromConfig(forceCreate)
		if createAndSeed {
			migrations.MigrateDatabase(0)
			seeding.SeedDatabase(environment)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().BoolVarP(
		&forceCreate,
		"force",
		"f",
		false,
		"If the database already exists, drop it",
	)

	createCmd.Flags().BoolVarP(
		&createAndSeed,
		"seed",
		"s",
		false,
		"Insert seed data into the database after it is created",
	)
}
