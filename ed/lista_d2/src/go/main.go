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

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next

}
func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}

func (l *LList) Size() int {
	return l.size
}

func (l *LList) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	p := l.root.next

	for p != l.root {
		sb.WriteString(strconv.Itoa(p.Value))
		if p.next != l.root {
			sb.WriteString(", ")
		}
		p = p.next
	}

	sb.WriteString("]")
	return sb.String()
}

func NewLList() *LList {
	rootNode := &Node{Value: 0}
	rootNode.next = rootNode
	rootNode.prev = rootNode
	return &LList{root: rootNode, size: 0}
}

func (l *LList) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func (l *LList) Remove(node *Node) {
	if node == l.root {
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
}

func (l *LList) Insert(node *Node, value int) {
	if node == nil || node == l.root {
		return
	}

	newNode := &Node{
		Value: value,
		prev:  node.prev,
		next:  node,
		root:  l.root,
	}

	node.prev.next = newNode
	node.prev = newNode
	l.size++
}

func (l *LList) Search(value int) *Node {
	for node := l.Front(); node != nil; node = node.Next() {
		if node.Value == value {
			return node
		}
	}
	return nil
}

func (l *LList) PushBack(value int) {
	newNode := &Node{
		Value: value,
		prev:  l.root.prev,
		next:  l.root,
		root:  l.root,
	}

	l.root.prev.next = newNode
	l.root.prev = newNode
	l.size++
}

func (l *LList) PopBack() {
	if l.size == 0 {
		return
	}

	last := l.root.prev

	last.prev.next = l.root
	l.root.prev = last.prev

	l.size--
}

func (l *LList) PushFront(value int) {
	newNode := &Node{
		Value: value,
		prev:  l.root,
		next:  l.root.next,
		root:  l.root,
	}

	l.root.next.prev = newNode
	l.root.next = newNode
	l.size++
}

func (l *LList) PopFront() {
	if l.size == 0 {
		return
	}

	first := l.root.next

	l.root.next = first.next
	first.next.prev = l.root
	l.size--
}

func (l *LList) Front() *Node {
	if l.root.next == l.root {
		return nil
	}
	return l.root.next
}

func (l *LList) Back() *Node {
	if l.root.prev == l.root {
		return nil
	}
	return l.root.prev
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
