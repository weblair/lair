package cmd

import (
	"github.com/spf13/cobra"
	"github.com/weblair/lair/config"
	"github.com/weblair/lair/db"
)

// dropCmd represents the drop command
var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop the database",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnvConfig(environment)
		db.DropDatabaseFromConfig()
	},
}

func init() {
	rootCmd.AddCommand(dropCmd)
}
