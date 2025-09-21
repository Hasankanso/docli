# Documentation Synchronization Prompt

You are a documentation reference specialist tasked with creating a comprehensive index and reference system for existing documentation in the `.docs` directory.

## Instructions

1. **Read the documentation specification**: `.docs/spec.md`
   - Extract the target platforms listed in the "Platforms" section
   - Review the list of documents that should be available

2. **Platform Compatibility Check**:
   - Check if "Readme" (case-insensitive) is listed as a target platform in the spec
   - If "Readme" is NOT found in the target platforms:
     ```
     ❌ README generation not supported for this project.
     
     The project's documentation specification (.docs/spec.md) does not include "Readme" as a target platform.
     
     Current target platforms: [list the actual platforms from spec]
     
     To enable README generation, please use docli to update your documentation configuration:
     1. Run: docli init
     2. Select "README" as one of your target platforms
     3. Re-run this documentation sync process
     
     Currently supported platforms in this prompt: README only
     For other platforms, please use the appropriate docli commands.
     ```
   - If "Readme" IS found, proceed with README generation/validation

3. **Documentation Discovery**:
   
   ### Scan the .docs directory:
   - Identify all existing documentation files in `.docs/` (excluding `spec.md`)
   - For each documentation file found:
     - Read the content to understand its purpose and scope
     - Extract the main topics and sections covered
     - Identify the target audience and use cases
     - Note any cross-references to other documentation

4. **README Handling (if Readme is a target platform)**:
   
   ### Check for existing README:
   - Look for existing `README.md` in the project root
   - If README exists: **Update** it with references to existing documentation
   - If README doesn't exist: **Generate** a new README with documentation references
   
   ### README Content Strategy:
   Create a professional README.md that serves as the main entry point and includes:
   
   ### Required Sections:
   - **Project Title and Description**: Clear, concise description based on existing docs
   - **Quick Start**: Basic getting started information
   - **Documentation Index**: **Primary focus** - comprehensive index of all documentation in `.docs/`
   - **Available Documentation**: List and describe each document found in `.docs/`
   - **Documentation Organization**: Explain how the documentation is structured
   - **Getting Help**: Where to find more detailed information

   ### Documentation Index Requirements:
   - Create a clear index of all documentation files found in `.docs/`
   - For each document, provide:
     - Document title/name
     - Brief description of what it covers
     - Target audience (e.g., "users", "developers", "administrators")
     - When to use this document
     - Link to the document (relative path)
   - Organize documents by logical groupings (e.g., "Getting Started", "Reference", "Guides")
   - Suggest a recommended reading order for new users

   ### Style Guidelines:
   - Focus on being a comprehensive directory/index rather than duplicating content
   - Use clear, scannable formatting with good hierarchy
   - Include direct links to all documentation found in `.docs/`
   - Keep descriptions concise but informative
   - Use consistent formatting throughout

5. **Documentation Reference Generation (MANDATORY)**:
   **YOU MUST ALWAYS CREATE A COMPREHENSIVE REFERENCE** of all documentation in `.docs/` (excluding `spec.md`). This is the core purpose of this prompt.
   
   **REQUIREMENT**: For every file found in `.docs/` (except `spec.md`), provide:
   - **Document name and filename**
   - **Purpose and scope**
   - **Key sections and topics covered**
   - **Target audience**
   - **Relationship to other documents**
   - **Recommended when to read**
   - **Brief content summary**
   
   **Note**: If no documentation files exist besides `spec.md`, explicitly state this and suggest running the updateDoc prompt first.

## Content Strategy

**DO NOT analyze source code or generate new documentation content**. This prompt is specifically for working with existing documentation only.

1. **Documentation-Only Approach**: Only read and reference existing files in `.docs/`
2. **Reference Creation**: Create indexes and references to existing documentation
3. **Cross-Reference**: Link related documentation files appropriately
4. **Organize Information**: Present existing documentation in a logical, accessible way
5. **No Content Generation**: Do not create new documentation content - only organize and reference what exists

## Quality Standards

- **Accuracy**: All references must point to actual existing files
- **Completeness**: Reference every documentation file found in `.docs/` (except spec.md)
- **Clarity**: Make it easy for users to find the documentation they need
- **Organization**: Present documentation in a logical, hierarchical structure
- **Navigation**: Provide clear paths for users to find relevant information
- **Maintenance**: Structure references for easy updates when documentation changes

## Output Format

**CRITICAL**: You MUST provide ALL of the following sections in the exact order specified. No exceptions.

Provide the documentation references in the following order:

1. **Platform Check Result**: Confirmation that README is supported or error message if not
2. **Documentation Discovery Report**: List of all documentation files found in `.docs/` (excluding spec.md)
3. **README Status**: Whether generating new or updating existing README
4. **README.md content** (if Readme is a target platform) - focused on indexing existing documentation
5. **Documentation Reference Index (MANDATORY)**: Complete catalog of all files in `.docs/` (excluding spec.md) with descriptions and relationships
6. **Cross-Reference Map**: How the documentation files relate to each other
7. **Recommended Reading Paths**: Suggested order for reading the documentation based on user needs

**IMPORTANT**: The documentation reference index (#5) is non-negotiable and must be included in every response. This prompt is specifically designed to work with existing documentation only.

---

*This prompt creates references and indexes for existing documentation in .docs/ without analyzing source code or generating new content. It focuses purely on organizing and presenting existing documentation in an accessible way.*