package cmd

import (
	"github.com/spf13/cobra"
	"github.com/weblair/lair/config"
	"github.com/weblair/lair/db"
)

// seedCmd is invoked when the user types 'lair seed.'
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Insert seed data into the database",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnvConfig(environment)
		db.SeedDatabase(environment)
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
