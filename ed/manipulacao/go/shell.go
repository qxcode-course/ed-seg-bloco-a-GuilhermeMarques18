package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	var man []int
	for _, v := range vet {
		if v > 0 {
			man = append(man, v)
		}
	}
	return man
}

func getCalmWomen(vet []int) []int {
	var women []int
	for _, v := range vet {
		if v < 0 && v > -10 {
			women = append(women, v)
		}
	}
	return women
}

func sortVet(vet []int) []int {
	sort.Ints(vet)
	return vet
}

func sortStress(vet []int) []int {
	sort.Slice(vet, func(i, j int) bool {
		si := vet[i]
		if si < 0 {
			si = -si
		}

		sj := vet[j]
		if sj < 0 {
			sj = -sj
		}
		return si < sj
	})
	return vet
}

func reverse(vet []int) []int {
	list := make([]int, len(vet))
	var pos int
	for i, v := range vet {
		pos = len(vet) - 1 - i
		list[pos] = v
	}
	return list
}

func unique(vet []int) []int {
	list := make(map[int]bool)
	res := []int{}
	for _, v := range vet {
		if !list[v] {
			list[v] = true
			res = append(res, v)
		}
	}
	return res
}

func repeated(vet []int) []int {
	list := make(map[int]int)
	res := []int{}
	for _, v := range vet {
		list[v]++
		if list[v] == 2 {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
