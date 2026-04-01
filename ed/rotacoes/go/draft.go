package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	vetor := make([]int, n)
	for i := range vetor {
		fmt.Scan(&vetor[i])
	}

	slice := inverte(m, vetor)
	fmt.Printf("[ ")
	for _, v := range slice {
		fmt.Print(v)
		fmt.Printf(" ")
	}
	fmt.Println("]")
}

func inverte(m int, vetor []int) []int {

	inverso := make([]int, len(vetor))

	for i := range vetor {
		result := (i + m) % len(vetor)
		inverso[result] = vetor[i]
	}

	return inverso
}
