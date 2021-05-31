package gredy

/*
title:broken calculator
content:
On a broken calculator that has a number showing on its display, we can perform two operations:

    Double: Multiply the number on the display by 2, or;
    Decrement: Subtract 1 from the number on the display.

Initially, the calculator is displaying the number x.
Return the minimum number of operations needed to display the number y.
Example one:
Input: x = 2, y = 3
Output: 2
Explanation: Use double operation and then decrement operation {2 -> 4 -> 3}.

题目大意：x经过double、decrement两者或其一的多次操作转换为y，输出为操作的次数
解题思路：x > y ,返回x-y , y>x ,一直折半y 返回x - y
*/
func brokenCalc(x int, y int) int {
	cnt := 0
	if x > y {
		return x - y
	} else if x < y {
		for x < y {
			if y%2 == 0 {
				y = y / 2
			} else {
				y = (y + 1) / 2
				cnt++
			}
			cnt++
		}
		return x - y + cnt

	}
	return 0
}
