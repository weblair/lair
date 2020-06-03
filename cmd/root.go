package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// environment represents the config that Lair should use when connecting to the database.
var environment string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lair",
	Short: "Database migration tool for Weblair projects",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&environment,
		"env",
		"e",
		"development",
		"Use the given environment's database configs",
	)
}
