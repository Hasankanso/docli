package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of docli",
	Long:  `Display the current version of docli.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("docli v2.0.0")
		fmt.Println("Built with Go")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
