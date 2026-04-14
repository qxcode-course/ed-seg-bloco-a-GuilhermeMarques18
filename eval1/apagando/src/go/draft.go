package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var person int
	list := make(map[int]int)

	for i := 0; i < n; i++ {
		fmt.Scan(&person)

		list[i] = person
	}

	var m, sair int
	fmt.Scan(&m)

	exit := make(map[int]bool)
	for i := 0; i < m; i++ {
		fmt.Scan(&sair)

		exit[sair] = true
	}

	for i := 0; i < n; i++ {
		v, ok := list[i]
		if ok && !exit[v] {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()
}
