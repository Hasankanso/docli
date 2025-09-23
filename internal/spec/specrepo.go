package spec

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lucsky/cuid"
)

// DocMetaData represents a single document configuration
type DocMetaData struct {
	ID          string   "json:\"id,omitempty\""
	Name        string   "json:\"name\""
	Description string   "json:\"description,omitempty\""
	FileHints   []string "json:\"file_hints,omitempty\""
}

func NewDocMetaData(name, description string, fileHints []string) *DocMetaData {
	return &DocMetaData{
		ID:          cuid.Slug(),
		Name:        name,
		Description: description,
		FileHints:   fileHints,
	}
}

type DocSpec struct {
	Platforms []string      "json:\"platforms,omitempty\""
	DocMeta   []DocMetaData "json:\"docmeta,omitempty\""
}

type SpecRepo struct {
	SpecFilePath     string
	SpecJsonFilePath string
}

func NewSpecRepo() *SpecRepo {
	return &SpecRepo{
		SpecFilePath:     ".docs/spec.md",
		SpecJsonFilePath: ".docs/spec.json",
	}
}

func (r *SpecRepo) InitSpec(platforms []string) error {
	if r.SpecExists() {
		return fmt.Errorf("documentation configuration already exists at %s", r.SpecFilePath)
	}
	emptySpec := &DocSpec{
		Platforms: platforms,
		DocMeta:   []DocMetaData{},
	}

	err := r.Save(emptySpec)
	if err != nil {
		return err
	}
	return nil
}

func (r *SpecRepo) SpecExists() bool {
	if _, err := os.Stat(r.SpecJsonFilePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func (r *SpecRepo) AddDocMeta(doc *DocMetaData) error {
	spec, err := r.loadJsonSpec()
	if err != nil {
		return err
	}

	spec.DocMeta = append(spec.DocMeta, *doc)

	return r.Save(spec)
}

func (r *SpecRepo) RemoveDocMeta(id string) error {
	spec, err := r.loadJsonSpec()
	if err != nil {
		return err
	}
	for i, doc := range spec.DocMeta {
		if doc.ID == id {
			spec.DocMeta = append(spec.DocMeta[:i], spec.DocMeta[i+1:]...)
			r.Save(spec)
			return nil
		}
	}

	return fmt.Errorf("document's Meta data with ID '%s' not found", id)
}

func (r *SpecRepo) GetAllDocMeta() ([]DocMetaData, error) {
	spec, err := r.loadJsonSpec()
	if err != nil {
		return nil, err
	}
	return spec.DocMeta, nil
}

func (r *SpecRepo) AddPlatform(platform string) error {
	spec, err := r.loadJsonSpec()
	if err != nil {
		return err
	}

	spec.Platforms = append(spec.Platforms, platform)
	return r.Save(spec)
}

func (r *SpecRepo) RemovePlatform(platform string) error {
	spec, err := r.loadJsonSpec()
	if err != nil {
		return err
	}

	for i, p := range spec.Platforms {
		if p == platform {
			spec.Platforms = append(spec.Platforms[:i], spec.Platforms[i+1:]...)
			return r.Save(spec)
		}
	}

	return fmt.Errorf("platform '%s' not found", platform)
}

func (r *SpecRepo) Save(config *DocSpec) error {

	err := r.saveJsonSpec(config)
	if err != nil {
		return fmt.Errorf("failed to save JSON spec: %w", err)
	}

	docsDir := ".docs"
	err = os.MkdirAll(docsDir, 0755)
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

func generateSpecContent(config *DocSpec) string {
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
	if len(config.DocMeta) == 0 {
		builder.WriteString("No documents configured yet, nothing to do.\n\n")
	} else {
		for i, doc := range config.DocMeta {
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

	return builder.String()
}

func (r *SpecRepo) loadJsonSpec() (*DocSpec, error) {

	content, err := os.ReadFile(r.SpecJsonFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read spec.json: %w", err)
	}
	spec := &DocSpec{}
	err = json.Unmarshal(content, spec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec.json: %w", err)
	}
	return spec, nil
}

func (r *SpecRepo) saveJsonSpec(spec *DocSpec) error {
	content, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal spec.json: %w", err)
	}
	err = os.WriteFile(r.SpecJsonFilePath, content, 0644)
	if err != nil {
		return fmt.Errorf("failed to write spec.json: %w", err)
	}
	return nil
}
