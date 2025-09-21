# How to Use Docli

Docli is a command-line tool for generating, managing, and working with documentation in various formats. This guide focuses on the main commands you'll use to get started and manage your documentation projects.

## Overview

Docli helps you:
- Generate documentation from source code
- Convert between different documentation formats
- Manage documentation workflows
- Create and maintain project documentation

## Main Commands

### Initialize a Documentation Project

The `init` command is your starting point for setting up documentation synchronization in your project.

```bash
docli init
```

This interactive command will:
1. Copy prompt files to `.github/prompts/` directory (updateDoc.prompt.md and syncDoc.prompt.md)
2. Check if a documentation configuration already exists (exits if found)
3. Guide you through selecting target platforms (Confluence, README)
4. Help you configure documents with names, descriptions, and file/folder sources
5. Create a `.docs/spec.md` file with your configuration

#### Platform Selection

You can choose from:
- **Confluence** - For team wikis and collaborative documentation
- **README** - For repository documentation

Select platforms by entering numbers separated by commas (e.g., `1,2` for both).

#### Document Configuration

For each document, you'll provide:
- **Title** - The name of your document
- **Description** - What the document covers
- **File/Folder Sources** - Where docli should look for relevant content

#### Example Interactive Session

```
Welcome to docli initialization!
This will guide you through setting up your documentation sync configuration.
📋 Copying needed prompt files...
✅ Prompt files copied successfully!
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

This shows a brief welcome message and reminds you to use `--help` for available commands.

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

- **Configuration exists** - If `docli init` reports existing configuration, remove `.docs/spec.md` or the entire `.docs` directory first
- **Prompt file warnings** - If you see warnings about prompt files, they're non-critical; the init process will continue
- **Permission errors** - Ensure you have write permissions in your project directory for both `.docs` and `.github/prompts` directories  
- **No input provided** - Docli will use sensible defaults (e.g., Confluence as default platform)

## Next Steps

After initializing your documentation project with `docli init`, you can:
1. Review and edit the generated `.docs/spec.md` file
2. Run additional docli commands to sync your documentation to the configured platforms
3. Update your source files and re-run docli commands to keep documentation current