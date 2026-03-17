package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	album := make([]int, b)
	possui := make(map[int]bool)

	for i := 0; i < b; i++ {
		fmt.Scan(&album[i])
		possui[album[i]] = true
	}
	var reper []string
	for i := 1; i < b; i++ {
		if album[i] == album[i-1] {
			reper = append(reper, strconv.Itoa(album[i]))
		}
	}

	if len(reper) == 0 {
		fmt.Println("N")
	} else {
		fmt.Println(strings.Join(reper, " "))
	}

	var faltantes []string
	for i := 1; i <= a; i++ {
		if !possui[i] {
			faltantes = append(faltantes, strconv.Itoa(i))
		}
	}
	if len(faltantes) == 0 {
		fmt.Println("N")
	} else {
		fmt.Println(strings.Join(faltantes, " "))
	}
}
