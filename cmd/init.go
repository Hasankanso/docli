package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Hasankanso/docli/internal/prompts"
	"github.com/spf13/cobra"
)

// Document represents a single document configuration
type Document struct {
	Name        string
	Description string
	FileHints   []string
}

// Config represents the entire documentation configuration
type Config struct {
	Platforms []string
	Documents []Document
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize documentation synchronization configuration",
	Long: `Initialize your documentation project by setting up platform configuration
and document specifications. This will guide you through an interactive setup
process and create a .docs/spec.md file with your configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		runInit()
	},
}

func runInit() {
	fmt.Println("Welcome to docli initialization!")
	fmt.Println("This will guide you through setting up your documentation sync configuration.")

	// Copy prompt files to user's local directory if they don't exist
	fmt.Println("📋 Copying needed prompt files...")
	err := prompts.CopyPromptFiles()
	if err != nil {
		fmt.Printf("Warning: Failed to copy prompt files: %v\n", err)
	} else {
		fmt.Println("✅ Prompt files copied successfully!")
	}

	// Step 2: Check if spec.md already exists
	specPath := filepath.Join(".docs", "spec.md")
	if _, err := os.Stat(specPath); err == nil {
		fmt.Println("ℹ️  Documentation configuration already exists!")
		fmt.Printf("Found existing configuration at: %s\n\n", specPath)
		fmt.Println("To start fresh, please remove the existing configuration first:")
		fmt.Printf("  rm %s\n", specPath)
		fmt.Println("  # or remove the entire .docs directory:")
		fmt.Println("  rm -rf .docs")
		fmt.Println("\nThen run 'docli init' again.")
		return
	}

	config := Config{}
	reader := bufio.NewReader(os.Stdin)

	// Step 3: Ask for platforms
	config.Platforms = askForPlatforms(reader)

	// Step 4: Collect document details (no need to ask count, collect until empty title)
	config.Documents = collectDocumentDetails(reader)

	// Step 5: Generate and save spec.md
	err = saveConfiguration(config)
	if err != nil {
		fmt.Printf("Error saving configuration: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n✅ Configuration saved to .docs/spec.md")
	fmt.Println("\nYou can now run other docli commands to sync your documentation!")
}

func askForPlatforms(reader *bufio.Reader) []string {
	fmt.Println("📋 Which platforms do you want to sync your documentation to?")
	fmt.Println("1. Confluence")
	fmt.Println("2. README")
	fmt.Println("\nYou can select multiple platforms by entering numbers separated by commas (e.g., 1,2)")
	fmt.Print("Select platforms (1): ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Default to Confluence if no input
	if input == "" {
		fmt.Println("Selected: Confluence")
		return []string{"confluence"}
	}

	// Parse comma-separated input
	selections := strings.Split(input, ",")
	var platforms []string
	platformMap := map[string]string{
		"1": "confluence",
		"2": "readme",
	}

	for _, selection := range selections {
		selection = strings.TrimSpace(selection)
		if platform, exists := platformMap[selection]; exists {
			// Avoid duplicates
			found := false
			for _, p := range platforms {
				if p == platform {
					found = true
					break
				}
			}
			if !found {
				platforms = append(platforms, platform)
			}
		}
	}

	if len(platforms) == 0 {
		fmt.Println("No valid platforms selected. Defaulting to Confluence.")
		return []string{"confluence"}
	}

	fmt.Printf("Selected: %s\n", strings.Join(platforms, ", "))
	return platforms
}

func collectDocumentDetails(reader *bufio.Reader) []Document {
	documents := make([]Document, 0)

	fmt.Println("\n📝 Now let's configure your documents:")
	fmt.Println("Enter document details one by one. Leave the title empty when you're done.")

	docNum := 1
	for {
		fmt.Printf("\n--- Document %d ---\n", docNum)

		// Ask for document name
		fmt.Printf("Enter title for document %d (or press Enter to finish): ", docNum)
		input, _ := reader.ReadString('\n')
		docName := strings.TrimSpace(input)

		// If empty title, we're done
		if docName == "" {
			break
		}

		// Ask for document description
		fmt.Printf("Enter description for '%s': ", docName)
		input, _ = reader.ReadString('\n')
		docDescription := strings.TrimSpace(input)

		// Ask for file/folder hints
		fmt.Printf("\n💡 Please provide file or folder names where we can find relevant content for '%s'.\n", docName)
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
			fmt.Println("No file hints provided for this document.")
		} else {
			fmt.Printf("Added %d file/folder hint(s) for '%s'\n", len(fileHints), docName)
		}

		documents = append(documents, Document{
			Name:        docName,
			Description: docDescription,
			FileHints:   fileHints,
		})

		docNum++
	}

	if len(documents) == 0 {
		fmt.Println("No documents configured.")
	} else {
		fmt.Printf("\n✅ Configured %d document(s) successfully!\n", len(documents))
	}

	return documents
}

func saveConfiguration(config Config) error {
	// Create .docs directory if it doesn't exist
	docsDir := ".docs"
	err := os.MkdirAll(docsDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create .docs directory: %w", err)
	}

	// Generate spec.md content
	content := generateSpecContent(config)

	// Write to spec.md
	specPath := filepath.Join(docsDir, "spec.md")
	err = os.WriteFile(specPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write spec.md: %w", err)
	}

	return nil
}

func generateSpecContent(config Config) string {
	var builder strings.Builder

	builder.WriteString("# Documentation Configuration\n\n")
	builder.WriteString("This file contains the configuration for your documentation synchronization.\n\n")

	// Platform section
	builder.WriteString("## Platforms\n\n")
	if len(config.Platforms) == 0 {
		builder.WriteString("No platforms configured.\n\n")
	} else {
		builder.WriteString("**Target Platforms:**\n")
		for _, platform := range config.Platforms {
			capitalized := strings.ToUpper(string(platform[0])) + platform[1:]
			builder.WriteString(fmt.Sprintf("- %s\n", capitalized))
		}
		builder.WriteString("\n")
	}

	// Documents section
	builder.WriteString("## Documents\n\n")
	if len(config.Documents) == 0 {
		builder.WriteString("No documents configured.\n")
	} else {
		for i, doc := range config.Documents {
			builder.WriteString(fmt.Sprintf("### %d. %s\n\n", i+1, doc.Name))

			if doc.Description != "" {
				builder.WriteString(fmt.Sprintf("**Description:** %s\n\n", doc.Description))
			}

			if len(doc.FileHints) > 0 {
				builder.WriteString("**File/Folder Sources:**\n")
				for _, hint := range doc.FileHints {
					builder.WriteString(fmt.Sprintf("- `%s`\n", hint))
				}
				builder.WriteString("\n")
			} else {
				builder.WriteString("*No file hints provided.*\n\n")
			}
		}
	}

	builder.WriteString("---\n")
	builder.WriteString("*Generated by docli init command*\n")

	return builder.String()
}
