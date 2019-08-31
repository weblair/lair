package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// TODO: Run database migrations from GitHub
// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a Gin project to staging or production (NOT IMPLEMENTED)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Fatal("Deployments are not implemented yet.")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
