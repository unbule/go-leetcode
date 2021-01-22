package main

import "fmt"

func removeDuplicates(S string) string {
	stack := []byte{}
	stack = append(stack, 0)
	for i := 0; i < len(S); i++ {
		if stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}

	}
	stack = stack[1:]
	return string(stack)
}

func main() {
	s := "abbbaca"
	fmt.Println(removeDuplicates(s))
}
