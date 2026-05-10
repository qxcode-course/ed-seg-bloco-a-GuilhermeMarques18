package main

import (
	"fmt"
)

func main() {
	var seq string
	var L int

	if _, err := fmt.Scan(&seq); err != nil {
		return
	}

	if _, err := fmt.Scan(&L); err != nil {
		return
	}

	if len(seq) > 100 {
		return
	}

	if L > 100 {
		return
	}

	standard := []rune(seq)
	res := len(standard)

	var backtrack func(pos int) bool
	backtrack = func(pos int) bool {
		if pos == res {
			return true
		}

		if standard[pos] != '.' {
			return backtrack(pos + 1)
		}

		for d := 0; d <= L; d++ {
			charD := rune(d + '0')
			if isValid(standard, pos, charD, L) {
				standard[pos] = charD
				if backtrack(pos + 1) {
					return true
				}
				standard[pos] = '.'
			}
		}
		return false
	}

	if backtrack(0) {
		fmt.Println(string(standard))
	}

}

func isValid(res []rune, pos int, charD rune, L int) bool {
	for i := pos - L; i <= pos+L; i++ {
		if i >= 0 && i < len(res) && i != pos {
			if res[i] == charD {
				return false
			}
		}
	}
	return true
}
