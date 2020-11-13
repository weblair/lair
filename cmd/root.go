package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/cobra"
)

// environment represents the config that Lair should use when connecting to the database.
var environment string

// loglevel represents the override value for Lair's logging.
var loglevel string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lair",
	Short: "Database migration tool for Weblair projects",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	ll, err := logrus.ParseLevel(loglevel)
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.Warn(fmt.Sprintf("%s is not a valid log level. Setting log level to 'info.'", loglevel))
	} else {
		logrus.SetLevel(ll)
	}

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

	rootCmd.PersistentFlags().StringVarP(
		&loglevel,
		"loglevel",
		"l",
		"info",
		"Override the default log level, valid values are: panic, fatal, error, warn, info, debug, trace",
	)

}
