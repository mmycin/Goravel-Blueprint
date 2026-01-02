package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goraveltpl",
	Short: "A CLI tool to generate Goravel projects from template",
	Long: color.New(color.FgCyan, color.Bold).Sprint(`
╔═══════════════════════════════════════════════════════╗
║                                                       ║
║          🚀 Goravel Template Generator 🚀            ║
║                                                       ║
║    Generate new Goravel projects from template       ║
║                                                       ║
╚═══════════════════════════════════════════════════════╝
`),
	Version: "1.0.0",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func er(msg interface{}) {
	color.New(color.FgRed, color.Bold).Println("❌ Error:", msg)
	os.Exit(1)
}

func success(msg string) {
	color.New(color.FgGreen, color.Bold).Println("✅", msg)
}

func info(msg string) {
	color.New(color.FgCyan).Println("ℹ️ ", msg)
}

func warn(msg string) {
	color.New(color.FgYellow).Println("⚠️ ", msg)
}

