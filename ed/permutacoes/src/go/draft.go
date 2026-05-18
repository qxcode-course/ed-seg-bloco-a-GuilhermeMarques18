package main

import (
	"fmt"
	"sort"
)

func Permutacao(seq string) []string {
	var result []string
	runes := []rune(seq)
	backtrack(runes, 0, &result)
	return result
}

func backtrack(runes []rune, start int, result *[]string) {
	if start == len(runes)-1 {
		*result = append(*result, string(runes))
		return
	}

	for i := start; i < len(runes); i++ {
		runes[start], runes[i] = runes[i], runes[start]

		backtrack(runes, start+1, result)

		runes[start], runes[i] = runes[i], runes[start]
	}
}

func main() {
	var s string

	if _, err := fmt.Scan(&s); err != nil {
		return
	}

	if len(s) < 1 || len(s) > 6 {
		return
	}

	result := Permutacao(s)

	sort.Strings(result)

	for _, v := range result {
		fmt.Println(v)
	}

}
