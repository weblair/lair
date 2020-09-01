package cmd

import (
	"github.com/spf13/cobra"
	"github.com/weblair/lair/internal/config"
	"github.com/weblair/lair/internal/seeding"
)

// seedCmd is invoked when the user types 'lair seed.'
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Insert seed data into the database",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnvConfig(environment)
		seeding.SeedDatabase(environment)
	},
}

// seedGenerateCmd is invoked when the user types 'lair seed generate [file]'.
var seedGenerateCmd = &cobra.Command{
	Use:   "generate [file]",
	Short: "Generate a seed YAML file from an Excel workbook",
	Long:  `The seed command relies on YAML files that define table data. Creating these for 
more than tiny datasets, however, can be extremely tedious. Using this command,
you can fill in your data in an Excel workbook and it will be converted to an
appropriate YAML file. Create a sheet for each table and place the column names
of the table in the first row. Order the sheets the way you would need to order
data in the YAML file to avoid FK problems.
`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		seeding.GenerateSeedFile(environment, args[0])
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
	seedCmd.AddCommand(seedGenerateCmd)
}
