package main

import (
	_ "embed"
	"os"

	"goraveltpl/cmd"
)

//go:embed repo.zip
var embeddedZip []byte

func main() {
	// Pass embedded zip to cmd package
	cmd.SetEmbeddedZip(embeddedZip)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
