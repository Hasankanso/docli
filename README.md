# docli

[![Go](https://img.shields.io/badge/Go-1.24.1-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A powerful command-line tool for generating, managing, and working with documentation in various formats. Streamline your documentation workflow with automated generation, platform synchronization, and intelligent content management.

## 🚀 Features

- **Interactive Setup** - Guided configuration process for documentation projects
- **Multi-Platform Sync** - Support for Confluence and README documentation
- **Source Code Analysis** - Automatically extract documentation from your codebase
- **Template Management** - Built-in prompt templates for consistent documentation
- **Flexible Configuration** - Customizable document specifications and source mappings
- **CLI-First Design** - Efficient command-line interface with comprehensive help

## 📦 Installation

### Download Binary

Download the latest release for your platform:
- [Windows](https://github.com/Hasankanso/docli/releases) - `docli.exe`
- [macOS](https://github.com/Hasankanso/docli/releases) - `docli`
- [Linux](https://github.com/Hasankanso/docli/releases) - `docli`

### Build from Source

```bash
# Clone the repository
git clone https://github.com/Hasankanso/docli.git
cd docli

# Build the binary
go build -o docli

# Run docli
./docli --help
```

## ⚡ Quick Start

### 1. Initialize Your Documentation Project

```bash
# Start the interactive setup
docli init
```

This will guide you through:
- Selecting target platforms (Confluence, README)
- Configuring your documents
- Setting up source file mappings
- Creating the configuration file at `.docs/spec.md`

### 2. Example Interactive Session

```
Welcome to docli initialization!
📋 Which platforms do you want to sync your documentation to?
1. Confluence
2. README

Select platforms (1): 1,2
Selected: confluence, readme

📝 Now let's configure your documents:
--- Document 1 ---
Enter title: API Documentation
Enter description: Complete API reference with examples
File/Folder sources: src/api/, docs/api/

✅ Configuration saved to .docs/spec.md
```

### 3. Verify Your Setup

```bash
# Check the generated configuration
cat .docs/spec.md

# View available commands
docli --help
```

## 🛠️ Usage

### Commands

#### `docli init`
Initialize documentation synchronization configuration for your project.

```bash
docli init

# With verbose output
docli init --verbose

# With quiet mode
docli init --quiet
```

**What it does:**
- Creates `.docs/spec.md` configuration file
- Copies prompt templates to `.github/prompts/`
- Sets up platform targeting (Confluence, README)
- Configures document specifications with source mappings

#### `docli version`
Display version information.

```bash
docli version
# Output: docli v1.0.0
```

#### `docli`
Show welcome message and available commands.

```bash
docli
# Output: Welcome to docli! Use --help to see available commands.
```

### Global Flags

- `--verbose, -v` - Enable detailed output
- `--quiet, -q` - Suppress non-essential messages  
- `--help, -h` - Show help information

## ⚙️ Configuration

After running `docli init`, your configuration is stored in `.docs/spec.md`:

```markdown
# Documentation Configuration

## Platforms
**Target Platforms:**
- Confluence
- Readme

## Documents

### 1. How to Use Docli
**Description:** emphasize main docli commands, ignore side commands such as version
**File/Folder Sources:**
- `cmd/ and go into related methods used in cmd, for a depth usage understanding`

### 2. Docli Exhaustive Commands List  
**Description:** list all available commands, flags and options
**File/Folder Sources:**
- `cmd/`
```

### Configuration Elements

- **Platforms** - Target destinations for your documentation (Confluence, README)
- **Documents** - Individual documentation pieces with descriptions and source hints
- **File/Folder Sources** - Where docli should look for relevant content

## 📁 Project Structure

```
docli/
├── cmd/                    # CLI command implementations
│   ├── root.go            # Root command and global configuration
│   ├── init.go            # Interactive initialization command
│   └── version.go         # Version command
├── internal/
│   └── prompts/           # Prompt template management
│       └── prompts.go     # Template copying and setup
├── .docs/                 # Generated documentation configuration
│   ├── spec.md           # Main configuration file
│   ├── how_to_use_docli.md
│   └── docli_exhaustive_commands_list.md
├── .github/
│   └── prompts/           # Documentation generation templates
├── main.go               # Application entry point
├── go.mod                # Go module definition
└── README.md             # This file
```

## 🔧 Development

### Prerequisites

- Go 1.24.1 or later
- Git

### Building

```bash
# Clone and build
git clone https://github.com/Hasankanso/docli.git
cd docli
go mod tidy
go build -o docli

# Run tests (if available)
go test ./...
```

### Dependencies

- **[Cobra](https://github.com/spf13/cobra)** - Modern CLI framework for Go
- **Go Standard Library** - File operations, I/O, and string processing

## 💡 Use Cases

### Documentation Teams
- **Standardize** documentation formats across projects
- **Automate** content generation from source code
- **Sync** documentation to multiple platforms simultaneously

### Open Source Projects  
- **Generate** comprehensive READMEs from code analysis
- **Maintain** consistent documentation structure
- **Extract** API documentation from source comments

### Enterprise Development
- **Integrate** with Confluence for team collaboration
- **Centralize** documentation configuration and templates
- **Scale** documentation processes across multiple repositories

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Development Process

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

### Getting Help

- **Documentation**: Check the generated `.docs/` files for detailed guides
- **Issues**: [GitHub Issues](https://github.com/Hasankanso/docli/issues)
- **Discussions**: [GitHub Discussions](https://github.com/Hasankanso/docli/discussions)

### Troubleshooting

**Configuration already exists:**
```bash
# Remove existing configuration to start fresh
rm .docs/spec.md
# or remove the entire .docs directory
rm -rf .docs
docli init
```

**Permission errors:**
- Ensure write permissions for `.docs/` and `.github/prompts/` directories
- Check that you're running from your project root directory

**Missing prompt files:**
- Warnings about prompt file copying are non-critical
- The init process will continue and create the necessary configuration

---

**Built with ❤️ using Go and Cobra**