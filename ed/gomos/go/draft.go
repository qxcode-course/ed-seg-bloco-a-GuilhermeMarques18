package main

import "fmt"

type Ponto struct {
	x, y int
}

func main() {
	var Q int
	var D string

	fmt.Scan(&Q, &D)

	py := make([]Ponto, Q)
	for i := 0; i < Q; i++ {
		fmt.Scan(&py[i].x, &py[i].y)
	}

	pyPassada := make([]Ponto, Q)
	copy(pyPassada, py)
	switch D {
	case "L":
		py[0].x--
	case "R":
		py[0].x++
	case "U":
		py[0].y--
	case "D":
		py[0].y++
	}
	for i := 1; i < Q; i++ {
		py[i] = pyPassada[i-1]
	}
	for i := 0; i < Q; i++ {
		fmt.Printf("%d %d\n", py[i].x, py[i].y)
	}

}
