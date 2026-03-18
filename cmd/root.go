package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "sudoku-solver",
	Short: "A Sudoku solver using Dancing Links",
}
