package main

import "fmt"

type MinStack struct {
	num []int
	min []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	minstack := new(MinStack)
	minstack.num = []int{}
	minstack.min = []int{}
	return *minstack
}

func (this *MinStack) Push(x int) {
	this.num = append(this.num, x)
	if len(this.min) == 0 || x <= this.min[len(this.min)-1] {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.num) < 0 {
		return
	}
	if this.Top() == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.num = this.num[:len(this.num)-1]

}

func (this *MinStack) Top() int {
	if len(this.num) < 0 {
		return -1
	}
	return this.num[len(this.num)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

func main() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println(stack.GetMin())
	fmt.Println(stack)
	fmt.Println(stack.Top())
	stack.Pop()
	fmt.Println(stack)
	fmt.Println(stack.GetMin())
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
