package cmd

import (
	"github.com/spf13/cobra"
	"github.com/weblair/lair/config"
	"github.com/weblair/lair/db"
)

var environment string
var forceCreate bool
var createAndSeed bool
var migrateSteps int
var resetAndSeed bool

// lair db
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Project database operations.",
	Long: ``,
}

// lair db create
var dbCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new database",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnvConfig(environment)
		db.CreateDatabaseFromConfig(forceCreate)
		if createAndSeed {
			db.MigrateDatabase(0)
			db.SeedDatabase()
		}
	},
}

// lair db drop
var dbDropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop the database",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnvConfig(environment)
		db.DropDatabaseFromConfig()
	},
}

// lair db migrate
var dbMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations against the current database",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnvConfig(environment)
		db.MigrateDatabase(migrateSteps)
	},
}

// lair db migrate current
var dbMigrateCurrentCmd = &cobra.Command{
	Use: "current",
	Short: "Print the current database version",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnvConfig(environment)
		db.GetDatabaseVersion(environment)
	},
}

// lair db migrate new [new migration]
var dbMigrateNewCmd = &cobra.Command{
	Use: "new [description]",
	Short: "Create a new database migration",
	Long: "",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db.CreateNewMigrationFile(args[0])
	},
}

var dbResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Drop then re-create the database",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnvConfig(environment)
		db.CreateDatabaseFromConfig(true)
		db.MigrateDatabase(0)
		if resetAndSeed {
			db.SeedDatabase()
		}
	},
}

var dbSeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Insert seed data into the database",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnvConfig(environment)
		db.SeedDatabase()
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(dbCreateCmd, dbDropCmd, dbMigrateCmd, dbResetCmd, dbSeedCmd)
	dbMigrateCmd.AddCommand(dbMigrateNewCmd, dbMigrateCurrentCmd)

	dbCmd.PersistentFlags().StringVarP(
		&environment,
		"env",
		"e",
		"development",
		"Use the given environment's database configs",
	)

	// db create flags
	dbCreateCmd.Flags().BoolVarP(
		&forceCreate,
		"force",
		"f",
		false,
		"If the database already exists, drop it",
	)
	dbCreateCmd.Flags().BoolVarP(
		&createAndSeed,
		"seed",
		"s",
		false,
		"Insert seed data into the database after it is created",
	)

	// db migrate flags
	dbMigrateCmd.Flags().IntVarP(
		&migrateSteps,
		"steps",
		"n",
		0,
		"Migrate up n steps if positive, down n steps if negative",
	)

	// db reset flags
	dbSeedCmd.Flags().BoolVarP(
		&resetAndSeed,
		"seed",
		"s",
		false,
		"Insert seed data into the database after it has been re-created",
	)
}
