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

var displayCmd = &cobra.Command{
	Use:   "display [puzzle]",
	Short: "Display a Sudoku puzzle",
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

		fmt.Println(solver.FormatGrid(puzzle, emptyChar))
	},
}

func init() {
	displayCmd.Flags().StringVarP(&emptyChar, "empty", "e", "-", "empty cell character")
	RootCmd.AddCommand(displayCmd)
}
