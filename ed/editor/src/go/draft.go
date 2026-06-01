package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	var left []rune
	var right []rune

	for _, c := range input {
		switch c {
		case 'R':
			left = append(left, '\n')

		case 'B':
			if len(left) > 0 {
				left = left[:len(left)-1]
			}

		case 'D':
			if len(right) > 0 {
				right = right[1:]
			}

		case '<':
			if len(left) > 0 {
				last := left[len(left)-1]
				left = left[:len(left)-1]
				right = append([]rune{last}, right...)
			}

		case '>':
			if len(right) > 0 {
				left = append(left, right[0])
				right = right[1:]
			}

		default:
			left = append(left, c)
		}
	}

	fmt.Print(string(left))
	fmt.Print("|")
	fmt.Println(string(right))
}
