package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	figurinha := make([]int, b)
	count := make(map[int]bool)

	for i := 0; i < b; i++ {
		fmt.Scan(&figurinha[i])
		count[figurinha[i]] = true
	}

	var reper []string
	for i := 1; i < b; i++ {
		if figurinha[i] == figurinha[i-1] {
			reper = append(reper, strconv.Itoa(figurinha[i]))
		}
	}

	if len(reper) == 0 {
		fmt.Println("N")
	} else {
		fmt.Println(strings.Join(reper, " "))
	}
	faltar := []string{}
	for i := 1; i <= a; i++ {
		if !count[i] {
			faltar = append(faltar, strconv.Itoa(i))
		}
	}
	if len(faltar) == 0 {
		fmt.Println("N")
	} else {
		fmt.Println(strings.Join(faltar, " "))
	}
}
