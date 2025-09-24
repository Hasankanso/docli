package docmeta

import (
	"fmt"

	"github.com/Hasankanso/docli/internal/logger"
	"github.com/Hasankanso/docli/internal/spec"
)

type CreateDocMetaCommand struct {
	DocMeta  *spec.DocMetaData
	SpecRepo *spec.SpecRepo
}

func NewCreateDocMetaCommand(NewSpecRepo *spec.SpecRepo, newDocMeta *spec.DocMetaData) *CreateDocMetaCommand {
	return &CreateDocMetaCommand{
		SpecRepo: NewSpecRepo,
		DocMeta:  newDocMeta,
	}
}

func (cmd *CreateDocMetaCommand) Run() {

	specExists := cmd.SpecRepo.SpecExists()
	if !specExists {
		logger.Error("No documentation configuration found")
		logger.Info("Please run 'docli init' first to initialize your project")
		return
	}

	// Save the updated configuration
	err := cmd.SpecRepo.AddDocMeta(cmd.DocMeta)
	if err != nil {
		logger.Fatal("Error saving configuration: %v", err)
	}

	logger.Success("Document metadata for '%s' added successfully", cmd.DocMeta.Name)
}

type DeleteDocMetaCommand struct {
	ID       string
	SpecRepo *spec.SpecRepo
}

func NewDeleteDocMetaCommand(NewSpecRepo *spec.SpecRepo, id string) *DeleteDocMetaCommand {
	return &DeleteDocMetaCommand{
		SpecRepo: NewSpecRepo,
		ID:       id,
	}
}

func (cmd *DeleteDocMetaCommand) Run() {
	specExists := cmd.SpecRepo.SpecExists()
	if !specExists {
		logger.Error("No documentation configuration found")
		logger.Info("Please run 'docli init' first to initialize your project")
		return
	}

	// Save the updated configuration
	err := cmd.SpecRepo.RemoveDocMeta(cmd.ID)
	if err != nil {
		logger.Fatal("Error deleting document metadata: %v", err)
	}
	logger.Success("Document metadata with ID '%s' deleted successfully", cmd.ID)
}

type ListDocMetaCommand struct {
	SpecRepo *spec.SpecRepo
}

func NewListDocMetaCommand(NewSpecRepo *spec.SpecRepo) *ListDocMetaCommand {
	return &ListDocMetaCommand{
		SpecRepo: NewSpecRepo,
	}
}
func (cmd *ListDocMetaCommand) Run() {
	logger.Info("Listing document metadata entries")
	specExists := cmd.SpecRepo.SpecExists()
	if !specExists {
		logger.Error("No documentation configuration found")
		logger.Info("Please run 'docli init' first to initialize your project")
		return
	}

	// List the document metadata entries
	docMetaList, err := cmd.SpecRepo.GetAllDocMeta()
	if err != nil {
		logger.Fatal("Error retrieving document metadata: %v", err)
	}

	if len(docMetaList) == 0 {
		logger.Info("No document metadata entries found")
		return
	}

	// Print header with tabs
	fmt.Printf("ID\t\tName\t\t\n")
	fmt.Printf("--\t\t----\t\t\n")

	for _, docMeta := range docMetaList {

		fmt.Printf("%s\t%s\t\t\n", docMeta.ID, docMeta.Name)
	}
}
