package cmd

import (
	"github.com/spf13/cobra"
	"github.com/weblair/lair/internal/config"
	"github.com/weblair/lair/internal/database"
)

// resetAndSeed is the flag that indicates seeding should take place after resetting
var resetAndSeed bool

// resetCmd is invoked when the user types 'lair reset.'
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Drop then re-create the database",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnvConfig(environment)
		database.CreateDatabaseFromConfig(true)
		database.MigrateDatabase(0)
		if resetAndSeed {
			database.SeedDatabase(environment)
		}
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)

	resetCmd.Flags().BoolVarP(
		&resetAndSeed,
		"seed",
		"s",
		false,
		"Insert seed data into the database after it has been re-created",
	)
}
