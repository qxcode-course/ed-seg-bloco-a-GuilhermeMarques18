package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Scan(&n)

	m := 101
	result := -1

	var a, b int
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)

		if a < 10 || b < 10 {
			continue
		}

		dif := int(math.Abs(float64(a - b)))

		if dif < m {
			m = dif
			result = i
		}
	}

	if result == -1 {
		fmt.Println("sem ganhador")
		return
	} else {
		fmt.Println(result)
	}

}
