package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n > 15 {
		return
	}

	xadrez := make([][]rune, n)

	for i := 0; i < n; i++ {
		xadrez[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			xadrez[i][j] = '.'
		}
	}

	result := backtracking(xadrez, 0)

	fmt.Println(result)
}

func backtracking(tabuleiro [][]rune, linha int) int {
	n := len(tabuleiro)

	if linha == n {
		return 1
	}

	total := 0

	for coluna := 0; coluna < n; coluna++ {
		if check(tabuleiro, linha, coluna) {
			tabuleiro[linha][coluna] = 'Q'

			total += backtracking(tabuleiro, linha+1)

			tabuleiro[linha][coluna] = '.'
		}
	}

	return total
}

func check(tabuleiro [][]rune, linha int, coluna int) bool {
	n := len(tabuleiro)

	for l := 0; l < linha; l++ {
		if tabuleiro[l][coluna] == 'Q' {
			return false
		}
	}

	for l, c := linha-1, coluna-1; l >= 0 && c >= 0; l, c = l-1, c-1 {
		if tabuleiro[l][c] == 'Q' {
			return false
		}
	}

	for l, c := linha-1, coluna+1; l >= 0 && c < n; l, c = l-1, c+1 {
		if tabuleiro[l][c] == 'Q' {
			return false
		}
	}

	return true
}
