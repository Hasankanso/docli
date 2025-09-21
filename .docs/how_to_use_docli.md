# How to Use Docli

Docli is a command-line tool for setting up documentation synchronization configuration. This guide focuses on the main commands you'll use to configure your documentation projects for future synchronization with platforms like Confluence and README.

## Overview

Docli helps you:
- Set up documentation synchronization configuration
- Organize documentation project structure
- Define target platforms for documentation deployment
- Manage documentation project workflows
- **Provide AI-ready prompts for documentation generation** - Docli includes specialized prompt templates that guide AI assistants in creating and maintaining your documentation

## Documentation Prompts

Docli provides specialized AI-ready prompt templates that guide documentation generation and maintenance. These prompts are automatically copied to your project's `.github/prompts/` directory during initialization:

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
1. Copy specialized documentation prompt files to `.github/prompts/` directory (updateDoc.prompt.md, syncDoc.prompt.md, and refineDoc.prompt.md) - these provide AI assistants with detailed instructions for generating, updating, and refining your documentation
2. Check if a documentation configuration already exists at `.docs/spec.md` (exits if found)
3. Guide you through selecting target platforms (Confluence, README)
4. Help you configure documents with names, descriptions, and file/folder sources
5. Create a `.docs/spec.md` file with your configuration

#### Platform Selection

You can choose from:
- **Confluence** - For team wikis and collaborative documentation
- **README** - For repository documentation

Select platforms by entering numbers separated by commas (e.g., `1,2` for both). If no input is provided, Confluence is selected by default.

#### Document Configuration

For each document, you'll provide:
- **Title** - The name of your document
- **Description** - What the document covers
- **File/Folder Sources** - Where docli should look for relevant content

#### Example Interactive Session

```
Welcome to docli initialization!
This will guide you through setting up your documentation sync configuration.
📋 Copying specialized documentation prompt files...
✅ Documentation prompt files copied successfully to .github/prompts/!
📋 Which platforms do you want to sync your documentation to?
1. Confluence
2. README

Select platforms (1): 1,2
Selected: confluence, readme

📝 Now let's configure your documents:
--- Document 1 ---
Enter title for document 1 (or press Enter to finish): API Documentation
Enter description for 'API Documentation': Complete API reference with examples

💡 Please provide file or folder names where we can find relevant content for 'API Documentation'.
You can specify multiple files/folders. Press Enter on an empty line when done.
  File/Folder 1 (or press Enter to finish): src/api/
  File/Folder 2 (or press Enter to finish): docs/api/
  File/Folder 3 (or press Enter to finish): (press Enter to finish)

--- Document 2 ---
Enter title for document 2 (or press Enter to finish): (press Enter to finish)

✅ Configuration saved to .docs/spec.md

You can now run other docli commands to sync your documentation!
```

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

## Configuration File

After running `docli init`, your configuration is stored in `.docs/spec.md`. This file contains:

- **Platforms** - Target platforms for documentation sync
- **Documents** - Each document's name, description, and source file hints

### Sample Configuration

```markdown
# Documentation Configuration

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

## Best Practices

1. **Start with init** - Always begin by running `docli init` to set up your documentation project
2. **Organize sources** - Provide clear file/folder hints to help docli find relevant content
3. **Use descriptive names** - Give your documents clear, descriptive titles and descriptions
4. **Check configuration** - Review the generated `.docs/spec.md` file to ensure it matches your needs

## Error Handling

If you encounter issues:

- **Configuration exists** - If `docli init` reports that documentation configuration already exists, you'll see a message with the path to the existing `.docs/spec.md` file. To start fresh, remove the file or entire `.docs` directory first
- **Prompt file warnings** - If you see warnings about failing to copy prompt files, they're non-critical; the init process will continue
- **Permission errors** - Ensure you have write permissions in your project directory for both `.docs` and `.github/prompts` directories  
- **No input provided** - Docli will use sensible defaults (e.g., Confluence as default platform, empty string handling)

## Next Steps

After initializing your documentation project with `docli init`, you can:
1. Review and edit the generated `.docs/spec.md` file
2. Run additional docli commands to sync your documentation to the configured platforms
3. Update your source files and re-run docli commands to keep documentation current