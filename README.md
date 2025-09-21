# Docli

A command-line tool for generating, managing, and working with documentation in various formats.

## Quick Start

```bash
# Initialize documentation configuration for your project
docli init

# View all available commands
docli --help

# Check version
docli version
```

## Documentation Index

This project includes comprehensive documentation to help you get started and master docli. All documentation is organized in the `.docs/` directory.

### 📚 Available Documentation

#### Getting Started
- **[How to Use Docli](.docs/how_to_use_docli.md)** 
  - **Purpose:** Main usage guide focusing on core docli commands
  - **Target Audience:** New users and daily users
  - **When to use:** Start here for getting up and running with docli
  - **Covers:** Interactive `init` command, configuration setup, platform selection, best practices

#### Reference Documentation
- **[Docli Exhaustive Commands List](.docs/docli_exhaustive_commands_list.md)**
  - **Purpose:** Complete technical reference of all commands, flags, and options
  - **Target Audience:** Advanced users and developers
  - **When to use:** When you need detailed command syntax, parameters, or troubleshooting
  - **Covers:** Full command hierarchy, global flags, data structures, error codes, file operations

### 📖 Recommended Reading Order

**For New Users:**
1. Start with [How to Use Docli](.docs/how_to_use_docli.md) to understand basic concepts and main workflows
2. Reference [Docli Exhaustive Commands List](.docs/docli_exhaustive_commands_list.md) when you need specific command details

**For Advanced Users:**
- Jump directly to [Docli Exhaustive Commands List](.docs/docli_exhaustive_commands_list.md) for complete command reference
- Use [How to Use Docli](.docs/how_to_use_docli.md) for workflow examples and best practices

### 🏗️ Documentation Organization

The documentation is structured to support different user needs:

- **`.docs/how_to_use_docli.md`** - Workflow-focused, example-driven guide
- **`.docs/docli_exhaustive_commands_list.md`** - Complete technical specification
- **`.docs/spec.md`** - Project documentation configuration (auto-generated)

## Core Features

- **Interactive Setup** - Guided configuration through `docli init`
- **Multi-Platform Support** - Generate documentation for Confluence and README formats
- **Source Integration** - Configure documents with file/folder source hints
- **Flexible Configuration** - Customize documentation structure via `.docs/spec.md`

## Getting Help

- **Quick Help:** `docli --help` for command overview
- **Detailed Guides:** See [documentation index](#-available-documentation) above
- **Command Reference:** Check [Docli Exhaustive Commands List](.docs/docli_exhaustive_commands_list.md)
- **Troubleshooting:** Error handling covered in [How to Use Docli](.docs/how_to_use_docli.md#error-handling)

## What's Next?

1. **Initialize your project:** Run `docli init` to set up documentation synchronization
2. **Explore the guides:** Read through the [available documentation](#-available-documentation)
3. **Configure your docs:** Customize `.docs/spec.md` after initialization
4. **Sync your documentation:** Use docli commands to keep your documentation current

---

*For complete command reference and advanced usage, see the comprehensive documentation in the `.docs/` directory.*