package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(grid) == 0 || len(grid[0]) == 0 {
		return false

	}
	m := len(grid)
	n := len(grid[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var dfs func(r, c, idx int) bool
	dfs = func(r, c, idx int) bool {
		if idx == len(word) {
			return true
		}
		if r < 0 || r >= m || c < 0 || c >= n {
			return false
		}
		if visited[r][c] {
			return false
		}
		if grid[r][c] != word[idx] {
			return false
		}
		visited[r][c] = true
		found := dfs(r+1, c, idx+1) ||
			dfs(r-1, c, idx+1) ||
			dfs(r, c+1, idx+1) ||
			dfs(r, c-1, idx+1)

		visited[r][c] = false
		return found
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
