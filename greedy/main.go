package main

import (
	"fmt"
	"time"
)

func main() {
	x := 1
	fmt.Println(x)
	go func() {
		time.Sleep(10)
	}()
	print("1111")
}
func print(str interface{}) {
	fmt.Println(str)
}
