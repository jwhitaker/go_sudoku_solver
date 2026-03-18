package solver

import (
	"fmt"
	"strings"
)

type Node struct {
	row    int
	col    int
	left   *Node
	right  *Node
	up     *Node
	down   *Node
	header *Node
}

type DLX struct {
	header   *Node
	columns  []*Node
	solution []int
	solved   bool
}

func NewDLX(cols int) *DLX {
	dlx := &DLX{
		columns:  make([]*Node, cols),
		solution: make([]int, 0),
	}

	dlx.header = &Node{}
	dlx.header.left = dlx.header
	dlx.header.right = dlx.header
	dlx.header.up = dlx.header
	dlx.header.down = dlx.header

	current := dlx.header
	for i := 0; i < cols; i++ {
		node := &Node{col: i}
		node.left = current
		node.right = dlx.header
		node.up = node
		node.down = node
		current.right = node
		dlx.header.left = node
		current = node
		dlx.columns[i] = node
	}

	return dlx
}

func (dlx *DLX) addRow(row int, columns []int) {
	var first *Node
	for _, col := range columns {
		header := dlx.columns[col]
		node := &Node{row: row, col: col, header: header}

		node.down = header
		node.up = header.up
		header.up.down = node
		header.up = node

		if first == nil {
			first = node
			node.left = node
			node.right = node
		} else {
			node.left = first.left
			node.right = first
			first.left.right = node
			first.left = node
		}
	}
}

func (dlx *DLX) cover(col *Node) {
	col.right.left = col.left
	col.left.right = col.right

	for row := col.down; row != col; row = row.down {
		for node := row.right; node != row; node = node.right {
			node.down.up = node.up
			node.up.down = node.down
		}
	}
}

func (dlx *DLX) uncover(col *Node) {
	for row := col.up; row != col; row = row.up {
		for node := row.left; node != row; node = node.left {
			node.down.up = node
			node.up.down = node
		}
	}
	col.right.left = col
	col.left.right = col
}

func (dlx *DLX) chooseColumn() *Node {
	min := 1<<31 - 1
	var selected *Node

	for col := dlx.header.right; col != dlx.header; col = col.right {
		count := 0
		for node := col.down; node != col; node = node.down {
			count++
		}
		if count < min {
			min = count
			selected = col
			if min == 0 {
				break
			}
		}
	}
	return selected
}

func (dlx *DLX) search(k int) {
	if dlx.header.right == dlx.header {
		dlx.solved = true
		return
	}

	col := dlx.chooseColumn()
	if col == nil {
		return
	}

	dlx.cover(col)

	for row := col.down; row != col; row = row.down {
		dlx.solution = append(dlx.solution, row.row)

		for node := row.right; node != row; node = node.right {
			dlx.cover(dlx.columns[node.col])
		}

		dlx.search(k + 1)

		if dlx.solved {
			return
		}

		for node := row.left; node != row; node = node.left {
			dlx.uncover(dlx.columns[node.col])
		}

		dlx.solution = dlx.solution[:len(dlx.solution)-1]
	}

	dlx.uncover(col)
}

func getBoxIndex(row, col int) int {
	return (row/3)*3 + (col / 3)
}

func buildExactCover(puzzle string, emptyChar string) (*DLX, []int, error) {
	dlx := NewDLX(324)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			cellIdx := r*9 + c
			rowIdxBase := 81 + r*9
			colIdxBase := 162 + c*9
			boxIdxBase := 243 + getBoxIndex(r, c)*9

			for n := 1; n <= 9; n++ {
				row := r*81 + c*9 + (n - 1)
				dlx.addRow(row, []int{
					cellIdx,
					rowIdxBase + (n - 1),
					colIdxBase + (n - 1),
					boxIdxBase + (n - 1),
				})
			}
		}
	}

	puzzle = strings.TrimSpace(puzzle)
	initialSolution := make([]int, 0)

	for i, ch := range puzzle {
		if string(ch) == emptyChar || ch == '0' {
			continue
		}
		var n int
		fmt.Sscanf(string(ch), "%d", &n)
		if n < 1 || n > 9 {
			continue
		}

		r := i / 9
		c := i % 9

		rowIdx := r*81 + c*9 + (n - 1)
		initialSolution = append(initialSolution, rowIdx)

		cellCol := r*9 + c
		rowCol := 81 + r*9 + (n - 1)
		colCol := 162 + c*9 + (n - 1)
		boxCol := 243 + getBoxIndex(r, c)*9 + (n - 1)

		dlx.cover(dlx.columns[cellCol])
		dlx.cover(dlx.columns[rowCol])
		dlx.cover(dlx.columns[colCol])
		dlx.cover(dlx.columns[boxCol])
	}

	return dlx, initialSolution, nil
}

func Solve(puzzle string, emptyChar string) (string, error) {
	puzzle = strings.TrimSpace(puzzle)
	if len(puzzle) != 81 {
		return "", fmt.Errorf("puzzle must be exactly 81 characters, got %d", len(puzzle))
	}

	validChars := "123456789" + emptyChar
	seen := make(map[rune]bool)
	for _, ch := range validChars {
		seen[ch] = true
	}
	seen['0'] = true

	for _, ch := range puzzle {
		if !seen[ch] {
			return "", fmt.Errorf("invalid character '%c' in puzzle", ch)
		}
	}

	dlx, initialSolution, err := buildExactCover(puzzle, emptyChar)
	if err != nil {
		return "", err
	}

	dlx.search(0)

	if !dlx.solved {
		return "", fmt.Errorf("no solution exists")
	}

	fullSolution := make([]int, len(initialSolution)+len(dlx.solution))
	copy(fullSolution, initialSolution)
	copy(fullSolution[len(initialSolution):], dlx.solution)

	solution := make([]byte, 81)
	for i := range solution {
		solution[i] = '.'
	}
	for _, rowIdx := range fullSolution {
		r := rowIdx / 81
		c := (rowIdx % 81) / 9
		n := (rowIdx % 9) + 1
		solution[r*9+c] = byte('0' + n)
	}

	return string(solution), nil
}

func FormatGrid(puzzle string, emptyChar string) string {
	var result strings.Builder
	for i, ch := range puzzle {
		if ch == '0' {
			ch = rune(emptyChar[0])
		}
		if i > 0 && i%9 == 0 {
			result.WriteString("\n")
		}
		if i%27 == 0 && i != 0 {
			result.WriteString("------+------+------\n")
		}
		if i%9 == 3 || i%9 == 6 {
			result.WriteString(" ")
		}
		result.WriteRune(ch)
		if i%9 != 8 {
			result.WriteString(" ")
		}
	}
	return result.String()
}

func IsValid(puzzle string, emptyChar string) bool {
	puzzle = strings.TrimSpace(puzzle)
	if len(puzzle) != 81 {
		return false
	}

	validChars := "123456789" + emptyChar
	seen := make(map[rune]bool)
	for _, ch := range validChars {
		seen[ch] = true
	}
	seen['0'] = true

	for _, ch := range puzzle {
		if !seen[ch] {
			return false
		}
	}

	grid := make([]int, 81)
	for i, ch := range puzzle {
		if ch == rune(emptyChar[0]) || ch == '0' {
			grid[i] = 0
		} else {
			n, _ := fmt.Sscanf(string(ch), "%d", &grid[i])
			if n == 0 {
				return false
			}
		}
	}

	for row := 0; row < 9; row++ {
		seen := make(map[int]bool)
		for col := 0; col < 9; col++ {
			val := grid[row*9+col]
			if val != 0 {
				if seen[val] {
					return false
				}
				seen[val] = true
			}
		}
	}

	for col := 0; col < 9; col++ {
		seen := make(map[int]bool)
		for row := 0; row < 9; row++ {
			val := grid[row*9+col]
			if val != 0 {
				if seen[val] {
					return false
				}
				seen[val] = true
			}
		}
	}

	for boxRow := 0; boxRow < 3; boxRow++ {
		for boxCol := 0; boxCol < 3; boxCol++ {
			seen := make(map[int]bool)
			for r := 0; r < 3; r++ {
				for c := 0; c < 3; c++ {
					row := boxRow*3 + r
					col := boxCol*3 + c
					val := grid[row*9+col]
					if val != 0 {
						if seen[val] {
							return false
						}
						seen[val] = true
					}
				}
			}
		}
	}

	return true
}
