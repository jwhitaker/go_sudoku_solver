package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jwhitaker/go_sudoku_solver/pkg/solver"
)

var isValidCmd = &cobra.Command{
	Use:   "isvalid [puzzle]",
	Short: "Check if a Sudoku puzzle is valid",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		var puzzle string
		if len(args) == 0 {
			reader := bufio.NewReader(os.Stdin)
			data, _ := io.ReadAll(reader)
			puzzle = string(data)
		} else {
			puzzle = args[0]
		}
		puzzle = strings.ReplaceAll(puzzle, "\n", "")
		puzzle = strings.ReplaceAll(puzzle, " ", "")

		valid := solver.IsValid(puzzle, emptyChar)
		if valid {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	},
}

func init() {
	isValidCmd.Flags().StringVarP(&emptyChar, "empty", "e", "-", "empty cell character")
	RootCmd.AddCommand(isValidCmd)
}
