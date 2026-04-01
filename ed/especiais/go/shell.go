package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	list := make(map[int]int)
	for _, v := range vet {
		absV := v
		if v < 0 {
			absV = -v
		}
		list[absV]++
	}

	keys := make([]int, 0, len(list))
	for k := range list {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	result := []Pair{}
	for _, k := range keys {
		result = append(result, Pair{k, list[k]})
	}
	return result
}

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return []Pair{}
	}
	slice := []Pair{}
	count := 1

	for i := 0; i < len(vet); i++ {

		if i == len(vet)-1 || vet[i] != vet[i+1] {
			slice = append(slice, Pair{vet[i], count})
			count = 1
		} else {
			count++
		}

	}
	return slice
}

func mnext(vet []int) []int {
	slice := make([]int, len(vet))

	for i := 0; i < len(vet); i++ {

		if vet[i] > 0 {
			var woman bool
			if i > 0 && vet[i-1] < 0 {
				woman = true
			}
			if i < len(vet)-1 && vet[i+1] < 0 {
				woman = true
			}
			if woman {
				slice[i] = 1
			} else {
				slice[i] = 0
			}
		} else {
			slice[i] = 0
		}
	}
	return slice
}

func alone(vet []int) []int {
	slice := make([]int, len(vet))

	for i := 0; i < len(vet); i++ {

		if vet[i] > 0 {
			var man bool

			if i > 0 && vet[i-1] < 0 {
				man = true
			}

			if i < len(vet)-1 && vet[i+1] < 0 {
				man = true
			}

			if !man {
				slice[i] = 1
			} else {
				slice[i] = 0
			}
		} else {
			slice[i] = 0
		}
	}
	return slice
}

func couple(vet []int) int {
	lovers := make(map[int]int)
	var count int
	for _, v := range vet {
		if lovers[-v] > 0 {
			count++
			lovers[-v]--
		} else {
			lovers[v]++
		}
	}
	return count
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	if len(seq) == 0 {
		return -1
	}
	for i := 0; i <= len(vet)-len(seq); i++ {
		count := true
		for j := 0; j < len(seq); j++ {
			if vet[i+j] != seq[j] {
				count = false
				break
			}
		}
		if count {
			return i
		}
	}
	return -1
}

func clear(vet []int, value int) []int {
	count := []int{}
	for _, v := range vet {
		if v != value {
			count = append(count, v)
		}
	}
	return count
}

func erase(vet []int, posList []int) []int {
	sort.Slice(posList, func(i, j int) bool {
		return posList[i] > posList[j]
	})

	for _, v := range posList {
		if v >= 0 && v < len(vet) {
			vet = append(vet[:v], vet[v+1:]...)
		}
	}
	return vet
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
