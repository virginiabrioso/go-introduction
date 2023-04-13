import (
	"fmt"
	"os"
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
	if len(s) > 10^4 || len(s) < 1 {
		fmt.Println("Len of s is not acceptable")
		os.Exit(1)
	}

	var stack Stack // create a stack variable of type Stack

	for _, ch := range s {
		stack.Push(string(ch))
	}

	var searchFor string
	var orgStackLen int
	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {

			if x == "(" || x == "[" || x == "{" {
				searchFor = "nothing"
			}

			if x == "]" {
				searchFor = "["
			} else if x == ")" {
				searchFor = "("
			} else if x == "}" {
				searchFor = "{"
			} else {
				searchFor = "nothing"
			}

			if searchFor == "nothing" {
				break // break here
			}

			orgStackLen = len(stack)

			// search for opening caracter and remove it
			for j := 0; j < len(stack); j++ {
				if stack[j] == searchFor {
					stack = (stack)[:j]
				}
			}

			//if no opening caracter found break it
			if orgStackLen == len(stack) {
				break
			}
		}
	}

	return stack.IsEmpty()
}