package main

import "fmt"

func div(n int) {
	if n == 0 {
		return
	}
	div(n / 2)

	resto := n % 2
	result := n / 2
	fmt.Println(result, resto)
}
func main() {
	var n int
	fmt.Scan(&n)
	div(n)
}
