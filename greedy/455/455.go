package greedy

import "sort"

/*
title:Assign Cookies
Assume you are an awesome parent and want to give your children some cookies.
But, you should give each child at most one cookie.
Each child i has a greed factor g[i], which is the minimum size of a cookie
that the child will be content with; and each cookie j has a size s[j].
If s[j] >= g[i], we can assign the cookie j to the child i,
and the child i will be content.
Your goal is to maximize the number of your content children and output the maximum number.

*/
/*
	题目大意：child对应i，g[i]里是child满意的最小饼干尺寸，cookie对应j，s[j]代表饼干尺寸，
	如果s[j] >= g[i],小孩拿到饼干就会很满意
*/
/*
思路: 先排序，排序后选择满足小孩胃口并且差值最小的饼干。
*/
func findContentChildren(g []int, s []int) (ans int) {
	sort.Ints(g)
	sort.Ints(s)
	glen := len(g)
	slen := len(s)
	for i, j := 0, 0; i < glen && j < slen; i++ {
		//遍历到满足小孩胃口，差值最小的饼干
		for j < slen && g[i] > s[j] {
			j++
		}
		if j < slen {
			ans++
			j++
		}
	}
	return
}
