package main

import "fmt"

func io(fila []int, pos int) {
	fmt.Print("[ ")
	for i := 0; i < len(fila); i++ {
		fmt.Print(fila[i])
		if i == pos {
			fmt.Print(">")
		}
		if i < len(fila)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println(" ]")
}

func main() {
	var qtd int
	var sword int

	fmt.Scan(&qtd, &sword)

	fila := make([]int, qtd)
	for i := 0; i < qtd; i++ {
		fila[i] = i + 1
	}

	pos := sword - 1

	io(fila, pos)

	for len(fila) > 1 {

		kill := (pos + 1) % len(fila)

		fila = append(fila[:kill], fila[kill+1:]...)

		if kill < pos {
			pos--
		}

		pos = (pos + 1) % len(fila)

		io(fila, pos)
	}
}
