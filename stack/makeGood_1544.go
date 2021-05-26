package main

import (
	"fmt"
	"unicode"
)

func makeGood(s string) string {

	for i := 0; i < len(s)-1; i++ {
		if s[i] == byte(unicode.ToLower(rune(s[i+1]))) || s[i] == byte(unicode.ToUpper(rune(s[i+1]))) {
			if !(s[i] == s[i+1]) {
				s = s[:i] + s[i+2:]
				i = -1
			}
		}

	}

	return s
}
func main() {
	// s := "leEeetcode"
	// s1 := "abBAcC"
	s2 := "qFxXfQo"
	// fmt.Println(makeGood(s))
	// fmt.Println(makeGood(s1))
	fmt.Println(makeGood(s2))
}
