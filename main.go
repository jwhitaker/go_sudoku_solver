package main

import (
	"os"

	"github.com/jwhitaker/go_sudoku_solver/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
