package main

import "fmt"

func subsetSum(idx int, current int, target int, arr []int) bool {
	if current == target {
		return true
	}

	if current > target || idx >= len(arr) {
		return false
	}

	return subsetSum(idx+1, current+arr[idx], target, arr) || subsetSum(idx+1, current, target, arr)
}
func main() {
	var n, k int

	if _, err := fmt.Scan(&n, &k); err != nil {
		return
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	if subsetSum(0, 0, k, arr) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
