package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print this build's version string",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString("VERSION"))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
