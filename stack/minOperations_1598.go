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
func minOperations2(logs []string) int {
	cnt := 0
	for _, val := range logs {
		if !strings.Contains(val, ".") {
			cnt++
		}
		if cnt > 0 && strings.Contains(val, "..") {
			cnt--
		}
	}
	return cnt
}

func main() {
	logs := []string{"d1/", "../", "./", "d2/"}
	fmt.Println(minOperations(logs))
	fmt.Println(minOperations2(logs))
}
