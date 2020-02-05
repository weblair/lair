package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/weblair/lair/config"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Operations on a Lair project's configuration files (NOT FULLY IMPLEMENTED)",
	Long:  ``,
}

// configHMACCmd represents the config hmac subcommand
var configHMACCmd = &cobra.Command{
	Use:   "hmac",
	Short: "Generate a new HMAC key",
	Long: `Generates a new cryptographically secure base-64 HMAC key and stores it in the 
.env file. If the .env file does not exist, a new one will be generated.`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Warn("HMAC Key generation is not fully implemented.")
		logrus.Warn("If a new .env file results from this command, it might not be complete.")
		config.CreateGinDotEnv()
	},
}

// configInitCmd represents the config init subcommand
var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a configuration file for a project (NOT IMPLEMENTED)",
	Long:  "",
}

// configInitDotenvCmd represents the config init dotenv subcommand
var configInitDotenvCmd = &cobra.Command{
	Use:   "dotenv",
	Short: "Initialize a dotenv file (NOT IMPLEMENTED)",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Fatal("Dotenv initialization is not implemented yet.")
	},

}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configHMACCmd, configInitCmd)
	configInitCmd.AddCommand(configInitDotenvCmd)
}
