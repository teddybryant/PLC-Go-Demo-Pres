package main

import (
	"fmt"
	"time"
)

func sumFibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	a, b := 0, 1
	sum := 1
	for i := 2; i < n; i++ {
		a, b = b, a+b
		sum += b
	}
	return sum
}

func main() {
	n := 1000000
	start := time.Now()
	sum := sumFibonacci(n)
	elapsed := time.Since(start)
	fmt.Printf("Sum of the first %d Fibonacci numbers: %d\n", n, sum)
	fmt.Printf("Go's Execution time: %s\n", elapsed)
}
