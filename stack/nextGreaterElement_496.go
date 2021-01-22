package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	num3 := []int{}

	for i := 0; i < len(nums1); i++ {
		flag := false
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				flag = true
			}
			if flag == true && nums1[i] < nums2[j] {
				num3 = append(num3, nums2[j])
				break
			}
		}
		if len(num3) < i+1 {
			num3 = append(num3, -1)
		}

	}
	return num3
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
