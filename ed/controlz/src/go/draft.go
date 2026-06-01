package main

import (
	"bufio"
	"fmt"
	"os"
)

type Action struct {
	kind string
	ch   rune
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	var left []rune
	var right []rune
	var undo []Action
	var redo []Action

	for _, c := range input {
		switch c {

		case 'R':
			left = append(left, '\n')

			undo = append(undo, Action{
				kind: "insert",
				ch:   '\n',
			})

			redo = nil

		case 'B':
			if len(left) > 0 {
				removed := left[len(left)-1]

				left = left[:len(left)-1]

				undo = append(undo, Action{
					kind: "delete",
					ch:   removed,
				})

				redo = nil
			}

		case 'D':
			if len(right) > 0 {
				removed := right[0]

				right = right[1:]

				undo = append(undo, Action{
					kind: "delete",
					ch:   removed,
				})

				redo = nil
			}

		case '<':
			if len(left) > 0 {
				last := left[len(left)-1]
				left = left[:len(left)-1]
				right = append([]rune{last}, right...)
			}

		case '>':
			if len(right) > 0 {
				left = append(left, right[0])
				right = right[1:]
			}

		case 'Z':
			if len(undo) > 0 {
				action := undo[len(undo)-1]
				undo = undo[:len(undo)-1]

				switch action.kind {
				case "insert":
					if len(left) > 0 {
						left = left[:len(left)-1]
					}

				case "delete":
					left = append(left, action.ch)
				}

				redo = append(redo, action)
			}

		case 'Y':
			if len(redo) > 0 {
				action := redo[len(redo)-1]
				redo = redo[:len(redo)-1]

				switch action.kind {
				case "insert":
					left = append(left, action.ch)

				case "delete":
					if len(left) > 0 {
						left = left[:len(left)-1]
					}
				}

				undo = append(undo, action)
			}

		default:
			left = append(left, c)

			undo = append(undo, Action{
				kind: "insert",
				ch:   c,
			})

			redo = nil
		}
	}

	fmt.Print(string(left))
	fmt.Print("|")
	fmt.Println(string(right))
}
