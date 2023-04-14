package main

import (
	"fmt"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func isValid(s string) bool {
	if len(s) <= 1 || len(s)%2 != 0 {
		fmt.Println("Len of s is not acceptable")
		return false
	}

	var stack Stack      // create a stack to save open caracteres
	var searchFor string // create a string to save desired open bracket
	var count int = 0    // counting

	for _, ch := range s {

		if string(ch) == "(" || string(ch) == "[" || string(ch) == "{" {
			stack.Push(string(ch))
			//fmt.Println(string(ch), "is an open caracter")
			count++
		} else {
			//fmt.Println(string(ch), "is a closed caracter")

			if count == 0 {
				return false // first interaction cannot be a closed caracter
			}

			x, y := stack.Pop() // take top list from stack
			if y == true {
				if string(ch) == "]" { // check with kind of closed bracket it is
					searchFor = "["
				} else if string(ch) == ")" {
					searchFor = "("
				} else if string(ch) == "}" {
					searchFor = "{"
				} else {
					return false
				}

				// check if top item from stack is open it
				if x != searchFor {
					fmt.Println(x, ":", searchFor)
					return false
				}

				count++

			} else {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

func main() {
	var s string = "))"
	fmt.Println(isValid(s))
}
