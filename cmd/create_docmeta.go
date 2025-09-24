package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/Hasankanso/docli/internal/docmeta"
	"github.com/Hasankanso/docli/internal/logger"
	"github.com/Hasankanso/docli/internal/spec"
	"github.com/spf13/cobra"
)

// createDocmetaCmd represents the create docmeta command
var createDocmetaCmd = &cobra.Command{
	Use:   "docmeta",
	Short: "Create a new document metadata entry",
	Long: `Create a new document metadata entry and add it to your spec.md file.
This command will guide you through an interactive process to define a new
document with its name, description, and file hints.`,
	Run: func(cmd *cobra.Command, args []string) {
		runCreateDocmeta()
	},
}

func runCreateDocmeta() {
	reader := bufio.NewReader(os.Stdin)
	newDocMeta := CollectSingleDocumentDetails(reader)
	if newDocMeta == nil {
		logger.Info("Document creation cancelled")
		return
	}
	specRepo := spec.NewSpecRepo()
	createCmd := docmeta.NewCreateDocMetaCommand(specRepo, newDocMeta)

	createCmd.Run()
}

func CollectSingleDocumentDetails(reader *bufio.Reader) *spec.DocMetaData {
	logger.Info("\n--- New Document Configuration ---")

	// Ask for document name
	logger.Info("Enter document title: ")
	input, _ := reader.ReadString('\n')
	docName := strings.TrimSpace(input)

	// If empty title, cancel
	if docName == "" {
		return nil
	}

	// Ask for document description
	logger.Info("Enter description for '%s': ", docName)
	input, _ = reader.ReadString('\n')
	docDescription := strings.TrimSpace(input)

	// Ask for file/folder hints
	logger.Info("\nPlease provide file or folder names where we can find relevant content for '%s'.", docName)
	logger.Info("You can specify multiple files/folders. Press Enter on an empty line when done.")

	var fileHints []string
	hintNum := 1
	for {
		logger.Info("  File/Folder %d (or press Enter to finish): ", hintNum)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			break
		}

		fileHints = append(fileHints, input)
		hintNum++
	}

	if len(fileHints) == 0 {
		logger.Info("No file hints provided for '%s'.", docName)
	} else {
		logger.Info("Added %d file/folder hint(s) for '%s'", len(fileHints), docName)
	}

	return spec.NewDocMetaData(docName, docDescription, fileHints)
}
