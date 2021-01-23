package main

import "fmt"

type MyQueue struct {
	instack  []int
	outstack []int
}

/** Initialize your data structure here. */
func queConstructor() MyQueue {
	q := new(MyQueue)
	q.instack = []int{}
	q.outstack = []int{}
	return *q
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.instack = append(this.instack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.instack) != 0 {
		s := []int{}
		s = append(s, this.outstack...)
		this.outstack = this.outstack[:0]
		for i := 0; i < len(this.instack); i++ {
			this.outstack = append(this.outstack, this.instack[len(this.instack)-i-1])
		}
		this.outstack = append(this.outstack, s...)
		this.instack = this.instack[:0]
		s = s[:0]
	}
	p := this.outstack[len(this.outstack)-1]
	this.outstack = this.outstack[:len(this.outstack)-1]
	return p
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.instack) != 0 {
		s := []int{}
		s = append(s, this.outstack...)
		this.outstack = this.outstack[:0]
		for i := 0; i < len(this.instack); i++ {
			this.outstack = append(this.outstack, this.instack[len(this.instack)-i-1])
		}
		this.outstack = append(this.outstack, s...)
		this.instack = this.instack[:0]
		s = s[:0]
	}
	return this.outstack[len(this.outstack)-1]

}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.instack) == 0 && len(this.outstack) == 0 {
		return true
	}
	return false
}

func main() {
	myqueue := queConstructor()
	myqueue.Push(1)
	myqueue.Push(2)
	myqueue.Push(3)
	myqueue.Push(4)
	fmt.Println(myqueue)
	fmt.Println(myqueue.Pop())
	myqueue.Push(5)
	fmt.Println(myqueue)
	fmt.Println(myqueue.Pop())
	fmt.Println(myqueue.Pop())
	fmt.Println(myqueue.Pop())
	fmt.Println(myqueue.Pop())
	fmt.Println(myqueue.Empty())
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
