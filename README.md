# Goravel Template CLI

A beautiful CLI tool to generate Goravel projects from a template with custom module names.

## Features

-   🎨 Beautiful colored output with emojis
-   🚀 Quick project generation from template
-   🌐 **Downloads template** - Always fetches the latest template from GitHub
-   🔄 **Dynamic module detection** - Automatically detects and replaces the old module name
-   📁 Preserves directory structure and `go.mod` dependencies
-   ✨ User-friendly interface

## Installation

To install the latest version:

```bash
go install github.com/mmycin/Goravel-Blueprint/cmd/goraveltpl@latest
```

## Usage

Simply run the `goraveltpl` command and you'll be prompted interactively:

```bash
goraveltpl new
```

The CLI will prompt you for:

-   **Project name**: The name of your project folder
-   **Module Name**: The Go module path (e.g., `github.com/username/my-project`)

## Requirements

-   Go 1.21 or higher (for building)
-   Internet connection (to download the template)

## What it does

1. Creates a new project folder with the specified project name
2. Downloads the latest `goravel-template.zip` from GitHub
3. Extracts the template archive into the new project
4. Detects the module name used in the template's `go.mod`
5. Replaces all occurrences of the old module name with your new module name in all files
6. Preserves the original `go.mod` (with updated module name) and its dependencies
7. Preserves file permissions and directory structure

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
ℹ️  Template: https://github.com/mmycin/Goravel-Blueprint/releases/download/Template/goravel-template.zip
ℹ️  Output: myapp

ℹ️  Downloading template...
✅ Template downloaded successfully
ℹ️  Detected template module: github.com/mmycin/goravel_test

  ✓ Created file: go.mod
  ✓ Created file: main.go
  ...

════════════════════════════════════════════════════════
✅ Project 'myapp' created successfully!

Next steps:
  cd myapp
  go mod tidy
  go run .
════════════════════════════════════════════════════════
```
