package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Operations on a Lair project's configuration files (NOT IMPLEMENTED)",
	Long:  ``,
}

// configHMACCmd represents the config hmac subcommand
var configHMACCmd = &cobra.Command{
	Use:   "hmac",
	Short: "Generate a new HMAC key (NOT IMPLEMENTED)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Fatal("HMAC key generation is not implemented yet.")
	},
}

// configInitCmd represents the config init subcommand
var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a configuration file for a project (NOT IMPLEMENTED)",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Fatal("Config initialization is not implemented yet.")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configHMACCmd, configInitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
