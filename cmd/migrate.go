package cmd

import (
	"github.com/spf13/cobra"
	"github.com/weblair/lair/internal/config"
	"github.com/weblair/lair/internal/migrations"
)

// migrateSteps indicates how many migration steps up or down should be run against the database.
var migrateSteps int

// migrateCmd is invoked when the user types 'lair migrate'.
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations against the current database",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnvConfig(environment)
		migrations.MigrateDatabase(migrateSteps)
	},
}

// migrateCurrentCmd is invoked when the user types 'lair migrate current'.
var migrateCurrentCmd = &cobra.Command{
	Use:   "current",
	Short: "Print the current database version",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnvConfig(environment)
		migrations.GetDatabaseVersion(environment)
	},
}

// migrateNewCmd is invoked when the user types 'lair migrate new [new migration]'.
var migrateNewCmd = &cobra.Command{
	Use:   "new [description]",
	Short: "Create a new database migration",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		migrations.CreateNewMigrationFile(args[0])
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(migrateNewCmd, migrateCurrentCmd)

	migrateCmd.Flags().IntVarP(
		&migrateSteps,
		"steps",
		"n",
		0,
		"Migrate up n steps if positive, down n steps if negative",
	)
}
