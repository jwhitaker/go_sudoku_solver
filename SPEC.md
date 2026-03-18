# Sudoku Solver with Dancing Links

## Project Overview
- **Project name**: sudoku-solver
- **Type**: CLI tool
- **Core functionality**: Solve Sudoku puzzles using the Dancing Links (DLX) algorithm
- **Target users**: Sudoku enthusiasts, developers learning DLX

## Functionality Specification

### Core Features
1. **Dancing Links (DLX) Algorithm**: Implement Algorithm X with dancing links for efficient exact cover solving
2. **CLI Interface**: Accept 81-character string from command line representing the puzzle
3. **Configurable Empty Cell**: Allow users to specify the empty cell character via `-e` flag (default: `-`)
4. **Input Validation**: Validate puzzle format (81 chars, valid digits, valid empty marker)
5. **Solution Output**: Display solved puzzle as 81-character string and formatted 9x9 grid

### Command Line Interface
- Positional argument: 81-character puzzle string
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
