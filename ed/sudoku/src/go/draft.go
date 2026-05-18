package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	matriz := make([][]rune, n)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		if scanner.Scan() {
			linha := scanner.Text()
			matriz[i] = []rune(linha)
		}
	}

	if resolver(matriz, 0) {
		imprimirMatriz(matriz)
	} else {
		fmt.Println("Não há solução")
	}
}

func naLinha(matriz [][]rune, lin int, num rune) bool {
	for c := 0; c < len(matriz); c++ {
		if matriz[lin][c] == num {
			return true
		}
	}
	return false
}

func naColuna(matriz [][]rune, col int, num rune) bool {
	for l := 0; l < len(matriz); l++ {
		if matriz[l][col] == num {
			return true
		}
	}
	return false
}

func noQuadrante(matriz [][]rune, lin, col int, num rune) bool {
	dim := len(matriz)
	tamQuadrante := int(math.Sqrt(float64(dim)))

	// Encontra o topo esquerdo do quadrante atual
	lInicio := (lin / tamQuadrante) * tamQuadrante
	cInicio := (col / tamQuadrante) * tamQuadrante

	for l := 0; l < tamQuadrante; l++ {
		for c := 0; c < tamQuadrante; c++ {
			if matriz[lInicio+l][cInicio+c] == num {
				return true
			}
		}
	}
	return false
}

func resolver(matriz [][]rune, index int) bool {
	nl := len(matriz)

	if index == nl*nl {
		return true
	}

	l := index / nl
	c := index % nl

	if matriz[l][c] != '.' {
		return resolver(matriz, index+1)
	}

	for i := 1; i <= nl; i++ {
		num := rune('0' + i)
		if !naLinha(matriz, l, num) && !naColuna(matriz, c, num) && !noQuadrante(matriz, l, c, num) {
			matriz[l][c] = num

			if resolver(matriz, index+1) {
				return true
			}

			matriz[l][c] = '.'
		}
	}

	return false
}

func imprimirMatriz(matriz [][]rune) {
	for _, linha := range matriz {
		fmt.Println(string(linha))
	}
}
