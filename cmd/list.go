package cmd

import (
	"github.com/spf13/cobra"
)

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List resources",
	Long: `List various types of resources in your documentation project.

Available resource types:
  docmeta - List all document metadata entries

Use the appropriate subcommand to list the specific type of resource you want to view.`,
}

func init() {
	RootCmd.AddCommand(ListCmd)
}
