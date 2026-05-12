package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	low := 0
	high := len(slice) - 1
	idx := -1

	for low <= high {
		mid := low + (high-low)/2

		if slice[mid] == value {
			idx = mid
			break
		} else if slice[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if idx != -1 {
		afterIdx := idx

		for i := idx + 1; i < len(slice); i++ {
			if slice[i] == value {
				afterIdx = i
			} else {
				break
			}
		}
		return afterIdx
	}
	return low
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
