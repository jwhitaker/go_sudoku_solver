package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jwhitaker/go_sudoku_solver/pkg/solver"
)

var difficulty string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a Sudoku puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		puzzle, err := solver.Generate(difficulty)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Puzzle:")
		fmt.Println(puzzle.String("-"))
		fmt.Println()
		fmt.Println("Solution:")
		fmt.Println(puzzle.SolvedString())
	},
}

func init() {
	generateCmd.Flags().StringVarP(&difficulty, "difficulty", "d", "easy", "difficulty level (easy, medium, hard)")
	RootCmd.AddCommand(generateCmd)
}
