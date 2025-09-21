package prompts

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func CopyPromptFiles() error {
	// Create .github/prompts directory if it doesn't exist
	promptsDir := filepath.Join(".github", "prompts")
	err := os.MkdirAll(promptsDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create .github/prompts directory: %w", err)
	}

	// GitHub repository information
	const (
		baseURL = "https://raw.githubusercontent.com/Hasankanso/docli/main"
	)

	// Copy syncDoc.prompt.md if it doesn't exist
	syncDocPath := filepath.Join(promptsDir, "syncDoc.prompt.md")
	if _, err := os.Stat(syncDocPath); os.IsNotExist(err) {
		syncDocURL := baseURL + "/.github/prompts/syncDoc.prompt.md"
		err = fetchFileFromGit(syncDocURL, syncDocPath)
		if err != nil {
			return fmt.Errorf("failed to fetch syncDoc.prompt.md: %w", err)
		}
		fmt.Println("✅ Created .github/prompts/syncDoc.prompt.md")
	}

	// Copy updateDoc.prompt.md if it doesn't exist
	updateDocPath := filepath.Join(promptsDir, "updateDoc.prompt.md")
	if _, err := os.Stat(updateDocPath); os.IsNotExist(err) {
		updateDocURL := baseURL + "/.github/prompts/updateDoc.prompt.md"
		err = fetchFileFromGit(updateDocURL, updateDocPath)
		if err != nil {
			return fmt.Errorf("failed to fetch updateDoc.prompt.md: %w", err)
		}
		fmt.Println("✅ Created .github/prompts/updateDoc.prompt.md")
	}

	return nil
}

func fetchFileFromGit(url, destPath string) error {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Make the HTTP request
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("network error fetching %s: %w", url, err)
	}
	defer resp.Body.Close()

	// Check if the response is successful
	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("file not found at %s (HTTP 404)", url)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server error fetching %s: HTTP %d", url, resp.StatusCode)
	}

	// Read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response from %s: %w", url, err)
	}

	// Validate that we got some content
	if len(data) == 0 {
		return fmt.Errorf("received empty file from %s", url)
	}

	// Write the data to the destination file
	err = os.WriteFile(destPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", destPath, err)
	}

	return nil
}
