package main

import (
	"fmt"
	"math"
)

func heron(l1 float64, l2 float64, l3 float64) float64 {
	p := (l1 + l2 + l3) / 2

	result := p * (p - l1) * (p - l2) * (p - l3)
	area := math.Sqrt(float64(result))

	return area
}

func main() {
	var l1 float64
	var l2 float64
	var l3 float64

	fmt.Scan(&l1, &l2, &l3)

	finish := heron(l1, l2, l3)
	fmt.Printf("%.2f\n", finish)

}
