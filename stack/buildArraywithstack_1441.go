package main

import "fmt"

func buildArray(target []int, n int) []string {
	arr := []string{}
	for i := 0; i < n; i++ {
		flag := false
		for j := 0; j < len(target); j++ {
			if target[j] == i+1 {
				target = target[1:]
				arr = append(arr, "Push")
				flag = true
			}
		}
		if !flag {
			arr = append(arr, "Push", "Pop")
		}
		if len(target) == 0 {
			return arr
		}
	}
	return arr
}

func main() {
	target := []int{1, 3}
	n := 3
	fmt.Println(buildArray(target, n))
}
