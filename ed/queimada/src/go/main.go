package main

import (
	"bufio"
	"fmt"
	"os"
)

func burnTrees(grid [][]rune, l, c int) {
	// se estiver fora da matriz, retorne
	if l < 0 || c < 0 || l > len(grid)-1 || c > len(grid[0])-1 {
		return
	}

	// se o elemento atual não for uma arvore, retorne
	if grid[l][c] != '#' {
		return
	}
	// queime a arvore colocando o caractere 'o' na posição atual
	// chame a recursão para todos os 4 vizinhos
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

func showGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
