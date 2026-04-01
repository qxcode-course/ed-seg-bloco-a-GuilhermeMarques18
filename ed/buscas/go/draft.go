package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	strings := make([]string, n)
	for i := range strings {
		fmt.Scan(&strings[i])
	}

	var m int
	fmt.Scan(&m)
	consultas := make([]string, m)
	for i := range consultas {
		fmt.Scan(&consultas[i])
	}

	result := matching(strings, consultas)

	for i, v := range result {
		fmt.Printf("%d", v)
		if i < len(result)-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Println()

}

func matching(strings []string, consultas []string) []int {
	verification := make(map[string]int)

	//Incrementa
	for _, s := range strings {
		verification[s]++
	}

	result := make([]int, len(consultas))

	//Go retorna valor zero se a chave não existir
	for i, v := range consultas {
		result[i] = verification[v]
	}
	return result
}
