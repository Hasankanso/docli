package cmd

import "github.com/spf13/cobra"

// CreateCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create resources",
	Long: `Create various types of resources in your documentation project.

Available resource types:
  docmeta - Create a new document metadata entry

The create command provides subcommands to create different types of resources
that help organize and manage your project documentation.`,
}

func init() {
	RootCmd.AddCommand(CreateCmd)
	CreateCmd.AddCommand(createDocmetaCmd)
}
