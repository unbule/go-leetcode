package main

import (
	"fmt"
	"strings"
)

func minOperations(logs []string) int {
	stack := []string{}
	for _, val := range logs {
		if !strings.Contains(val, ".") {
			stack = append(stack, val)
		}
		if len(stack) > 0 && strings.Contains(val, "..") {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack)
}

func main() {
	logs := []string{"d1/", "../", "../", "../"}
	minOperations(logs)
	fmt.Println(minOperations(logs))
}
