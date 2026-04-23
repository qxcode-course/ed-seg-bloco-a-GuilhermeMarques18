package main

import (
	"fmt"
)

func mdc(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	result := mdc(b, a%b)

	return result
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
