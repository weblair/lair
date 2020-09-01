package cmd

import (
	"github.com/spf13/cobra"
	"github.com/weblair/lair/internal/config"
	"github.com/weblair/lair/internal/database"
)

// dropCmd represents the drop command
var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop the database",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnvConfig(environment)
		database.DropDatabaseFromConfig()
	},
}

func init() {
	rootCmd.AddCommand(dropCmd)
}
