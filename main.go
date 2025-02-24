package main

import (
	"financetracker/cmd/cli"
	"os"
)

func main() {
	command := cli.InitializeCliCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
