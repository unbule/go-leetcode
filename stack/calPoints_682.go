package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calPoints(ops []string) int {
	stack := []string{}
	for _, val := range ops {
		if strings.Contains(val, "C") {
			stack = stack[:len(stack)-1]
		} else if strings.Contains(val, "D") {
			num, _ := strconv.Atoi(stack[len(stack)-1])
			num = num * 2
			stack = append(stack, strconv.Itoa(num))
		} else if strings.Contains(val, "+") {
			num1, _ := strconv.Atoi(stack[len(stack)-1])
			num2, _ := strconv.Atoi(stack[len(stack)-2])
			sum := num1 + num2
			stack = append(stack, strconv.Itoa(sum))
		} else {
			stack = append(stack, val)
		}
	}
	cnt := 0
	for _, val := range stack {
		num, _ := strconv.Atoi(val)
		cnt += num
	}
	return cnt
}

func main() {
	ops := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	fmt.Println(calPoints(ops))
}
