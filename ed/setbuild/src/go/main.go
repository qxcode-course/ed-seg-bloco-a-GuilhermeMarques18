package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, capacity),
		capacity: capacity,
	}
}

func (s *Set) String() string {
	sb := strings.Builder{}
	sb.WriteRune('[')
	for i, value := range s.data[:s.size] {
		sb.WriteString(fmt.Sprintf("%d", value))
		if i < s.size-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteRune(']')
	return sb.String()
}

func (s *Set) insert(value int, idx int) error {
	if idx < 0 || idx > s.size {
		return errors.New("index out of bounds")
	}
	for i := s.size; i > idx; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[idx] = value
	s.size++
	return nil
}

func (s *Set) Insert(value int) {
	idx := s.binarySearch(value)

	if idx < s.size && s.data[idx] == value {
		return
	}

	if s.size == s.capacity {
		newCap := s.capacity * 2
		if newCap == 0 {
			newCap = 1
		}
		s.reserve(newCap)
	}
	s.insert(value, idx)
}

func (s *Set) reserve(newCapacity int) {
	if newCapacity <= s.capacity {
		return
	}
	newData := make([]int, newCapacity)

	copy(newData, s.data)

	s.data = newData
	s.capacity = newCapacity
}

func (s *Set) binarySearch(value int) int {
	low := 0
	high := s.size - 1

	for low <= high {
		mid := (low + high) / 2

		if s.data[mid] == value {
			return mid
		}
		if s.data[mid] < value {
			low = mid + 1

		} else {
			high = mid - 1
		}
	}
	return low
}

func (s *Set) erase(idx int) error {
	if idx < 0 || idx >= s.size {
		return errors.New("index out of bounds")
	}

	for i := idx; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}

	s.size--
	return nil
}

func (s *Set) Erase(value int) bool {
	idx := s.binarySearch(value)

	if idx < s.size && s.data[idx] == value {
		s.erase(idx)
		return true
	}

	fmt.Println("value not found")
	return false
}

func (s *Set) Contains(value int) bool {
	idx := s.binarySearch(value)

	if idx < s.size && s.data[idx] == value {
		return true
	}
	return false
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	s := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			s = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, err := strconv.Atoi(part)
				if err == nil {
					s.Insert(value)
				}
			}
		case "show":
			fmt.Println(s.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			s.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if s.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
