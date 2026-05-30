package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
	l.size++
}

func (l *LList) Size() int {
	return l.size
}

func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}

func (l *LList) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for n := l.root.next; n != l.root; n = n.next {
		sb.WriteString(fmt.Sprintf("%d", n.Value))
		if n.next != l.root {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func equals(la *LList, lb *LList) bool {
	if la.size != lb.size {
		return false
	}

	currA := la.root.next
	currB := lb.root.next

	for currA != la.root {
		if currA.Value != currB.Value {
			return false
		}
		currA = currA.next
		currB = currB.next
	}
	return true
}
func addsorted(la *LList, value int) {
	curr := la.root.next

	for curr != la.root && curr.Value < value {
		curr = curr.next
	}
	la.insertBefore(curr, value)
}

func reverse(l *LList) {
	if l.size <= 1 {
		return
	}
	curr := l.root

	for {
		last := curr.next

		curr.next = curr.prev
		curr.prev = last
		curr = last
		if curr == l.root {
			break
		}
	}
}

func merge(la *LList, lb *LList) *LList {
	merged := NewLList()

	currA := la.root.next
	currB := lb.root.next

	for currA != la.root && currB != lb.root {
		if currA.Value <= currB.Value {
			merged.PushBack(currA.Value)
			currA = currA.next
		} else {
			merged.PushBack(currB.Value)
			currB = currB.next
		}
	}

	for currA != la.root {
		merged.PushBack(currA.Value)
		currA = currA.next
	}

	for currB != lb.root {
		merged.PushBack(currB.Value)
		currB = currB.next
	}

	return merged
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(lla)
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)
			fmt.Println(lla)
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
