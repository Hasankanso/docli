package spec

import (
	"github.com/Hasankanso/docli/internal/logger"
	"github.com/Hasankanso/docli/internal/prompts"
)

type InitSpecCommand struct {
	SpecRepo  *SpecRepo
	Platforms []string
}

func NewInitSpecCommand(NewSpecRepo *SpecRepo, platforms []string) *InitSpecCommand {
	return &InitSpecCommand{
		SpecRepo:  NewSpecRepo,
		Platforms: platforms,
	}
}

func (cmd *InitSpecCommand) Run() {

	// Check if spec already exists
	if cmd.SpecRepo.SpecExists() {
		logger.Error("A spec file already exists at %s", cmd.SpecRepo.SpecFilePath)
		logger.Info("If you want to re-initialize, please back up your existing %s and delete it"+
			" before running 'docli init' again.", cmd.SpecRepo.SpecFilePath)
		return
	}

	// Step 1: Copy prompt files
	logger.Info("Copying needed prompt files...")
	err := prompts.CopyPromptFiles()
	if err != nil {
		logger.Fatal("Failed to copy prompt files: %v", err)
	}

	// Step 2: Initialize spec repository
	logger.Info("Initializing documentation configuration...")
	err = cmd.SpecRepo.InitSpec(cmd.Platforms)
	if err != nil {
		logger.Fatal("Failed to initialize documentation configuration: %v", err)
	}
	logger.Success("Documentation configuration initialized successfully")
}
