.PHONY: build clean

build:
	go build -o sudoku-solver .

clean:
	rm -f sudoku-solver
