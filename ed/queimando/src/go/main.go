package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l int
	c int
}

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()
	stack.Push(Pos{l, c})

	for !stack.IsEmpty() {
		after := stack.Pop()

		if after.l < 0 || after.c < 0 || after.l >= len(grid) || after.c >= len(grid[0]) {
			continue
		}

		if grid[after.l][after.c] == '#' {

			grid[after.l][after.c] = 'o'

			stack.Push(Pos{after.l - 1, after.c})
			stack.Push(Pos{after.l + 1, after.c})
			stack.Push(Pos{after.l, after.c - 1})
			stack.Push(Pos{after.l, after.c + 1})
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
