package prompts

import (
	"fmt"
	"os"
	"path/filepath"
)

func CopyPromptFiles() error {
	// Create .github/prompts directory if it doesn't exist
	promptsDir := filepath.Join(".github", "prompts")
	err := os.MkdirAll(promptsDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create .github/prompts directory: %w", err)
	}

	// Get the executable path to find the source prompt files
	execPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	// Get the directory containing the executable
	execDir := filepath.Dir(execPath)

	// Try different possible locations for the prompt files
	var sourcePromptsDir string

	// Check if we're running from the development directory
	devPromptsDir := filepath.Join(execDir, ".github", "prompts")
	if _, err := os.Stat(devPromptsDir); err == nil {
		sourcePromptsDir = devPromptsDir
	} else {
		// Check if prompts are bundled with the executable
		bundledPromptsDir := filepath.Join(execDir, "prompts")
		if _, err := os.Stat(bundledPromptsDir); err == nil {
			sourcePromptsDir = bundledPromptsDir
		} else {
			// Fallback: try relative to executable assuming it's in docli repo
			fallbackPromptsDir := filepath.Join(filepath.Dir(execDir), ".github", "prompts")
			if _, err := os.Stat(fallbackPromptsDir); err == nil {
				sourcePromptsDir = fallbackPromptsDir
			} else {
				return fmt.Errorf("could not find source prompt files")
			}
		}
	}

	// Copy syncDoc.prompt.md if it doesn't exist
	syncDocPath := filepath.Join(promptsDir, "syncDoc.prompt.md")
	if _, err := os.Stat(syncDocPath); os.IsNotExist(err) {
		sourceSyncDoc := filepath.Join(sourcePromptsDir, "syncDoc.prompt.md")
		err = copyFile(sourceSyncDoc, syncDocPath)
		if err != nil {
			return fmt.Errorf("failed to copy syncDoc.prompt.md: %w", err)
		}
		fmt.Println("✅ Created .github/prompts/syncDoc.prompt.md")
	}

	// Copy updateDoc.prompt.md if it doesn't exist
	updateDocPath := filepath.Join(promptsDir, "updateDoc.prompt.md")
	if _, err := os.Stat(updateDocPath); os.IsNotExist(err) {
		sourceUpdateDoc := filepath.Join(sourcePromptsDir, "updateDoc.prompt.md")
		err = copyFile(sourceUpdateDoc, updateDocPath)
		if err != nil {
			return fmt.Errorf("failed to copy updateDoc.prompt.md: %w", err)
		}
		fmt.Println("✅ Created .github/prompts/updateDoc.prompt.md")
	}

	return nil
}

func copyFile(src, dst string) error {
	sourceData, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("failed to read source file %s: %w", src, err)
	}

	err = os.WriteFile(dst, sourceData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write destination file %s: %w", dst, err)
	}

	return nil
}
