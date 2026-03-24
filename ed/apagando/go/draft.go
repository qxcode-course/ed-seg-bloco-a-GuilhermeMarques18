package main

import "fmt"

func main() {
	var n, idn int
	fmt.Scan(&n)
	list := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&idn)

		list[i] = idn
	}

	var m, idm int
	fmt.Scan(&m)

	desistentes := make(map[int]bool)
	for i := 0; i < m; i++ {
		fmt.Scan(&idm)
		desistentes[idm] = true
	}

	for i := 0; i < n; i++ {
		v, ok := list[i]
		if ok && !desistentes[v] {
			fmt.Printf("%d ", v)
		}
	}

	fmt.Println()

}
