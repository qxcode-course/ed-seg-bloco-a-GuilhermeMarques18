package main

import (
	"fmt"
)

func balanceada(key string) bool {
	var load []rune

	chaveamento := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, v := range key {
		if v == '(' || v == '{' || v == '[' {
			load = append(load, v)
			continue
		}
		if len(load) == 0 {
			return false
		}

		if open, i := chaveamento[v]; i {

			if len(chaveamento) == 0 {
				return false
			}

			top := load[len(load)-1]
			load = load[:len(load)-1]

			if top != open {
				return false
			}

		}
	}
	return len(load) == 0
}
func main() {
	var n string
	fmt.Scan(&n)

	if balanceada(n) {
		fmt.Println("balanceado")
	} else {
		fmt.Println("nao balanceado")
	}

}
