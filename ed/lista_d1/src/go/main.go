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
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	rootNode := &Node{Value: 0}
	rootNode.next = rootNode
	rootNode.prev = rootNode
	return &LList{
		root: rootNode,
		size: 0,
	}
}

func (l *LList) Size() int {
	return l.size
}

func (l *LList) insertAfter(node *Node, value int) {
	newNode := &Node{Value: value, prev: node, next: node.next}
	node.next.prev = newNode
	node.next = newNode
	l.size++
}

func (l *LList) remove(node *Node) {
	if node == l.root {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
}

func (l *LList) PushFront(value int) {
	l.insertAfter(l.root, value)
}

func (l *LList) PushBack(value int) {
	l.insertAfter(l.root.prev, value)
}

func (l *LList) PopFront() {
	if l.size > 0 {
		l.remove(l.root.next)
	}
}

func (l *LList) PopBack() {
	if l.size > 0 {
		l.remove(l.root.prev)
	}
}

func (l *LList) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func (l *LList) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	curr := l.root.next
	for curr != l.root {
		sb.WriteString(strconv.Itoa(curr.Value))
		if curr.next != l.root {
			sb.WriteString(", ")
		}
		curr = curr.next
	}
	sb.WriteString("]")
	return sb.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		cmd := args[0]
		fmt.Printf("$%s\n", line)
		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
