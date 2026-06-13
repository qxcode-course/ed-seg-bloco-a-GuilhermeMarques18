package main

import (
	"fmt"
)

type Bomba struct {
	gas, dis int
}

func test(saldo int) bool {
	return saldo >= 0
}
func main() {
	var n int

	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	cam := make([]Bomba, n)

	for i := range cam {
		_, err := fmt.Scan(&cam[i].gas, &cam[i].dis)
		if err != nil {
			return
		}
	}

	saldo := 0
	for _, bomba := range cam {
		saldo += bomba.gas - bomba.dis
	}

	if test(saldo) {
		tanque := 0
		pos := 0

		for i, bomba := range cam {
			tanque += bomba.gas - bomba.dis

			if tanque < 0 {
				pos = i + 1
				tanque = 0
			}
		}
		fmt.Println(pos)

	}

}
