package main

import "fmt"

// x: número que está sendo testado
// div: divisor que está sendo testado
func eh_primo(x int, div int) bool {
	if x < 2 {
		return false
	}
	if div*div > x {
		return true
	}
	eh_primo(x, div+1)

	result := x%div == 0
	if result {
		return false
	}
	return true
}

func main() {
	var x int
	fmt.Scan(&x)
	if eh_primo(x, 2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
