package main

import (
	"os"

	"github.com/james-d-elliott/gocopyfw/internal/command"
)

func main() {
	if err := command.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
