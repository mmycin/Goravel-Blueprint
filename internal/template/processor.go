package template

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

const oldModuleName = "github.com/mmycin/goravel-test"

type Processor struct {
	zipData    []byte
	projectDir string
	moduleName string
}

func NewProcessor(zipData []byte, projectDir, moduleName string) *Processor {
	return &Processor{
		zipData:    zipData,
		projectDir: projectDir,
		moduleName: moduleName,
	}
}

func (p *Processor) Process() error {
	// Create project directory
	if err := os.MkdirAll(p.projectDir, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Open zip archive
	zipReader, err := zip.NewReader(bytes.NewReader(p.zipData), int64(len(p.zipData)))
	if err != nil {
		return fmt.Errorf("failed to open zip archive: %w", err)
	}

	// Extract and process each file in the zip
	for _, file := range zipReader.File {
		if err := p.extractAndProcessFile(file); err != nil {
			return fmt.Errorf("failed to extract file %s: %w", file.Name, err)
		}
	}

	// Generate go.mod file
	if err := p.generateGoMod(); err != nil {
		return fmt.Errorf("failed to generate go.mod: %w", err)
	}
	color.New(color.FgGreen).Printf("  ✓ Created file: go.mod\n")

	return nil
}

func (p *Processor) extractAndProcessFile(file *zip.File) error {
	// Skip directories (they end with /)
	if file.FileInfo().IsDir() {
		relPath := file.Name
		destPath := filepath.Join(p.projectDir, relPath)
		if err := os.MkdirAll(destPath, 0755); err != nil {
			return err
		}
		color.New(color.FgBlue).Printf("  📁 Created directory: %s\n", relPath)
		return nil
	}

	// Open file from zip
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	// Read file content
	content, err := io.ReadAll(rc)
	if err != nil {
		return err
	}

	// Create destination path
	relPath := file.Name
	destPath := filepath.Join(p.projectDir, relPath)

	// Create destination directory if needed
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return err
	}

	// Determine file mode
	mode := file.FileInfo().Mode()
	if p.isExecutableFile(relPath) {
		mode = 0755
	} else {
		mode = 0644
	}

	// Replace module name if it's a text file
	contentStr := string(content)
	if p.isTextFile(relPath) {
		contentStr = strings.ReplaceAll(contentStr, oldModuleName, p.moduleName)
	}

	// Write file
	if err := os.WriteFile(destPath, []byte(contentStr), mode); err != nil {
		return err
	}

	color.New(color.FgGreen).Printf("  ✓ Created file: %s\n", relPath)
	return nil
}

func (p *Processor) generateGoMod() error {
	goModPath := filepath.Join(p.projectDir, "go.mod")
	goModContent := fmt.Sprintf("module %s\n\ngo 1.21\n\nrequire (\n\tgithub.com/goravel/framework v1.13.0\n)\n", p.moduleName)

	return os.WriteFile(goModPath, []byte(goModContent), 0644)
}


func (p *Processor) isTextFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	textExts := []string{
		".go", ".mod", ".sum", ".md", ".txt", ".json", ".yaml", ".yml",
		".toml", ".ini", ".conf", ".env", ".sh", ".bat", ".ps1",
		".html", ".css", ".js", ".ts", ".xml", ".sql",
	}
	for _, textExt := range textExts {
		if ext == textExt {
			return true
		}
	}
	// Also check if file has no extension but might be text (like Dockerfile, Makefile)
	baseName := strings.ToLower(filepath.Base(path))
	if baseName == "dockerfile" || baseName == "makefile" || strings.HasPrefix(baseName, ".env") {
		return true
	}
	return false
}

func (p *Processor) isExecutableFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	execExts := []string{".sh", ".bat", ".ps1"}
	for _, execExt := range execExts {
		if ext == execExt {
			return true
		}
	}
	// Check if file has no extension but might be executable
	baseName := strings.ToLower(filepath.Base(path))
	if baseName == "makefile" {
		return false
	}
	return false
}
