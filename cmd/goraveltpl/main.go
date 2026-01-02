package main

import (
	"os"

	"github.com/mmycin/Goravel-Blueprint/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
