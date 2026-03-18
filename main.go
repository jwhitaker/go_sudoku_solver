package main

import (
	"os"
	"sudoku-solver/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
