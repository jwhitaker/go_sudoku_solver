# Sudoku Solver with Dancing Links

## Project Overview
- **Project name**: sudoku-solver
- **Type**: CLI tool
- **Core functionality**: Solve Sudoku puzzles using the Dancing Links (DLX) algorithm
- **Target users**: Sudoku enthusiasts, developers learning DLX

## Functionality Specification

### Core Features
1. **Dancing Links (DLX) Algorithm**: Implement Algorithm X with dancing links for efficient exact cover solving
2. **Solve Command**: Accept 81-character puzzle string from positional argument or stdin
3. **Generate Command**: Generate a 81 character string of a valid Sudoku puzzle
4. **Display Command**: Render a Sudoku board from an 81-character string
5. **IsValid Command**: Check if a puzzle is valid (no duplicates in row/col/box)
6. **Input Validation**: Validate puzzle format (81 chars, valid digits, valid empty marker)

### Command Line Interface
- `solve [puzzle]`: Solve a Sudoku puzzle (positional argument or stdin)
  - `-e, --empty` flag: Character representing empty cells (default: `-`)
- `generate`: Generate a new Sudoku puzzle
  - `-d, --difficulty` flag: Difficulty level - easy (30 removed), medium (40 removed), hard (50 removed) (default: easy)
  - `-e, --empty` flag: Character representing empty cells (default: `-`)
- `display [puzzle]`: Display a Sudoku board (positional argument or stdin)
  - `-e, --empty` flag: Character representing empty cells (default: `-`)
- `isvalid [puzzle]`: Check if a puzzle is valid (positional argument or stdin)
  - `-e, --empty` flag: Character representing empty cells (default: `-`)
- `-h, --help` flag: Display usage information

### Input Validation
- Puzzle must be exactly 81 characters (or 82+ including newlines)
- Only digits 1-9 and the empty cell marker allowed
- Each row must have at most 9 digits
- No duplicate digits in any row, column, or 3x3 box (pre-filled cells)

### Edge Cases
- Invalid puzzle format: Show error and exit
- Unsolvable puzzle: Display "No solution exists"
- Multiple solutions: Display first solution found

## Technical Specification

### Puzzle Structure (generator package)
- `Puzzle` struct holds both the puzzle and its solved state (generator package):
  - `Cells [81]int`: The puzzle grid (0 = empty)
  - `Solved [81]int`: The solved grid
- `String(emptyChar)` method: Returns 81-char string representation with empty char
- `SolvedString()` method: Returns 81-char string of solution
- `FormatGrid(s, emptyChar string)` function: Formats 81-char string as readable 9x9 grid
- `IsValid(puzzle, emptyChar string)` function: Returns true if puzzle is valid

### Dancing Links Structure
- Implement exact cover matrix for Sudoku constraints:
  1. Cell constraint (81 columns) - each cell must have exactly one number
  2. Row constraint (81 columns) - each row must have each number 1-9
  3. Column constraint (81 columns) - each column must have each number 1-9
  4. Box constraint (81 columns) - each box must have each number 1-9
- Total: 324 columns

### Algorithm
- Algorithm X with dancing links for O(1) cover/uncover operations
- Search with heuristic: choose column with fewest nodes

## Acceptance Criteria
1. Correctly solves valid Sudoku puzzles
2. Reports "No solution" for unsolvable puzzles
3. Empty cell marker is configurable
4. Provides clear error messages for invalid input
5. Returns solution in under 1 second for typical puzzles
6. Generates valid, solvable Sudoku puzzles at three difficulty levels
