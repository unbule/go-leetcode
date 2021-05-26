package greedy

import "strings"

/*
title:1190. Reverse Substrings Between Each Pair of Parentheses

You are given a string s that consists of lower case English letters and brackets.

Reverse the strings in each pair of matching parentheses, starting from the innermost one.

Your result should not contain any brackets.

https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses/
*/
/*
题目大意：从最里面的括号开始，翻转每一个括号里的字母
*/
/*
解题思路：记录括号的位置，遇到右括号时翻转右括号位置到左括号位置的字符串,最后清除字符串中的括号。
*/

func reverseParentheses(s string) string {
	str := []byte(s)
	var stack []int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			//左括号入栈
			stack = append(stack, i)
		} else if s[i] == ')' {
			index := stack[len(stack)-1]
			//翻转字符串
			for l, r := index+1, i-1; l < r; l, r = l+1, r-1 {
				str[l], str[r] = str[r], str[l]
			}
			//左括号出栈
			stack = stack[:len(stack)-1]
		}
	}
	//去掉所有左括号和右括号
	sb := strings.Builder{}
	for i := 0; i < len(str); i++ {
		if str[i] != '(' && str[i] != ')' {
			sb.WriteByte(str[i])
		}
	}
	return sb.String()
}
