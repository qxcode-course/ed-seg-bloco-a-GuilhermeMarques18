package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{data: make([]int, capacity), size: 0, capacity: capacity}
}

func (ms *MultiSet) Insert(value int) {
	if ms.size == ms.capacity {
		newCap := ms.capacity * 2

		if newCap == 0 {
			newCap = 1
		}
		newData := make([]int, newCap)
		copy(newData, ms.data[:ms.size])
		ms.data = newData
		ms.capacity = newCap
	}

	idx := sort.SearchInts(ms.data[:ms.size], value)

	for i := ms.size; i > idx; i-- {
		ms.data[i] = ms.data[i-1]
	}

	ms.data[idx] = value
	ms.size++
}

func (ms *MultiSet) String() string {
	if ms.size == 0 {
		return "[]"
	}
	var strValues []string
	for i := 0; i < ms.size; i++ {
		strValues = append(strValues, strconv.Itoa(ms.data[i]))
	}
	return "[" + strings.Join(strValues, ", ") + "]"
}

func (ms *MultiSet) Contains(value int) bool {

	idx := sort.SearchInts(ms.data[:ms.size], value)

	return idx < ms.size && ms.data[idx] == value
}

func (ms *MultiSet) Erase(value int) {
	idx := sort.SearchInts(ms.data[:ms.size], value)
	if idx < ms.size && ms.data[idx] == value {

		for i := idx; i < ms.size-1; i++ {
			ms.data[i] = ms.data[i+1]
		}

		ms.size--
	} else {
		fmt.Println("value not found")
	}
}

func (ms *MultiSet) Count(value int) int {
	idx := sort.SearchInts(ms.data[:ms.size], value)
	count := 0

	for i := idx; i < ms.size && ms.data[i] == value; i++ {
		count++
	}

	return count
}

func (ms *MultiSet) Unique() int {
	if ms.size == 0 {
		return 0
	}
	unique := 1
	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i-1] {
			unique++
		}
	}
	return unique
}

func (ms *MultiSet) Clear() {
	ms.size = 0
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	var ms *MultiSet

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			val, _ := strconv.Atoi(args[1])
			ms.Erase(val)
		case "contains":
			val, _ := strconv.Atoi(args[1])
			if ms.Contains(val) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			val, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(val))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
