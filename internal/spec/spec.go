package spec

import (
	"github.com/Hasankanso/docli/internal/common"
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
		common.Error("A spec file already exists at %s", cmd.SpecRepo.SpecFilePath)
		common.Info("If you want to re-initialize, please back up your existing %s and delete it"+
			" before running 'docli init' again.", cmd.SpecRepo.SpecFilePath)
		return
	}

	// Step 1: Copy prompt files
	common.Info("Copying needed prompt files...")
	err := prompts.CopyPromptFiles()
	if err != nil {
		common.Fatal("Failed to copy prompt files: %v", err)
	}

	// Step 2: Initialize spec repository
	common.Info("Initializing documentation configuration...")
	err = cmd.SpecRepo.InitSpec(cmd.Platforms)
	if err != nil {
		common.Fatal("Failed to initialize documentation configuration: %v", err)
	}
	common.Success("Documentation configuration initialized successfully")
}
