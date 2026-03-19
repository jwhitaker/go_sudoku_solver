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

		pageWidth := 210.0
		pageHeight := 297.0
		margin := 15.0
		gridSize := 180.0
		cellSize := gridSize / 9.0

		startX := (pageWidth - gridSize) / 2
		startY := margin + 15

		drawLargeGrid := func(p *gofpdf.Fpdf, puzzle string) {
			for i := 0; i <= 9; i++ {
				lw := 0.4
				if i%3 == 0 {
					lw = 1.0
				}
				p.SetLineWidth(lw)
				p.Line(startX, startY+float64(i)*cellSize, startX+gridSize, startY+float64(i)*cellSize)
				p.Line(startX+float64(i)*cellSize, startY, startX+float64(i)*cellSize, startY+gridSize)
			}

			p.SetFont("Helvetica", "", 16)
			for i, ch := range puzzle {
				if string(ch) != "-" {
					row := i / 9
					col := i % 9
					x := startX + float64(col)*cellSize + cellSize/2
					y := startY + float64(row)*cellSize + cellSize/2 + 5
					p.Text(x-p.GetStringWidth(string(ch))/2, y, string(ch))
				}
			}
		}

		solutionSize := 50.0
		solutionCellSize := solutionSize / 9.0

		drawSmallGrid := func(p *gofpdf.Fpdf, puzzle string, startX, startY float64) {
			for i := 0; i <= 9; i++ {
				lw := 0.2
				if i%3 == 0 {
					lw = 0.5
				}
				p.SetLineWidth(lw)
				p.Line(startX, startY+float64(i)*solutionCellSize, startX+solutionSize, startY+float64(i)*solutionCellSize)
				p.Line(startX+float64(i)*solutionCellSize, startY, startX+float64(i)*solutionCellSize, startY+solutionSize)
			}

			p.SetFont("Helvetica", "", 5)
			for i, ch := range puzzle {
				row := i / 9
				col := i % 9
				x := startX + float64(col)*solutionCellSize + solutionCellSize/2
				y := startY + float64(row)*solutionCellSize + solutionCellSize/2 + 1.5
				p.Text(x-p.GetStringWidth(string(ch))/2, y, string(ch))
			}
		}

		pdf.AddPage()
		pdf.SetFont("Helvetica", "B", 36)
		pdf.Cell(0, pageHeight/2-20, title)
		pdf.Ln(10)
		pdf.SetFont("Helvetica", "", 14)
		pdf.Cell(0, 10, fmt.Sprintf("%d Easy | %d Medium | %d Hard", easyCount, mediumCount, hardCount))
		pdf.Ln(10)
		pdf.SetFont("Helvetica", "", 12)
		pdf.Cell(0, 10, fmt.Sprintf("%d Puzzles", len(puzzles)))

		pdf.AddPage()

		for i, puzzle := range puzzles {
			if i > 0 {
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

			pdf.SetFont("Helvetica", "B", 14)
			pdf.Cell(0, 8, fmt.Sprintf("Puzzle %d (%s)", i+1, diff))

			drawLargeGrid(pdf, puzzle.String("-"))
		}

		pdf.AddPage()

		solutionsPerRow := 3
		solutionsPerCol := 3
		solutionsPerPage := solutionsPerRow * solutionsPerCol
		solutionSpacingX := solutionSize + 10
		solutionSpacingY := solutionSize + 12
		startXSolutions := (pageWidth - (float64(solutionsPerRow)*solutionSize + float64(solutionsPerRow-1)*10)) / 2

		for i := range puzzles {
			if i > 0 && i%solutionsPerPage == 0 {
				pdf.AddPage()
			}

			pageIndex := i % solutionsPerPage
			pageRow := pageIndex / solutionsPerRow
			pageCol := pageIndex % solutionsPerRow

			x := startXSolutions + float64(pageCol)*solutionSpacingX
			y := margin + 10 + float64(pageRow)*solutionSpacingY

			pdf.SetFont("Helvetica", "", 7)
			pdf.Text(x+solutionSize/2-2.5, y-2.5, fmt.Sprintf("%d", i+1))

			drawSmallGrid(pdf, puzzles[i].SolvedString(), x, y)
		}

		err := pdf.OutputFileAndClose(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing PDF: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Created %s\n", filename)
	},
}

var easy, medium, hard, title string

func init() {
	bookCmd.Flags().StringVarP(&easy, "easy", "e", "5", "number of easy puzzles")
	bookCmd.Flags().StringVarP(&medium, "medium", "m", "5", "number of medium puzzles")
	bookCmd.Flags().StringVarP(&hard, "hard", "H", "5", "number of hard puzzles")
	bookCmd.Flags().StringVarP(&title, "title", "t", "Sudoku Puzzles", "title of the book")
	RootCmd.AddCommand(bookCmd)
}
