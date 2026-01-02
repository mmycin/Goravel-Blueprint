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

type Processor struct {
	zipData       []byte
	projectDir    string
	moduleName    string
	oldModuleName string
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

	// Detect old module name from go.mod in the zip
	if err := p.detectOldModuleName(zipReader); err != nil {
		return fmt.Errorf("failed to detect old module name: %w", err)
	}
	color.New(color.FgCyan).Printf("ℹ️  Detected template module: %s\n", p.oldModuleName)

	// Determine if there's a common top-level directory
	topLevelDir := ""
	if len(zipReader.File) > 0 {
		firstFile := zipReader.File[0].Name
		if strings.Contains(firstFile, "/") {
			parts := strings.Split(firstFile, "/")
			potentialDir := parts[0]
			isTopLevel := true
			for _, file := range zipReader.File {
				if !strings.HasPrefix(file.Name, potentialDir+"/") && file.Name != potentialDir+"/" {
					isTopLevel = false
					break
				}
			}
			if isTopLevel {
				topLevelDir = potentialDir + "/"
			}
		}
	}

	// Extract and process each file in the zip
	for _, file := range zipReader.File {
		if err := p.extractAndProcessFile(file, topLevelDir); err != nil {
			return fmt.Errorf("failed to extract file %s: %w", file.Name, err)
		}
	}

	return nil
}

func (p *Processor) detectOldModuleName(zipReader *zip.Reader) error {
	for _, file := range zipReader.File {
		if strings.HasSuffix(file.Name, "go.mod") {
			rc, err := file.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			content, err := io.ReadAll(rc)
			if err != nil {
				return err
			}

			lines := strings.Split(string(content), "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "module ") {
					p.oldModuleName = strings.TrimSpace(strings.TrimPrefix(line, "module "))
					return nil
				}
			}
		}
	}
	return fmt.Errorf("go.mod not found in template")
}

func (p *Processor) extractAndProcessFile(file *zip.File, topLevelDir string) error {
	// Strip top-level directory if it exists
	relPath := file.Name
	if topLevelDir != "" {
		if relPath == topLevelDir {
			return nil // Skip the top-level directory itself
		}
		relPath = strings.TrimPrefix(relPath, topLevelDir)
	}

	// Skip directories (they end with /)
	if file.FileInfo().IsDir() {
		destPath := filepath.Join(p.projectDir, relPath)
		if err := os.MkdirAll(destPath, 0755); err != nil {
			return err
		}
		// color.New(color.FgBlue).Printf("  📁 Created directory: %s\n", relPath)
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
	if p.isTextFile(relPath) && p.oldModuleName != "" {
		contentStr = strings.ReplaceAll(contentStr, p.oldModuleName, p.moduleName)
	}

	// Write file
	if err := os.WriteFile(destPath, []byte(contentStr), mode); err != nil {
		return err
	}

	// Only log significant files or keep it quiet to avoid spam
	// color.New(color.FgGreen).Printf("  ✓ Created file: %s\n", relPath)
	return nil
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
