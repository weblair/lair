package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/weblair/lair/generator"
)

var withAuth bool

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [project name]",
	Short: "Create a new Gin project",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if withAuth {
			logrus.Warn("Initializng a project with auth middleware is not supported yet.")
			logrus.Warn("This project will be initialized without the middleware.")
		}
		generator.NewGinProject(args[0], withAuth)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().BoolVarP(
		&withAuth,
		"auth",
		"a",
		false,
		"Init project with auth middleware",
	)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
