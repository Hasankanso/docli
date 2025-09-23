package cmd

import (
	"github.com/Hasankanso/docli/internal/docmeta"
	"github.com/Hasankanso/docli/internal/spec"
	"github.com/spf13/cobra"
)

// DeleteDocmetaCmd represents the delete docmeta command
var DeleteDocmetaCmd = &cobra.Command{
	Use:   "docmeta <id>",
	Short: "Delete a document metadata entry by id",
	Long: `Delete a document metadata entry from your spec.md file by providing the document id.
The id should be provided in quotes if it contains spaces.

Example:
  docli delete docmeta "API Documentation"
  docli delete docmeta README`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		runDeleteDocmeta(id)
	},
}

func runDeleteDocmeta(id string) {
	specRepo := spec.NewSpecRepo()
	deleteCmd := docmeta.NewDeleteDocMetaCommand(specRepo, id)
	deleteCmd.Run()
}

func init() {
	RootCmd.AddCommand(DeleteCmd)
	DeleteCmd.AddCommand(DeleteDocmetaCmd)
}
