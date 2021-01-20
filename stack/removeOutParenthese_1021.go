package main

import (
	"fmt"
)

func main() {
	str := "(()())(())"
	patternArr := []int{}
	stack := []byte{}
	j := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			stack = append(stack, str[i])
		} else {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			patternArr = append(patternArr, i+1)
			j++
		}
	}
	c := []byte(str)
	c[0] = ' '
	for i := 0; i < j; i++ {
		c[patternArr[i]-1] = ' '
	}
	for i := 0; i < j-1; i++ {
		c[patternArr[i]] = ' '
	}
	new := []byte{}
	for i := 0; i < len(c); i++ {
		if c[i] != ' ' {
			new = append(new, c[i])
		}
	}
	fmt.Println(string(new))
}
