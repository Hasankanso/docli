package prompts

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Hasankanso/docli/internal/logger"
)

// GitHubFile represents a file in a GitHub directory
type GitHubFile struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Type        string `json:"type"`
	DownloadURL string `json:"download_url"`
}

// GitHubFileContent represents the content of a file from GitHub API
type GitHubFileContent struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}

func CopyPromptFiles() error {
	// Create .github/prompts directory if it doesn't exist
	promptsDir := filepath.Join(".github", "prompts")
	err := os.MkdirAll(promptsDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create .github/prompts directory: %w", err)
	}

	// GitHub repository information
	const (
		apiURL = "https://api.github.com/repos/Hasankanso/docli/contents/.github/prompts"
	)

	// Get list of all prompt files from GitHub
	promptFiles, err := getPromptFilesFromGitHub(apiURL)
	if err != nil {
		return fmt.Errorf("failed to get prompt files list: %w", err)
	}

	// Copy each prompt file if it doesn't exist locally
	for _, file := range promptFiles {
		// Only process .prompt.md files
		if !strings.HasSuffix(file.Name, ".prompt.md") {
			continue
		}

		localPath := filepath.Join(promptsDir, file.Name)
		if _, err := os.Stat(localPath); os.IsNotExist(err) {
			err = fetchFileContentFromGitHub(file.Path, localPath)
			if err != nil {
				return fmt.Errorf("failed to fetch %s: %w", file.Name, err)
			}
			logger.Success("Created .github/prompts/%s", file.Name)
		}
	}

	return nil
}

func getPromptFilesFromGitHub(apiURL string) ([]GitHubFile, error) {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Make the HTTP request to GitHub API
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("network error fetching directory listing: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response is successful
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("prompts directory not found in repository (HTTP 404)")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server error fetching directory listing: HTTP %d", resp.StatusCode)
	}

	// Read and parse the JSON response
	var files []GitHubFile
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&files)
	if err != nil {
		return nil, fmt.Errorf("failed to parse GitHub API response: %w", err)
	}

	return files, nil
}

func fetchFileContentFromGitHub(filePath, destPath string) error {
	// Construct GitHub API URL for file content
	const repoAPI = "https://api.github.com/repos/Hasankanso/docli/contents/"
	fileURL := repoAPI + filePath

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Make the HTTP request to GitHub API
	resp, err := client.Get(fileURL)
	if err != nil {
		return fmt.Errorf("network error fetching %s: %w", filePath, err)
	}
	defer resp.Body.Close()

	// Check if the response is successful
	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("file not found at %s (HTTP 404)", filePath)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server error fetching %s: HTTP %d", filePath, resp.StatusCode)
	}

	// Parse the JSON response to get file content
	var fileContent GitHubFileContent
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&fileContent)
	if err != nil {
		return fmt.Errorf("failed to parse GitHub API response for %s: %w", filePath, err)
	}

	// Decode base64 content
	decodedContent, err := base64.StdEncoding.DecodeString(fileContent.Content)
	if err != nil {
		return fmt.Errorf("failed to decode base64 content for %s: %w", filePath, err)
	}

	// Write the decoded content to the destination file
	err = os.WriteFile(destPath, decodedContent, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", destPath, err)
	}

	return nil
}
