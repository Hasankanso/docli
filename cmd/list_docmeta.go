package cmd

import (
	"github.com/Hasankanso/docli/internal/docmeta"
	"github.com/Hasankanso/docli/internal/logger"
	"github.com/Hasankanso/docli/internal/spec"
	"github.com/spf13/cobra"
)

// ListDocmetaCmd represents the list docmeta command
var ListDocmetaCmd = &cobra.Command{
	Use:   "docmeta",
	Short: "List all document metadata entries",
	Long: `List all document metadata entries from your spec.md file.
This command displays all configured documents with their names and descriptions.`,
	Run: func(cmd *cobra.Command, args []string) {
		runListDocmeta()
	},
}

func runListDocmeta() {
	specRepo := spec.NewSpecRepo()
	if !specRepo.SpecExists() {
		logger.Error("No spec.md file found. Please run 'docli init' to initialize your project")
		return
	}
	ListDocMetaCmd := docmeta.NewListDocMetaCommand(specRepo)
	ListDocMetaCmd.Run()
}

func init() {
	ListCmd.AddCommand(ListDocmetaCmd)
}
