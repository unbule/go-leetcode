package heap

import (
	"container/heap"
	"fmt"
)

type Heap struct {
	Array []int
}

func New(nums []int) *Heap {
	heap.Interface
	arr := make([]int, len(nums)+1)
	arr[0] = 0
	arr = append(arr, nums...)
	return &Heap{
		Array: arr,
	}
}

func main() {
	num := []int{3, 2, 1, 5, 6, 4}
	num2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 2
	k2 := 4
	fmt.Println(findKthLargest(num, k))
	fmt.Println(findKthLargest(num2, k2))
}

func findKthLargest(nums []int, k int) int {
	max := nums[0]
	cnt := 0
	fmt.Println(arr)
	for i := 0; i < arr.Len(); i++ {
		if arr[i] > max {
			cnt++
		}
		if cnt == k {
			return arr[i]
		}
	}
	return 0
}
