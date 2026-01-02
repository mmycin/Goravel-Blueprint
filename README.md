# Goravel Template CLI

A beautiful CLI tool to generate Goravel projects from a template with custom module names.

## Features

- 🎨 Beautiful colored output with emojis
- 🚀 Quick project generation from template
- 🔄 Automatic module name replacement
- 📁 Preserves directory structure
- ✨ User-friendly interface
- 📦 **Standalone binary** - Template is embedded, no external files needed

## Installation

```bash
go build -o goraveltpl.exe
```

Or install globally:
```bash
go install
```

## Usage

Simply run the `new` command and you'll be prompted interactively:

```bash
goraveltpl new
```

The CLI will prompt you for:
- **Project name**: The name of your project folder
- **Module Name**: The Go module path (e.g., `github.com/username/my-project`)

## Requirements

- Go 1.21 or higher (for building)
- The compiled binary is **standalone** - no external template files needed at runtime

## What it does

1. Creates a new project folder with the specified project name
2. Extracts the embedded `repo.zip` template archive into the new project
3. Replaces all occurrences of `github.com/mmycin/goravel-test` with your module name in all files
4. Generates `go.mod` with your module name
5. Preserves file permissions and directory structure

**Note:** The template is embedded as `repo.zip` in the binary during compilation, so the compiled `goraveltpl` binary is completely standalone and doesn't require any external files when running.

## Example

```bash
$ goraveltpl new

════════════════════════════════════════════════════════
  🚀 Creating New Goravel Project
════════════════════════════════════════════════════════

Enter Project name > myapp
Enter Module Name > github.com/john/myapp

ℹ️  Project Name: myapp
ℹ️  Module Name: github.com/john/myapp
ℹ️  Template: Embedded in binary
ℹ️  Output: myapp

  📁 Created directory: app
  📁 Created directory: app/http
  ✓ Created file: .gitkeep
  ...

════════════════════════════════════════════════════════
✅ Project 'myapp' created successfully!

Next steps:
  cd myapp
  go mod tidy
  go run main.go
════════════════════════════════════════════════════════
```

