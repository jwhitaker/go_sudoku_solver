package solver

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Puzzle struct {
	Cells  [81]int
	Solved [81]int
}

func Generate(difficulty string) (*Puzzle, error) {
	rand.Seed(time.Now().UnixNano())

	solved := generateSolved()
	puzzle := removeNumbers(solved, difficulty)

	return &Puzzle{
		Cells:  puzzle,
		Solved: solved,
	}, nil
}

func generateSolved() [81]int {
	var grid [81]int
	for i := range grid {
		grid[i] = 0
	}
	solveSudoku(&grid)
	return grid
}

func solveSudoku(grid *[81]int) bool {
	for i := 0; i < 81; i++ {
		if grid[i] == 0 {
			nums := rand.Perm(9)
			for _, n := range nums {
				num := n + 1
				if isValid(grid, i/9, i%9, num) {
					grid[i] = num
					if solveSudoku(grid) {
						return true
					}
					grid[i] = 0
				}
			}
			return false
		}
	}
	return true
}

func isValid(grid *[81]int, row, col, num int) bool {
	for c := 0; c < 9; c++ {
		if grid[row*9+c] == num {
			return false
		}
	}
	for r := 0; r < 9; r++ {
		if grid[r*9+col] == num {
			return false
		}
	}
	boxRow, boxCol := row/3*3, col/3*3
	for r := boxRow; r < boxRow+3; r++ {
		for c := boxCol; c < boxCol+3; c++ {
			if grid[r*9+c] == num {
				return false
			}
		}
	}
	return true
}

func removeNumbers(solved [81]int, difficulty string) [81]int {
	var puzzle [81]int
	copy(puzzle[:], solved[:])

	attempts := getAttempts(difficulty)
	positions := rand.Perm(81)
	removed := 0

	for _, pos := range positions {
		if removed >= attempts {
			break
		}
		if puzzle[pos] != 0 {
			puzzle[pos] = 0
			removed++
		}
	}

	return puzzle
}

func getAttempts(difficulty string) int {
	switch difficulty {
	case "medium":
		return 40
	case "hard":
		return 50
	default:
		return 30
	}
}

func (p *Puzzle) String(emptyChar string) string {
	var result strings.Builder
	for i := 0; i < 81; i++ {
		if p.Cells[i] == 0 {
			result.WriteString(emptyChar)
		} else {
			result.WriteString(fmt.Sprintf("%d", p.Cells[i]))
		}
	}
	return result.String()
}

func (p *Puzzle) SolvedString() string {
	var result strings.Builder
	for i := 0; i < 81; i++ {
		result.WriteString(fmt.Sprintf("%d", p.Solved[i]))
	}
	return result.String()
}
