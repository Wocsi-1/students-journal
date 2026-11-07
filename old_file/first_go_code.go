package main

import (
	"fmt"
	"time"
)

func main() {
	var n int = 1000000
	var TestSlice = []int{}
	var TestSlice2 = make([]int, 0, n)

	fmt.Printf("Total timr without preallocation: %v\n", timeLoop(TestSlice, n))
	fmt.Printf("Total time with preallocation: %v", timeLoop(TestSlice2, n))
}
func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
