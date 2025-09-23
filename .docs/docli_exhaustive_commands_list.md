# Docli Exhaustive Commands List

This document provides a comprehensive reference of all available commands, flags, and options in docli.

## Root Command

### `docli`

**Usage:** `docli`

**Description:** A documentation CLI tool

**Long Description:**
```
docli is a command-line tool for generating, managing, 
and working with documentation in various formats.

Use docli to:
- Generate documentation from source code
- Convert between different documentation formats
- Manage documentation workflows
- Create and maintain project documentation
```

**Behavior:** When run without subcommands, displays a welcome message: "Welcome to docli! Use --help to see available commands."

## Subcommands

### `docli init`

**Usage:** `docli init`

**Description:** Initialize basic documentation project structure

**Long Description:**
```
Initialize your documentation project by setting up the basic configuration
structure. This will copy prompt files and create the initial spec.md file
with platform configuration. Use 'docli create docmeta' to add document metadata.
```

**Functionality:**
- Downloads prompt files from GitHub repository and copies them to `.github/prompts/` directory
- Checks for existing `.docs/spec.json` configuration (exits with error if found)
- Interactive platform selection (Confluence, README)
- Creates initial configuration structure in both `.docs/spec.md` and `.docs/spec.json`
- Sets up foundation for document metadata management

**Interactive Flow:**
1. Welcome message and prompt file downloading:
   - Creates `.github/prompts/` directory if it doesn't exist
   - Downloads and copies `.prompt.md` files from official GitHub repository if they don't exist locally
   - Uses GitHub API to fetch file list and content

2. Platform selection with options:
   - `1` - Confluence
   - `2` - README
   - Multiple selections via comma-separated input (e.g., `1,2`)
   - Default: Confluence if no input provided

3. Configuration creation:
   - Creates empty document metadata structure
   - Generates both JSON and Markdown format configuration files

**Error Conditions:**
- Exits with fatal error if `.docs/spec.json` already exists
- Provides instructions to remove existing configuration
- Network errors during prompt file download are handled gracefully

**File Operations:**
- Creates `.github/prompts` directory with permissions `0755`
- Downloads prompt files from GitHub with permissions `0644`
- Creates `.docs` directory with permissions `0755`
- Writes `spec.md` and `spec.json` files with permissions `0644`

### `docli create`

**Usage:** `docli create <resource>`

**Description:** Create resources

**Long Description:**
```
Create various types of resources in your documentation project.

Available resource types:
  docmeta - Create a new document metadata entry

The create command provides subcommands to create different types of resources
that help organize and manage your project documentation.
```

**Subcommands:**
- `docmeta` - Create a new document metadata entry

### `docli create docmeta`

**Usage:** `docli create docmeta`

**Description:** Create a new document metadata entry

**Long Description:**
```
Create a new document metadata entry and add it to your spec.md file.
This command will guide you through an interactive process to define a new
document with its name, description, and file hints.
```

**Functionality:**
- Checks for existing configuration (requires `docli init` to be run first)
- Interactive collection of document details
- Generates unique CUID identifier for the document
- Saves configuration to both `.docs/spec.md` and `.docs/spec.json`

**Interactive Flow:**
1. Document title prompt (required, empty cancels operation)
2. Document description prompt (optional)
3. File/folder hints collection loop:
   - Multiple file/folder paths can be provided
   - Empty input terminates collection
   - No validation on paths provided

**Data Structure:**
```go
type DocMetaData struct {
    ID          string   // Unique CUID identifier
    Name        string   // Document title
    Description string   // Document description
    FileHints   []string // Array of file/folder paths
}
```

**Error Conditions:**
- Exits if no configuration found (requires `docli init` first)
- Document creation cancelled if empty title provided

### `docli list`

**Usage:** `docli list <resource>`

**Description:** List resources

**Long Description:**
```
List various types of resources in your documentation project.

Available resource types:
  docmeta - List all document metadata entries

Use the appropriate subcommand to list the specific type of resource you want to view.
```

**Subcommands:**
- `docmeta` - List all document metadata entries

### `docli list docmeta`

**Usage:** `docli list docmeta`

**Description:** List all document metadata entries

**Long Description:**
```
List all document metadata entries from your spec.md file.
This command displays all configured documents with their names and descriptions.
```

**Functionality:**
- Checks for existing configuration (requires `docli init` to be run first)
- Loads document metadata from `.docs/spec.json`
- Displays tabular output with ID and Name columns

**Output Format:**
```
ID                      Name
--                      ----
cjld2cyuq0000t3rmniod1foy    API Documentation
cjld2cyuq0001t3rmniod1foz    User Guide
```

**Error Conditions:**
- Exits if no configuration found (requires `docli init` first)
- Shows message if no document metadata entries found

### `docli delete`

**Usage:** `docli delete <resource> <id>`

**Description:** Delete resources

**Long Description:**
```
Delete various types of resources from your documentation project.

Available resource types:
  docmeta - Delete a document metadata entry by id

Use the appropriate subcommand to delete the specific type of resource you want to remove.
```

**Subcommands:**
- `docmeta` - Delete a document metadata entry by ID

### `docli delete docmeta`

**Usage:** `docli delete docmeta <id>`

**Description:** Delete a document metadata entry by id

**Long Description:**
```
Delete a document metadata entry from your spec.md file by providing the document id.
The id should be provided in quotes if it contains spaces.

Example:
  docli delete docmeta "cjld2cyuq0000t3rmniod1foy"
  docli delete docmeta cjld2cyuq0000t3rmniod1foy
```

**Arguments:**
- `<id>` - The unique identifier of the document metadata to delete (required)

**Functionality:**
- Checks for existing configuration (requires `docli init` to be run first)
- Removes document metadata entry by ID from internal storage
- Updates both `.docs/spec.md` and `.docs/spec.json` files

**Error Conditions:**
- Exits if no configuration found (requires `docli init` first)
- Exits with fatal error if document ID not found

### `docli version`

**Usage:** `docli version`

**Description:** Print the version number of docli

**Long Description:** Display the current version of docli.

**Output:**
```
docli v2.0.0
Built with Go
```

These flags are available for all commands through the root command's persistent flags:

### `--verbose, -v`

**Type:** Boolean  
**Default:** `false`  
**Description:** Enable verbose output  
**Usage:** `docli [command] --verbose` or `docli [command] -v`

### `--quiet, -q`

**Type:** Boolean  
**Default:** `false`  
**Description:** Enable quiet mode  
**Usage:** `docli [command] --quiet` or `docli [command] -q`

### `--help, -h`

**Type:** Boolean  
**Description:** Show help information  
**Usage:** `docli --help` or `docli [command] --help`

## Command Hierarchy

```
docli (root)
├── init
├── create
│   └── docmeta
├── list
│   └── docmeta
├── delete
│   └── docmeta
├── version
├── --verbose, -v (global flag)
├── --quiet, -q (global flag)
└── --help, -h (global flag)
```

## Data Types and Structures

### Document Metadata Structure
```go
type DocMetaData struct {
    ID          string   `json:"id,omitempty"`          // Unique CUID identifier
    Name        string   `json:"name"`                  // Document title
    Description string   `json:"description,omitempty"` // Document description
    FileHints   []string `json:"file_hints,omitempty"`  // Array of file/folder paths
}
```

### Documentation Specification Structure
```go
type DocSpec struct {
    Platforms []string      `json:"platforms,omitempty"` // Array of platform names
    DocMeta   []DocMetaData `json:"docmeta,omitempty"`   // Array of document metadata
}
```

### Platform Options
- `confluence` - Confluence platform
- `readme` - README platform

### Unique Identifier Generation
- Uses CUID (Collision-resistant Unique Identifier) for document IDs
- Generated via `github.com/lucsky/cuid` package
- Format: Short, URL-safe strings (e.g., `cjld2cyuq0000t3rmniod1foy`)

## File Operations

### Configuration File Formats

Docli maintains configuration in two synchronized formats:

#### JSON Format (`.docs/spec.json`)
```json
{
  "platforms": ["confluence", "readme"],
  "docmeta": [
    {
      "id": "cjld2cyuq0000t3rmniod1foy",
      "name": "API Documentation",
      "description": "Complete API reference with examples",
      "file_hints": ["src/api/", "docs/api/"]
    }
  ]
}
```

#### Markdown Format (`.docs/spec.md`)
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
```

### Directory Structure Created

```
.docs/
├── spec.md     # Human-readable configuration
└── spec.json   # Machine-readable configuration

.github/
└── prompts/
    ├── updateDoc.prompt.md     # Downloaded from GitHub
    ├── syncDoc.prompt.md       # Downloaded from GitHub
    └── refineDoc.prompt.md     # Downloaded from GitHub
```

### Prompt File Management

Prompt files are downloaded from the official GitHub repository:
- **Repository**: `https://github.com/Hasankanso/docli`
- **API Endpoint**: `https://api.github.com/repos/Hasankanso/docli/contents/.github/prompts`
- **Process**: 
  1. Fetch directory listing via GitHub API
  2. Filter for `.prompt.md` files
  3. Download file content via GitHub API (base64 encoded)
  4. Decode and save locally if file doesn't exist

## Exit Codes

- `0` - Success
- `1` - Error (used by commands when configuration issues occur, network errors, or file operation failures)

## Dependencies

Based on the imports in the source code:
- `github.com/spf13/cobra` - CLI framework
- `github.com/lucsky/cuid` - CUID generation for unique identifiers
- Go standard library packages:
  - `bufio` - Buffered I/O
  - `encoding/base64` - Base64 encoding/decoding
  - `encoding/json` - JSON marshaling/unmarshaling
  - `fmt` - Formatted I/O
  - `net/http` - HTTP client for GitHub API
  - `os` - Operating system interface
  - `path/filepath` - File path manipulation
  - `slices` - Slice operations
  - `strings` - String operations
  - `time` - Time operations (HTTP timeouts)
- Internal packages:
  - `github.com/Hasankanso/docli/internal/common` - Logging and utilities
  - `github.com/Hasankanso/docli/internal/docmeta` - Document metadata operations
  - `github.com/Hasankanso/docli/internal/prompts` - Prompt file management
  - `github.com/Hasankanso/docli/internal/spec` - Specification repository management

## Input Validation

### Platform Selection
- Accepts comma-separated numeric input
- Maps input: `1` → `confluence`, `2` → `readme`
- Filters duplicates using `slices.Contains`
- Defaults to `confluence` for invalid/empty input

### Document Metadata Validation
- Document names: Accepts any non-empty string; empty name cancels operation
- Document descriptions: Accepts any string including empty
- File hints: Accepts any string including empty; no path validation performed
- No character limits enforced on any field

### Document ID Validation
- Must be valid CUID format (generated internally)
- Used for delete operations
- No user input validation as IDs are system-generated

## Error Messages

### Init Command Errors

**Existing Configuration:**
```
FATAL: A spec file already exists at .docs/spec.json
If you want to re-initialize, please back up your existing .docs/spec.json and delete it before running 'docli init' again.
```

**Network/GitHub API Errors:**
```
FATAL: Failed to copy prompt files: [error description]
```

**File System Errors:**
```
FATAL: Failed to initialize documentation configuration: [error description]
```

### Document Metadata Command Errors

**No Configuration Found:**
```
ERROR: No documentation configuration found
Please run 'docli init' first to initialize your project
```

**Document Not Found (Delete):**
```
FATAL: Error deleting document metadata: document's Meta data with ID '[id]' not found
```

**File Operation Errors:**
```
FATAL: Error saving configuration: [error description]
FATAL: Error retrieving document metadata: [error description]
```

## Success Messages

### Init Command Success
```
Copying needed prompt files...
SUCCESS: Created .github/prompts/[filename]  (for each prompt file)
Initializing documentation configuration...
SUCCESS: Documentation configuration initialized successfully
```

### Document Metadata Success
```
SUCCESS: Document metadata for '[document name]' added successfully
SUCCESS: Document metadata with ID '[id]' deleted successfully
```

### List Command Output
```
Listing document metadata entries
ID                      Name
--                      ----
[document-id]          [document-name]
```

### Network Error Handling
```
FATAL: Failed to copy prompt files: network error fetching directory listing: [error]
FATAL: Failed to copy prompt files: server error fetching directory listing: HTTP [code]
FATAL: Failed to copy prompt files: failed to parse GitHub API response: [error]
```