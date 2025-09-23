package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "docli",
	Short: "A documentation CLI tool",
	Long: `docli is a command-line tool for generating, managing, 
and working with documentation in various formats.

Use docli to:
- Generate documentation from source code
- Convert between different documentation formats
- Manage documentation workflows
- Create and maintain project documentation`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to docli! Use --help to see available commands.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	// Global flags
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().BoolP("quiet", "q", false, "quiet mode")
}
