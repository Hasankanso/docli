package cmd

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/Hasankanso/docli/internal/common"
	"github.com/Hasankanso/docli/internal/spec"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize basic documentation project structure",
	Long: `Initialize your documentation project by setting up the basic configuration
structure. This will copy prompt files and create the initial spec.md file
with platform configuration. Use 'docli create docmeta' to add document metadata.`,
	Run: func(cmd *cobra.Command, args []string) {
		runInit()
	},
}

func runInit() {
	common.Info("Welcome to docli initialization!")

	specRepo := spec.NewSpecRepo()
	reader := bufio.NewReader(os.Stdin)

	// Step 3: Ask for platforms
	platforms := askForPlatforms(reader)

	// Step 4: Initialize spec repository and save initial config
	spec.NewInitSpecCommand(specRepo, platforms).Run()
}

func askForPlatforms(reader *bufio.Reader) []string {
	fmt.Println("Which platforms do you want to sync your documentation to?")
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
			found := slices.Contains(platforms, platform)
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

func init() {
	RootCmd.AddCommand(initCmd)
}
