package main

import (
	"fmt"
)

func main() {
	var num, sword, signal int
	fmt.Scan(&num, &sword, &signal)
	wheel := make([]int, num)

	inversion := signal
	for i := 0; i < num; i++ {
		wheel[i] = (i + 1) * inversion
		inversion *= -1
	}

	pos := 0
	for i, v := range wheel {
		if v == sword || v == -sword {
			pos = i
			break
		}
	}

	for len(wheel) > 1 {
		io(wheel, pos)
		kill := 0
		if wheel[pos] > 0 {
			kill = (pos + 1) % len(wheel)
		} else {
			kill = (pos - 1 + len(wheel)) % len(wheel)
		}

		killed := kill < pos
		wheel = append(wheel[:kill], wheel[kill+1:]...)

		if killed {
			pos--
		}

		if wheel[pos] > 0 {
			pos = kill % len(wheel)
		} else {
			pos = (kill - 1 + len(wheel)) % len(wheel)
		}
	}
	io(wheel, pos)
}
func io(vet []int, pos int) {
	fmt.Print("[ ")

	for i, v := range vet {
		if i == pos {
			if v > 0 {
				fmt.Printf("%d> ", v)
			} else {
				fmt.Printf("<%d ", v)
			}
		} else {
			fmt.Printf("%d ", v)
		}
	}

	fmt.Println("]")
}
