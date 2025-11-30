package main

import (
	"fmt"
	"os"

	"github.com/helton-godoy/shantilly-runtime/internal/app"
)

const (
	// Version information
	Version   = "1.0.0-alpha"
	Commit    = "dev"
	BuildDate = "unknown"
)

func main() {
	if err := app.Run(Version, Commit, BuildDate); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
