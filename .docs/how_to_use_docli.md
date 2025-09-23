# How to Use Docli

Docli is a command-line tool for managing documentation projects with configuration-driven synchronization. This guide focuses on the main commands you'll use to configure and manage your documentation projects for future synchronization with platforms like Confluence and README.

## Overview

Docli helps you:
- Set up documentation synchronization configuration
- Create and manage document metadata entries
- Organize documentation project structure  
- Define target platforms for documentation deployment
- Manage documentation project workflows
- **Provide AI-ready prompts for documentation generation** - Docli includes specialized prompt templates that guide AI assistants in creating and maintaining your documentation

## Documentation Prompts

Docli provides specialized AI-ready prompt templates that are automatically downloaded from the official repository during initialization. These prompts are copied to your project's `.github/prompts/` directory:

### Available Prompts

- **updateDoc.prompt.md** - Comprehensive documentation generation that analyzes your source code and creates complete documentation based on your configuration
- **syncDoc.prompt.md** - Creates documentation indexes and README files that reference your existing documentation  
- **refineDoc.prompt.md** - Quick, targeted refinements to existing documentation based on user feedback without code analysis

### How Prompts Work

These prompt files contain detailed instructions that help AI assistants understand:
- Your project structure and documentation goals
- How to analyze source code and extract relevant information
- What format and style to use for documentation
- How to organize and cross-reference multiple documents
- How to maintain consistency across documentation updates

### Using the Prompts

After running `docli init`, you can use these prompts with AI assistants to:
1. Generate initial documentation from your configured specifications
2. Keep documentation synchronized with code changes
3. Create comprehensive README files and documentation indexes
4. Make quick improvements and refinements to existing docs

The prompts work with your `.docs/spec.md` configuration file to provide context-aware documentation assistance.

## Main Commands

### Initialize a Documentation Project

The `init` command is your starting point for setting up documentation synchronization in your project.

```bash
docli init
```

This interactive command will:
1. Download and copy specialized documentation prompt files to `.github/prompts/` directory from the official repository
2. Check if a documentation configuration already exists at `.docs/spec.json` (exits if found)
3. Guide you through selecting target platforms (Confluence, README)
4. Create initial configuration structure in both `.docs/spec.md` and `.docs/spec.json`
5. Set up the foundation for adding document metadata

#### Platform Selection

You can choose from:
- **Confluence** - For team wikis and collaborative documentation
- **README** - For repository documentation

Select platforms by entering numbers separated by commas (e.g., `1,2` for both). If no input is provided, Confluence is selected by default.

#### Initial Setup

The init command creates the basic structure but doesn't collect document details during initialization. Use `docli create docmeta` after initialization to add your document configurations.

### Create Document Metadata

After initializing your project, use the `create docmeta` command to add new document configurations:

```bash
docli create docmeta
```

This interactive command will:
1. Prompt for document title (required)
2. Ask for document description
3. Collect file/folder source hints
4. Generate a unique ID for the document
5. Save the configuration to both `.docs/spec.md` and `.docs/spec.json`

#### Example Interactive Session

```
--- New Document Configuration ---
Enter document title: API Documentation
Enter description for 'API Documentation': Complete API reference with examples

Please provide file or folder names where we can find relevant content for 'API Documentation'.
You can specify multiple files/folders. Press Enter on an empty line when done.
  File/Folder 1 (or press Enter to finish): src/api/
  File/Folder 2 (or press Enter to finish): docs/api/
  File/Folder 3 (or press Enter to finish): (press Enter to finish)

Added 2 file/folder hint(s) for 'API Documentation'
SUCCESS: Document metadata for 'API Documentation' added successfully
```

### List Document Metadata

View all configured documents in your project:

```bash
docli list docmeta
```

This command displays all document metadata entries in a tabular format showing:
- Document ID (unique identifier)
- Document name
- Description and file hints are visible in the detailed spec.md file

### Delete Document Metadata

Remove a document configuration by its ID:

```bash
docli delete docmeta <document-id>
```

Examples:
```bash
# Delete by document ID
docli delete docmeta cjld2cyuq0000t3rmniod1foy

# If document name contains spaces, use quotes
docli delete docmeta "cjld2cyuq0000t3rmniod1foy"
```

To find the document ID, use `docli list docmeta` first.

### Getting Help

Display the welcome message and available options:

```bash
docli
```

This shows the welcome message: "Welcome to docli! Use --help to see available commands."

View detailed help for any command:

```bash
docli --help
docli init --help
```

## Global Flags

Docli supports global flags that work with all commands:

- `--verbose, -v` - Enable verbose output for detailed information
- `--quiet, -q` - Enable quiet mode to suppress non-essential output
- `--help, -h` - Show help information

### Examples

```bash
# Run init with verbose output
docli init --verbose

# Run init in quiet mode
docli init --quiet
```

## Configuration Files

After running `docli init`, your configuration is stored in two files:

- **`.docs/spec.md`** - Human-readable Markdown format for documentation and review
- **`.docs/spec.json`** - Machine-readable JSON format for internal operations

Both files contain the same information: target platforms and document metadata entries with unique IDs.

### Sample Configuration

**.docs/spec.md format:**
```markdown
# Documentation Configuration

This file contains the configuration for your documentation synchronization.

## Platforms

**Target Platforms:**
- Confluence
- Readme

## Documents

### 1. API Documentation

**Description:** Complete API reference with examples

**File/Folder Sources:**
- `src/api/`
- `docs/api/`

### 2. User Guide

**Description:** Step-by-step user instructions

**File/Folder Sources:**
- `examples/`
- `README.md`
```

**.docs/spec.json format:**
```json
{
  "platforms": ["confluence", "readme"],
  "docmeta": [
    {
      "id": "cjld2cyuq0000t3rmniod1foy",
      "name": "API Documentation", 
      "description": "Complete API reference with examples",
      "file_hints": ["src/api/", "docs/api/"]
    },
    {
      "id": "cjld2cyuq0001t3rmniod1foz",
      "name": "User Guide",
      "description": "Step-by-step user instructions", 
      "file_hints": ["examples/", "README.md"]
    }
  ]
}
```

## Best Practices

1. **Start with init** - Always begin by running `docli init` to set up your documentation project structure
2. **Add documents incrementally** - Use `docli create docmeta` to add document configurations one at a time
3. **Use descriptive names** - Give your documents clear, descriptive titles and descriptions
4. **Organize sources** - Provide clear file/folder hints to help docli find relevant content
5. **Review regularly** - Use `docli list docmeta` to review your document configurations
6. **Clean up when needed** - Remove outdated document metadata with `docli delete docmeta`
7. **Check configuration** - Review the generated `.docs/spec.md` file to ensure it matches your needs

## Error Handling

If you encounter issues:

- **Configuration exists** - If `docli init` reports that documentation configuration already exists, you'll see a message with the path to the existing `.docs/spec.json` file. To start fresh, remove the file or entire `.docs` directory first
- **Prompt file errors** - If you see errors about downloading prompt files, check your internet connection; the init process will continue but prompt files won't be available
- **Permission errors** - Ensure you have write permissions in your project directory for both `.docs` and `.github/prompts` directories  
- **No configuration found** - Commands like `create docmeta`, `list docmeta`, and `delete docmeta` require an existing configuration. Run `docli init` first
- **Document not found** - When deleting document metadata, ensure you're using the correct document ID from `docli list docmeta`

## Command Reference

| Command | Purpose |
|---------|---------|
| `docli` | Show welcome message |
| `docli init` | Initialize documentation project |
| `docli create docmeta` | Add new document metadata |
| `docli list docmeta` | List all document metadata |
| `docli delete docmeta <id>` | Remove document metadata by ID |
| `docli version` | Show version information |

## Next Steps

After setting up your documentation project with docli:

1. **Initialize**: Run `docli init` to set up the basic structure
2. **Configure**: Add document metadata using `docli create docmeta`
3. **Review**: Check your configuration with `docli list docmeta` and by examining `.docs/spec.md`
4. **Use prompts**: Leverage the AI-ready prompt files in `.github/prompts/` for documentation generation
5. **Iterate**: Add, modify, or remove document configurations as your project evolves