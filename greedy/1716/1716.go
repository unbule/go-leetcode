package greedy

/*
title :Calculate Money in Leetcode Bank
content:
Hercy wants to save money for his first car. He puts money in the Leetcode bank every day.
He starts by putting in $1 on Monday, the first day.
Every day from Tuesday to Sunday,
he will put in $1 more than the day before. On every subsequent Monday,
he will put in $1 more than the previous Monday.
Given n, return the total amount of money he will have in the Leetcode bank at the end of the nth day.
题目大意：hercy每天都要存钱，在一周内，每天比之前多存一元。每周的周一则比上一周的周一多一元
思路：n对7取余，计算出有多少周，周数+1则是每周开始的钱数，最后取余的结果是剩下多少天
*/

func totalMoney(n int) (total int) {
	//周数
	WeekDays := n / 7
	//最后一周的天数
	laskWeekDays := n % 7
	for i := 0; i < WeekDays; i++ {
		for j := i + 1; j <= i+7; j++ {
			total = total + j
		}
	}
	for i := 0; i < laskWeekDays; i++ {
		WeekDays = WeekDays + 1
		total = total + WeekDays
	}
	return
}
