package cmd

import (
	"github.com/spf13/cobra"
)

// DeleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete resources",
	Long: `Delete various types of resources from your documentation project.

Available resource types:
  docmeta - Delete a document metadata entry by title

Use the appropriate subcommand to delete the specific type of resource you want to remove.`,
}

func init() {
	RootCmd.AddCommand(DeleteCmd)
}
