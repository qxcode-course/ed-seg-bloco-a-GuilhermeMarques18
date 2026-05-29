package main

import (
	"fmt"
	"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	sb := strings.Builder{}
	sb.WriteRune('[')
	sb.WriteRune(' ')
	for node := l.Front(); node != l.End(); node = node.next {
		if node == sword {
			sb.WriteString(fmt.Sprintf("%v> ", node.Value))
		} else {
			sb.WriteString(fmt.Sprintf("%v ", node.Value))
		}
	}
	sb.WriteRune(']')
	return sb.String()
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	nextNo := it.next
	if nextNo == l.End() {
		nextNo = nextNo.next
	}
	return nextNo
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)

	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
