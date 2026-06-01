package main

import "fmt"

func main() {

	var n, x int
	fmt.Scan(&n)

	animal := make(map[int]int)
	casal := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&x)

		if animal[-x] > 0 {
			casal++
			animal[-x]--
		} else {
			animal[x]++
		}
	}
	fmt.Println(casal)
}
