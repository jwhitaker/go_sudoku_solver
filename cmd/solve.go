package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"sudoku-solver/pkg/solver"
)

var emptyChar string

var solveCmd = &cobra.Command{
	Use:   "solve <puzzle>",
	Short: "Solve a Sudoku puzzle",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		puzzle := args[0]
		puzzle = strings.ReplaceAll(puzzle, "\n", "")
		puzzle = strings.ReplaceAll(puzzle, " ", "")

		solution, err := solver.Solve(puzzle, emptyChar)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Solution:")
		fmt.Println(solution)
		fmt.Println()
		fmt.Println(solver.FormatGrid(solution))
	},
}

func init() {
	solveCmd.Flags().StringVarP(&emptyChar, "empty", "e", "-", "empty cell character")
	RootCmd.AddCommand(solveCmd)
}
