package main

import (
	"fmt"
	"strings"
)

func decodeString(s string) string {
	type Numstr struct {
		num int
		s   string
	}
	stack := []byte{}
	news := ""
	hash := []Numstr{}
	flag := false
	numstr := new(Numstr)
	ss := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			stack = append(stack, s[i])
			numstr.num = int(s[i-1] - '0')
			flag = true
		}
		if flag == true && s[i] != '[' && s[i] != ']' {
			ss = append(ss, s[i])
		}
		if s[i] == ']' {
			stack = stack[:len(stack)-1]
			numstr.s = string(ss)
			hash = append(hash, *numstr)
			flag = false
			ss = ss[:0]
		}
	}
	for i := 0; i < len(hash); i++ {
		news += strings.Repeat(hash[i].s, hash[i].num)
		fmt.Println(news)
	}
	return string(news)
}

func main() {
	s1 := "3[a]2[bc]" //"aaabcbc"
	// s2 := "3[a2[c]]"      //"accaccacc"
	// s3 := "2[abc]3[cd]ef" //"abcabccdcdcdef"
	// s4 := "abc3[cd]xyz"   //"abccdcdcdxyz"
	fmt.Printf("\n\n%s\n", decodeString(s1))
}
