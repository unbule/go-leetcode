package main

import "fmt"

func backspaceCompare(S string, T string) bool {
	sstack := []byte{}
	tstack := []byte{}
	for _, val := range S {
		if val == '#' {
			if len(sstack) > 0 {
				sstack = sstack[:len(sstack)-1]
			}
		} else {
			sstack = append(sstack, byte(val))
		}
	}
	for _, val := range T {
		if val == '#' {
			if len(tstack) > 0 {
				tstack = tstack[:len(tstack)-1]
			}
		} else {
			tstack = append(tstack, byte(val))
		}
	}
	fmt.Println(sstack)
	fmt.Println(tstack)
	if len(sstack) == len(tstack) {
		for i, val := range sstack {
			if tstack[i] != val {
				return false
			}
		}
	} else {
		return false
	}
	return true
}
func main() {

	S := "xywrrmp"
	T := "xywrrm#p"
	fmt.Println(backspaceCompare(S, T))
}
