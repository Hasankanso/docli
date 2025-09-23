package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Hasankanso/docli/internal/common"
	"github.com/Hasankanso/docli/internal/docmeta"
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
		common.Info("Document creation cancelled")
		return
	}
	specRepo := spec.NewSpecRepo()
	createCmd := docmeta.NewCreateDocMetaCommand(specRepo, newDocMeta)

	createCmd.Run()
}

func CollectSingleDocumentDetails(reader *bufio.Reader) *spec.DocMetaData {
	fmt.Println("\n--- New Document Configuration ---")

	// Ask for document name
	fmt.Print("Enter document title: ")
	input, _ := reader.ReadString('\n')
	docName := strings.TrimSpace(input)

	// If empty title, cancel
	if docName == "" {
		return nil
	}

	// Ask for document description
	fmt.Printf("Enter description for '%s': ", docName)
	input, _ = reader.ReadString('\n')
	docDescription := strings.TrimSpace(input)

	// Ask for file/folder hints
	fmt.Printf("\nPlease provide file or folder names where we can find relevant content for '%s'.\n", docName)
	fmt.Println("You can specify multiple files/folders. Press Enter on an empty line when done.")

	var fileHints []string
	hintNum := 1
	for {
		fmt.Printf("  File/Folder %d (or press Enter to finish): ", hintNum)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			break
		}

		fileHints = append(fileHints, input)
		hintNum++
	}

	if len(fileHints) == 0 {
		fmt.Printf("No file hints provided for '%s'.\n", docName)
	} else {
		fmt.Printf("Added %d file/folder hint(s) for '%s'\n", len(fileHints), docName)
	}

	return spec.NewDocMetaData(docName, docDescription, fileHints)
}
