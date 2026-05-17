package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyList struct {
	data []int
}

type Iterator struct {
	data  []int
	index int
}

type ReverseIterator struct {
	data  []int
	index int
}

type CyclicIterator struct {
	data  []int
	index int
}

func NewMyList(values []int) *MyList {
	return &MyList{data: values}
}

func (l *MyList) Iterator() *Iterator {
	return &Iterator{data: l.data, index: -1}
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.data)-1
}

func (i *Iterator) Next() int {
	if i.index == len(i.data) {
		panic(fmt.Errorf("No more elements"))
	}
	i.index += 1
	return i.data[i.index]
}

func (ri *MyList) ReverseIterator() *ReverseIterator {
	return &ReverseIterator{data: ri.data, index: len(ri.data)}
}

func (ci *MyList) CyclicIterator() *CyclicIterator {
	return &CyclicIterator{
		data:  ci.data,
		index: -1,
	}
}

func (ri *ReverseIterator) HasNext() bool {
	if ri.index > 0 {
		return true
	}
	return false
}

func (ri *ReverseIterator) Next() int {
	if len(ri.data) == 0 || ri.index <= 0 {
		panic(fmt.Errorf("No more elements"))
	}
	ri.index -= 1
	return ri.data[ri.index]
}

func (ci *CyclicIterator) HasNext() bool {
	if len(ci.data) == 0 {
		return false
	}
	return true
}

func (ci *CyclicIterator) Next() int {
	if len(ci.data) == 0 {
		panic(fmt.Errorf("No more elements"))
	}
	ci.index += 1
	if ci.index >= len(ci.data) {
		ci.index = 0
	}
	return ci.data[ci.index]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mylist := NewMyList([]int{})
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			break
		case "read":
			for i := 1; i < len(args); i++ {
				slice := make([]int, len(args)-1)
				for i, value := range args[1:] {
					slice[i], _ = strconv.Atoi(value)
				}
				mylist = NewMyList(slice)
			}
		case "show":
			fmt.Print("[ ")
			for it := mylist.Iterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "reverse":
			fmt.Print("[ ")
			for it := mylist.ReverseIterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "cyclic":
			qtd, _ := strconv.Atoi(args[1])
			fmt.Print("[ ")
			it := mylist.CyclicIterator()
			for range qtd {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		}
	}

}
