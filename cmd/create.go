package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"goraveltpl/internal/template"
)

var embeddedZip []byte

func SetEmbeddedZip(zipData []byte) {
	embeddedZip = zipData
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new Goravel project from template",
	Long: color.New(color.FgCyan).Sprint(`
Create a new Goravel project interactively.
You will be prompted to enter:
  - Project name: The name of your project folder
  - Module name: The Go module path (e.g., github.com/username/project)
`),
	Run: func(cmd *cobra.Command, args []string) {
		// Display banner
		color.New(color.FgMagenta, color.Bold).Println("\n" + strings.Repeat("═", 60))
		color.New(color.FgMagenta, color.Bold).Println("  🚀 Creating New Goravel Project")
		color.New(color.FgMagenta, color.Bold).Println(strings.Repeat("═", 60) + "\n")

		// Create scanner for reading input
		scanner := bufio.NewScanner(os.Stdin)

		// Prompt for project name
		color.New(color.FgCyan, color.Bold).Print("Enter Project name")
		color.New(color.FgWhite).Print(" > ")
		var projectName string
		if scanner.Scan() {
			projectName = strings.TrimSpace(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			er(fmt.Sprintf("Error reading input: %v", err))
		}
		if projectName == "" {
			er("Project name cannot be empty")
		}

		// Prompt for module name
		color.New(color.FgCyan, color.Bold).Print("Enter Module Name")
		color.New(color.FgWhite).Print(" > ")
		var moduleName string
		if scanner.Scan() {
			moduleName = strings.TrimSpace(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			er(fmt.Sprintf("Error reading input: %v", err))
		}
		if moduleName == "" {
			er("Module name cannot be empty")
		}

		// Validate module name format
		if !strings.Contains(moduleName, "/") {
			warn("Module name should typically be in format: github.com/username/project")
		}

		fmt.Println()

		// Create project directory
		projectDir := projectName
		if _, err := os.Stat(projectDir); err == nil {
			er(fmt.Sprintf("Directory '%s' already exists", projectDir))
		}

		info(fmt.Sprintf("Project Name: %s", projectName))
		info(fmt.Sprintf("Module Name: %s", moduleName))
		info(fmt.Sprintf("Template: Embedded in binary"))
		info(fmt.Sprintf("Output: %s", projectDir))
		fmt.Println()

		// Process template
		processor := template.NewProcessor(embeddedZip, projectDir, moduleName)
		if err := processor.Process(); err != nil {
			er(err.Error())
		}

		// Success message
		fmt.Println()
		color.New(color.FgGreen, color.Bold).Println(strings.Repeat("═", 60))
		success(fmt.Sprintf("Project '%s' created successfully!", projectName))
		color.New(color.FgCyan).Printf("\nNext steps:\n")
		color.New(color.FgWhite).Printf("  cd %s\n", projectDir)
		color.New(color.FgWhite).Printf("  go mod tidy\n")
		color.New(color.FgWhite).Printf("  go run main.go\n")
		color.New(color.FgGreen, color.Bold).Println(strings.Repeat("═", 60))
	},
}

