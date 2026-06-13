package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Polonesa struct {
	itens []string
}

func (p *Polonesa) PushBack(item string) {
	p.itens = append(p.itens, item)
}

func (p *Polonesa) PopBack() string {
	if len(p.itens) == 0 {
		return ""
	}
	idc := len(p.itens) - 1
	item := p.itens[idc]
	p.itens = p.itens[:idc]
	return item
}

func (p *Polonesa) Top() string {
	if len(p.itens) == 0 {
		return ""
	}
	return p.itens[len(p.itens)-1]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n := scanner.Text()

	p := Polonesa{}
	var exit []string

	num := strings.Fields(n)

	for _, v := range num {
		if unicode.IsDigit(rune(v[0])) {
			exit = append(exit, v)
		} else {
			for len(p.itens) > 0 && acento(p.Top()) >= acento(v) {
				top := p.PopBack()
				exit = append(exit, top)
			}
			p.PushBack(v)
		}
	}
	for len(p.itens) > 0 {
		exit = append(exit, p.PopBack())
	}
	fmt.Println(strings.Join(exit, " "))

}

func acento(operador string) int {
	if operador == "+" || operador == "-" {
		return 1
	} else if operador == "*" || operador == "/" {
		return 2
	} else {
		return 3
	}
}
