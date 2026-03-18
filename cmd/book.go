package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jung-kurt/gofpdf"
	"github.com/spf13/cobra"

	"github.com/jwhitaker/go_sudoku_solver/pkg/generator"
)

var bookCmd = &cobra.Command{
	Use:   "book <filename>",
	Short: "Generate a PDF book of Sudoku puzzles",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		easyCount, _ := strconv.Atoi(easy)
		mediumCount, _ := strconv.Atoi(medium)
		hardCount, _ := strconv.Atoi(hard)

		puzzles := make([]*generator.Puzzle, 0)

		for i := 0; i < easyCount; i++ {
			puzzle, err := generator.Generate("easy")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error generating easy puzzle: %v\n", err)
				os.Exit(1)
			}
			puzzles = append(puzzles, puzzle)
		}

		for i := 0; i < mediumCount; i++ {
			puzzle, err := generator.Generate("medium")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error generating medium puzzle: %v\n", err)
				os.Exit(1)
			}
			puzzles = append(puzzles, puzzle)
		}

		for i := 0; i < hardCount; i++ {
			puzzle, err := generator.Generate("hard")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error generating hard puzzle: %v\n", err)
				os.Exit(1)
			}
			puzzles = append(puzzles, puzzle)
		}

		if len(puzzles) == 0 {
			fmt.Fprintln(os.Stderr, "No puzzles to generate")
			os.Exit(1)
		}

		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.SetAutoPageBreak(true, 10)

		drawGrid := func(p *gofpdf.Fpdf, puzzle string, startX, startY float64) {
			p.SetLineWidth(0.5)
			for i := 0; i <= 9; i++ {
				lw := 0.3
				if i%3 == 0 {
					lw = 0.8
				}
				p.SetLineWidth(lw)
				p.Line(startX, startY+float64(i)*9, startX+81, startY+float64(i)*9)
				p.Line(startX+float64(i)*9, startY, startX+float64(i)*9, startY+81)
			}

			p.SetFont("Helvetica", "", 14)
			for i, ch := range puzzle {
				row := i / 9
				col := i % 9
				x := startX + float64(col)*9 + 3
				y := startY + float64(row)*9 + 3
				if string(ch) != "-" {
					p.Text(x, y, string(ch))
				}
			}
		}

		pdf.AddPage()
		pdf.SetFont("Helvetica", "B", 24)
		pdf.Cell(0, 15, "Sudoku Puzzles")
		pdf.Ln(20)

		for i, puzzle := range puzzles {
			if i > 0 && i%6 == 0 {
				pdf.AddPage()
			}

			diff := "Easy"
			if i >= easyCount {
				if i >= easyCount+mediumCount {
					diff = "Hard"
				} else {
					diff = "Medium"
				}
			}

			pdf.SetFont("Helvetica", "", 10)
			pdf.Cell(0, 5, fmt.Sprintf("Puzzle %d (%s)", i+1, diff))
			pdf.Ln(5)

			drawGrid(pdf, puzzle.String("-"), 10, pdf.GetY())
			pdf.Ln(90)
		}

		pdf.AddPage()
		pdf.SetFont("Helvetica", "B", 24)
		pdf.Cell(0, 15, "Solutions")
		pdf.Ln(20)

		for i, puzzle := range puzzles {
			if i > 0 && i%6 == 0 {
				pdf.AddPage()
			}

			pdf.SetFont("Helvetica", "", 10)
			pdf.Cell(0, 5, fmt.Sprintf("Solution %d", i+1))
			pdf.Ln(5)

			drawGrid(pdf, puzzle.SolvedString(), 10, pdf.GetY())
			pdf.Ln(90)
		}

		err := pdf.OutputFileAndClose(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing PDF: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Created %s\n", filename)
	},
}

var easy, medium, hard string

func init() {
	bookCmd.Flags().StringVarP(&easy, "easy", "e", "5", "number of easy puzzles")
	bookCmd.Flags().StringVarP(&medium, "medium", "m", "5", "number of medium puzzles")
	bookCmd.Flags().StringVarP(&hard, "hard", "H", "5", "number of hard puzzles")
	RootCmd.AddCommand(bookCmd)
}
