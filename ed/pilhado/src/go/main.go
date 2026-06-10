package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	var n, m int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n, &m)

	matriz := make([][]rune, n)
	visitado := make([][]bool, n)
	for i := range matriz {
		matriz[i] = make([]rune, m)
		visitado[i] = make([]bool, m)
	}

	var initial, finish Post
	for i := 0; i < n; i++ {
		scanner.Scan()
		linha := scanner.Text()
		for j := 0; j < m && j < len(linha); j++ {
			matriz[i][j] = rune(linha[j])
			if matriz[i][j] == 'I' {
				initial = Post{l: i, c: j}
			} else if matriz[i][j] == 'F' {
				finish = Post{l: i, c: j}
			}
		}
	}

	caminho := NewStack[Post]()
	becos := NewStack[Post]()
	caminho.Push(initial)

	dl := []int{-1, 0, 1, 0}
	dc := []int{0, 1, 0, -1}

	for !caminho.IsEmpty() {
		atual := caminho.Top()
		visitado[atual.l][atual.c] = true
		if atual.l == finish.l && atual.c == finish.c {
			break
		}
		encontrou := false
		for i := 0; i < 4; i++ {
			nl := atual.l + dl[i]
			nc := atual.c + dc[i]
			if nl < 0 || nl >= n || nc < 0 || nc >= m {
				continue
			}
			if matriz[nl][nc] != '#' && !visitado[nl][nc] {
				caminho.Push(Post{l: nl, c: nc})
				encontrou = true
				break
			}
		}
		if !encontrou {
			beco := caminho.Pop()
			becos.Push(beco)
		}
	}

	for !caminho.IsEmpty() {
		pos := caminho.Pop()
		matriz[pos.l][pos.c] = '.'

	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(string(matriz[i][j]))
		}
		fmt.Println()
	}
}
