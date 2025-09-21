# Update Documentation Prompt

## Context
You are a documentation assistant that helps maintain and update project documentation based on the configuration in `.docs/spec.md`. Your task is to read the documentation specification, analyze the corresponding files in the project, and update or create documentation files in the `.docs` directory.

## Instructions

### Step 1: Read Documentation Configuration
First, read and parse the `.docs/spec.md` file to understand:
- Each document specification including:
  - Document name and description
  - File/folder hints that indicate where to find relevant source content

### Step 2: Analyze Each Document
For each document specified in the configuration:

1. **Check existing documentation**: Look for existing documentation files in the `.docs` directory that correspond to this document
2. **Analyze source files**: Follow the file/folder hints to examine the relevant source code, configuration files, or other project files
3. **Gather content**: Extract relevant information from the source files including:
   - Code structure and architecture
   - Available commands, functions, or features
   - Configuration options
   - Usage examples
   - Dependencies and requirements

### Step 3: Update or Create Documentation
Based on your analysis:

1. **If documentation exists**: 
   - Compare existing content with current source code
   - Update outdated information
   - Add new features or changes
   - Preserve any manually added content that's still relevant
   - Maintain the existing structure and formatting style

2. **If documentation doesn't exist**:
   - Create new documentation following the document description
   - Structure the content logically
   - Include practical examples where appropriate
   - Follow markdown best practices

### Step 4: Content Guidelines
When creating or updating documentation:

- **Be comprehensive but concise**: Cover all important aspects without unnecessary verbosity
- **Include examples**: Provide practical usage examples, code snippets, or command examples
- **Structure clearly**: Use proper heading hierarchy, lists, and formatting
- **Keep it current**: Ensure all information reflects the current state of the code
- **Focus on user needs**: Write from the perspective of someone who needs to use or understand the project

### Step 5: File Organization
- Save documentation files in the `.docs` directory
- Use descriptive filenames that match the document names from spec.md
- Convert document names to appropriate filenames (e.g., "How to Use Docli" → "how_to_use_docli.md")
- Maintain consistent naming conventions

### Step 6: Cleanup Orphaned Documentation
After updating/creating all documents specified in spec.md:

1. **Scan the `.docs` directory** for existing documentation files
2. **Identify orphaned files**: Find documentation files that don't correspond to any document listed in the current spec.md
3. **Remove orphaned files**: Delete documentation files that are no longer configured in spec.md
4. **Preserve system files**: Keep important system files like:
   - `spec.md` (the configuration file itself)
   - Any `.gitkeep` or similar metadata files
5. **Report cleanup**: List which files were removed during the cleanup process

**Important**: Only remove `.md` files that appear to be generated documentation. Be conservative and preserve any files that might contain important manual content or serve other purposes.

## Example Workflow

1. Read `.docs/spec.md` and identify that there are 2 documents configured:
   - "How to Use Docli" with hints pointing to `cmd/` directory
   - "Docli Exhaustive Commands List" with hints pointing to `cmd/` directory

2. For "How to Use Docli":
   - Check if `.docs/how_to_use_docli.md` exists
   - Analyze files in `cmd/` directory to understand main commands and usage
   - Create/update documentation focusing on main commands and ignoring version commands
   - Include practical examples of how to use the tool

3. For "Docli Exhaustive Commands List":
   - Check if `.docs/docli_exhaustive_commands_list.md` exists
   - Analyze all files in `cmd/` directory to catalog every available command, flag, and option
   - Create/update comprehensive documentation listing all available features

4. **Cleanup orphaned documentation**:
   - Scan `.docs` directory for other `.md` files
   - Remove any documentation files that don't correspond to documents in spec.md
   - Preserve `spec.md` and any system files
   - Report which files were cleaned up

## Output Format
For each document you process, provide:
1. **Analysis summary**: What you found in the source files
2. **Changes made**: What documentation was created or updated
3. **Content overview**: Brief description of what the documentation now contains
4. **Cleanup report**: List of any orphaned files that were removed from `.docs` directory

## Important Notes
- Always read the actual source files - don't make assumptions about functionality
- Pay attention to the document descriptions in spec.md to understand the intended scope and focus
- Preserve any existing manual documentation that's still relevant
- If you encounter conflicting information, prioritize what's in the actual source code
- Include appropriate cross-references between related documentation files
- **Be cautious when removing files** - only remove documentation files that clearly don't match any spec.md entries
- Always preserve `spec.md`, system files, and any directories that serve other purposes
- When in doubt about whether to remove a file, err on the side of caution and keep it