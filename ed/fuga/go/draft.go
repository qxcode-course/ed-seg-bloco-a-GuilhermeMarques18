package main

import "fmt"

func fuga(H int, P int, F int, D int) string {
	pos := F

	for {
		if D == -1 {
			pos = (pos - 1 + 16) % 16
		} else {
			pos = (pos + 1) % 16
		}

		if pos == P {
			return "N"
		}
		if pos == H {
			return "S"
		}
	}
}
func main() {
	var H int
	var P int
	var F int
	var D int

	fmt.Scan(&H, &P, &F, &D)

	fmt.Println(fuga(H, P, F, D))

}
